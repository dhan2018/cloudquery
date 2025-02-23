package client

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"

	// Import all autorest modules
	_ "github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	subscriptions []string
	services      map[string]*services.Services
	logger        hclog.Logger
	// this is set by table client multiplexer
	SubscriptionId string
	orgId          int
}

func NewAzureClient(log hclog.Logger, subscriptionId []string) *Client {
	return &Client{
		logger:        log,
		subscriptions: subscriptionId,
		services:      make(map[string]*services.Services),
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// SetSubscriptionServices adds a set of services to an azure subscription
func (c *Client) SetSubscriptionServices(subscriptionId string, servicesSet services.Services) {
	c.services[subscriptionId] = &servicesSet
}

// Services returns all available services of client, returns subscriptionId's specific services
func (c Client) Services() *services.Services {
	return c.services[c.SubscriptionId]
}

// withSubscription allows multiplexer to create a new client with given subscriptionId
func (c Client) withSubscription(subscriptionId string) *Client {
	return &Client{
		subscriptions:  c.subscriptions,
		services:       c.services,
		logger:         c.logger.With("subscription_id", subscriptionId),
		SubscriptionId: subscriptionId,
		orgId:          c.orgId,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	providerConfig := config.(*Config)

	// Old-SDK auth with ordering:
	//  1. Environment
	//  2. AzureCLI
	logger.Info("Trying to authenticate via environment variables")
	creds, err := auth.GetSettingsFromEnvironment()
	creds_json, err := json.Marshal(creds.Values)
	logger.Info(string(creds_json))
	logger.Info("dumping error")
	if err != nil {
		logger.Error(err.Error())
	}
	azureAuth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		logger.Error(err.Error())
		logger.Info("Trying to authenticate via CLI")
		azureAuth, err = auth.NewAuthorizerFromCLI()
		if err != nil {
			logger.Error(err.Error())
			return nil, diag.FromError(err, diag.USER)
		}
	}

	// New-SDK auth; chained credentials with ordering:
	//  1. Environment
	//  2. ManagedIdentity
	//  3. AzureCLI
	azCred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, diag.FromError(err, diag.USER)
	}

	orgIdStr := os.Getenv("ORG_ID")
	if orgIdStr == "" {
		logger.Error("ORG_ID environment variable is not set.")
		return nil, diag.FromError(err, diag.USER)
	}
	orgId, err := strconv.Atoi(orgIdStr)
	if err != nil {
		logger.Error("ORG_ID environment variable is not a valid integer.")
		return nil, diag.FromError(err, diag.USER)
	}
	logger.Info("Org Id is", "org_id_str", orgIdStr, "org_id", orgId)
	client := NewAzureClient(logger, providerConfig.Subscriptions)
	client.orgId = orgId
	logger.Info("client.orgId", "org_id", client.orgId)

	if len(providerConfig.Subscriptions) == 0 {
		ctx := context.Background()
		svc := subscription.NewSubscriptionsClient()
		svc.Authorizer = azureAuth
		res, err := svc.List(ctx)
		if err != nil {
			return nil, classifyError(err, diag.USER, "")
		}
		subscriptions := make([]string, 0)
		for res.NotDone() {
			for _, sub := range res.Values() {
				switch sub.State {
				case subscription.Disabled:
					logger.Info("Not fetching from subscription because it is disabled", "subscription", *sub.SubscriptionID)
				case subscription.Deleted:
					logger.Info("Not fetching from subscription because it is deleted", "subscription", *sub.SubscriptionID)
				default:
					subscriptions = append(subscriptions, *sub.SubscriptionID)
				}
			}
			err := res.NextWithContext(ctx)
			if err != nil {
				return nil, classifyError(err, diag.INTERNAL, "")
			}
		}
		client.subscriptions = subscriptions
		logger.Info("No subscriptions specified, going to using all available ones", "subscriptions", subscriptions)
	}

	if len(client.subscriptions) == 0 {
		return nil, diag.FromError(errors.New("could not find any subscription"), diag.USER)
	}

	for _, sub := range client.subscriptions {
		svcs, err := services.InitServices(sub, azureAuth, azCred)
		if err != nil {
			return nil, classifyError(err, diag.INTERNAL, sub)
		}
		client.SetSubscriptionServices(sub, svcs)
	}

	// Return the initialized client and it will be passed to your resources
	return client, nil
}
