package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource webhooks --config gen.hcl --output .
func Webhooks() *schema.Table {
	return &schema.Table{
		Name:         "aws_codepipeline_webhooks",
		Description:  "The detail returned for each webhook after listing webhooks, such as the webhook URL, the webhook name, and the webhook ARN",
		Resolver:     fetchCodepipelineWebhooks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("codepipeline"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "authentication",
				Description: "Supported options are GITHUB_HMAC, IP, and UNAUTHENTICATED",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Definition.Authentication"),
			},
			{
				Name:        "authentication_allowed_ip_range",
				Description: "The property used to configure acceptance of webhooks in an IP address range For IP, only the AllowedIPRange property must be set",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Definition.AuthenticationConfiguration.AllowedIPRange"),
			},
			{
				Name:        "authentication_secret_token",
				Description: "The property used to configure GitHub authentication",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Definition.AuthenticationConfiguration.SecretToken"),
			},
			{
				Name:        "name",
				Description: "The name of the webhook",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Definition.Name"),
			},
			{
				Name:        "target_action",
				Description: "The name of the action in a pipeline you want to connect to the webhook",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Definition.TargetAction"),
			},
			{
				Name:        "target_pipeline",
				Description: "The name of the pipeline you want to connect to the webhook",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Definition.TargetPipeline"),
			},
			{
				Name:        "url",
				Description: "A unique URL generated by CodePipeline",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the webhook",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_code",
				Description: "The number code of the error",
				Type:        schema.TypeString,
			},
			{
				Name:        "error_message",
				Description: "The text of the error message about the webhook",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_triggered",
				Description: "The date and time a webhook was last successfully triggered, in timestamp format",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "tags",
				Description: "The tags associated with the webhook",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_codepipeline_webhook_filters",
				Description: "The event criteria that specify when a webhook notification is sent to your URL",
				Resolver:    schema.PathTableResolver("Definition.Filters"),
				Columns: []schema.Column{
					{
						Name:        "webhook_cq_id",
						Description: "Unique CloudQuery ID of aws_codepipeline_webhooks table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "json_path",
						Description: "A JsonPath expression that is applied to the body/payload of the webhook",
						Type:        schema.TypeString,
					},
					{
						Name:        "match_equals",
						Description: "The value selected by the JsonPath expression must match what is supplied in the MatchEquals field",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCodepipelineWebhooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().CodePipeline
	config := codepipeline.ListWebhooksInput{}
	for {
		response, err := svc.ListWebhooks(ctx, &config)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Webhooks

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
