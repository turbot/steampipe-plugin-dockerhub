package dockerhub

import (
	"context"

	"github.com/docker/hub-tool/pkg/hub"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDockerHubTag(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "dockerhub_tag",
		Description: "Get details of all the tags in your DockerHub.",
		List: &plugin.ListConfig{
			ParentHydrate: listRepositories,
			Hydrate:       listTags,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the tag.",
			},
			{
				Name:        "full_size",
				Type:        proto.ColumnType_INT,
				Description: "Full size of the tag.",
			},
			{
				Name:        "last_updated",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating the last update time of the tag.",
			},
			{
				Name:        "last_updater_user_name",
				Type:        proto.ColumnType_STRING,
				Description: "Username by whom the tag was last updated.",
			},
			{
				Name:        "images",
				Type:        proto.ColumnType_JSON,
				Description: "List of images associated with the tag.",
			},
			{
				Name:        "last_pulled",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating when the tag was last pulled.",
			},
			{
				Name:        "last_pushed",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating when the tag was last pushed.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the tag.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	repo := h.Item.(hub.Repository)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("dockerhub_tag.listTags", "connection_error", err)
		return nil, err
	}

	tags, _, err := client.GetTags(repo.Name)
	if err != nil {
		logger.Error("dockerhub_tag.listTags", "api_error", err)
		return nil, err
	}

	for _, tag := range tags {
		d.StreamListItem(ctx, tag)
	}

	return nil, nil
}
