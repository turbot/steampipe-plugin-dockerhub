package dockerhub

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin/transform"
)

const pluginName = "steampipe-plugin-dockerhub"

// Plugin creates this (dockerhub) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:               pluginName,
		DefaultTransform:   transform.FromCamel().Transform(transform.NullIfZeroValue),
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		DefaultGetConfig: &plugin.GetConfig{
			IgnoreConfig: &plugin.IgnoreConfig{ShouldIgnoreErrorFunc: func(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
				return isNotFoundError([]string{"resource not found"})(err)
			}},
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "account_id",
				Hydrate: getDockerHubAccountId,
			},
		},
		TableMap: map[string]*plugin.Table{
			"dockerhub_repository": tableDockerHubRepository(ctx),
			"dockerhub_tag":        tableDockerHubTag(ctx),
			"dockerhub_token":      tableDockerHubToken(ctx),
			"dockerhub_user":       tableDockerHubUser(ctx),
		},
	}

	return p
}
