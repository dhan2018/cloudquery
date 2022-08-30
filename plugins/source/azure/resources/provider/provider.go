package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/account"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/authorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/batch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/compute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/container"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/cosmosdb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/datalake"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/eventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/frontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/iothub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/keyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/logic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/monitor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/mysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/network"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/postgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/redis"
	resources2 "github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/search"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/security"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/servicebus"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/sql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/storage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/streamanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/subscription"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/web"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	Version = "Development"
)

func AddDefaultColumns(table *schema.Table) *schema.Table {
	// Add the default columns to the main tables
	defaultColumns := []schema.Column{{
		Name:        "org_id",
		Description: "FinOps Org Id",
		Type:        schema.TypeInt,
		Resolver:    client.ResolveOrgId,
	}}
	table.Columns = append(defaultColumns, table.Columns...)

	// We also need to add the default columns to the relation tables
	for i := 0; i < len(table.Relations); i++ {
		table.Relations[i].Columns = append(defaultColumns, table.Relations[i].Columns...)
	}

	return table
}

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:         Version,
		Name:            "azure",
		Configure:       client.Configure,
		ErrorClassifier: client.ErrorClassifier,
		ResourceMap: map[string]*schema.Table{
			"account.locations":                  AddDefaultColumns(account.Locations()),
			"authorization.role_assignments":     AddDefaultColumns(authorization.AuthorizationRoleAssignments()),
			"authorization.role_definitions":     AddDefaultColumns(authorization.AuthorizationRoleDefinitions()),
			"batch.accounts":                     AddDefaultColumns(batch.BatchAccounts()),
			"cdn.profiles":                       AddDefaultColumns(cdn.Profiles()),
			"compute.disks":                      AddDefaultColumns(compute.ComputeDisks()),
			"compute.virtual_machines":           AddDefaultColumns(compute.ComputeVirtualMachines()),
			"compute.virtual_machine_scale_sets": AddDefaultColumns(compute.VirtualMachineScaleSets()),
			"container.managed_clusters":         AddDefaultColumns(container.ContainerManagedClusters()),
			"container.registries":               AddDefaultColumns(container.ContainerRegistries()),
			"cosmosdb.accounts":                  AddDefaultColumns(cosmosdb.CosmosDBAccounts()),
			"cosmosdb.sql_databases":             AddDefaultColumns(cosmosdb.CosmosDBSqlDatabases()),
			"cosmosdb.mongodb_databases":         AddDefaultColumns(cosmosdb.CosmosDBMongoDBDatabases()),
			"datalake.storage_accounts":          AddDefaultColumns(datalake.StorageAccounts()),
			"datalake.analytics_accounts":        AddDefaultColumns(datalake.AnalyticsAccounts()),
			"eventhub.namespaces":                AddDefaultColumns(eventhub.EventHubNamespaces()),
			"frontdoor.front_doors":              AddDefaultColumns(frontdoor.FrontDoors()),
			"iothub.hubs":                        AddDefaultColumns(iothub.IothubHubs()),
			// This resource is currently not working
			// https://github.com/cloudquery/cq-provider-azure/issues/107
			"keyvault.vaults":      AddDefaultColumns(keyvault.KeyvaultVaults()),
			"keyvault.managed_hsm": AddDefaultColumns(keyvault.KeyvaultManagedHSM()),
			"logic.app_workflows":  AddDefaultColumns(logic.LogicAppWorkflows()),
			"mariadb.servers":      AddDefaultColumns(mariadb.MariadbServers()),
			"monitor.log_profiles": AddDefaultColumns(monitor.MonitorLogProfiles()),
			// This resource is currently not working
			"monitor.diagnostic_settings":          AddDefaultColumns(monitor.MonitorDiagnosticSettings()),
			"monitor.activity_logs":                AddDefaultColumns(monitor.MonitorActivityLogs()),
			"monitor.activity_log_alerts":          AddDefaultColumns(monitor.MonitorActivityLogAlerts()),
			"mysql.servers":                        AddDefaultColumns(mysql.MySQLServers()),
			"network.express_route_circuits":       AddDefaultColumns(network.NetworkExpressRouteCircuits()),
			"network.express_route_gateways":       AddDefaultColumns(network.NetworkExpressRouteGateways()),
			"network.express_route_ports":          AddDefaultColumns(network.NetworkExpressRoutePorts()),
			"network.interfaces":                   AddDefaultColumns(network.NetworkInterfaces()),
			"network.public_ip_addresses":          AddDefaultColumns(network.NetworkPublicIPAddresses()),
			"network.route_filters":                AddDefaultColumns(network.NetworkRouteFilters()),
			"network.route_tables":                 AddDefaultColumns(network.NetworkRouteTables()),
			"network.security_groups":              AddDefaultColumns(network.NetworkSecurityGroups()),
			"network.virtual_networks":             AddDefaultColumns(network.NetworkVirtualNetworks()),
			"network.watchers":                     AddDefaultColumns(network.NetworkWatchers()),
			"postgresql.servers":                   AddDefaultColumns(postgresql.PostgresqlServers()),
			"redis.services":                       AddDefaultColumns(redis.RedisServices()),
			"resources.groups":                     AddDefaultColumns(resources2.ResourcesGroups()),
			"resources.policy_assignments":         AddDefaultColumns(resources2.ResourcesPolicyAssignments()),
			"resources.links":                      AddDefaultColumns(resources2.ResourcesLinks()),
			"security.assessments":                 AddDefaultColumns(security.SecurityAssessments()),
			"search.services":                      AddDefaultColumns(search.SearchServices()),
			"security.auto_provisioning_settings":  AddDefaultColumns(security.SecurityAutoProvisioningSettings()),
			"security.contacts":                    AddDefaultColumns(security.SecurityContacts()),
			"security.pricings":                    AddDefaultColumns(security.SecurityPricings()),
			"security.settings":                    AddDefaultColumns(security.SecuritySettings()),
			"security.jit_network_access_policies": AddDefaultColumns(security.SecurityJitNetworkAccessPolicies()),
			"servicebus.namespaces":                AddDefaultColumns(servicebus.Namespaces()),
			"sql.servers":                          AddDefaultColumns(sql.SQLServers()),
			"sql.managed_instances":                AddDefaultColumns(sql.SqlManagedInstances()),
			"storage.accounts":                     AddDefaultColumns(storage.StorageAccounts()),
			"streamanalytics.jobs":                 AddDefaultColumns(streamanalytics.StreamanalyticsJobs()),
			"subscription.subscriptions":           AddDefaultColumns(subscription.Subscriptions()),
			"subscription.tenants":                 AddDefaultColumns(subscription.Tenants()),
			"web.apps":                             AddDefaultColumns(web.WebApps()),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
