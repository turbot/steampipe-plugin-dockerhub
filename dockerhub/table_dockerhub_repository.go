package dockerhub

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDockerHubRepository(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "dockerhub_repository",
		Description: "Get details of all the repositories in your DockerHub.",
		List: &plugin.ListConfig{
			Hydrate: listRepositories,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the repository.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description of the repository.",
			},
			{
				Name:        "last_updated",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating the last update time of the repository.",
			},
			{
				Name:        "pull_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of pulls or downloads of the repository.",
			},
			{
				Name:        "star_count",
				Type:        proto.ColumnType_INT,
				Description: "Number of stars or likes received by the repository.",
			},
			{
				Name:        "is_private",
				Type:        proto.ColumnType_BOOL,
				Description: "Boolean value indicating if the repository is private or not.",
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

func listRepositories(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	user, err := getUserInfo(ctx, d)
	if err != nil {
		logger.Error("dockerhub_repository.getUserInfo", "error", err)
		return nil, err
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("dockerhub_repository.listRepositories", "connection_error", err)
		return nil, err
	}

	repositories, _, err := client.GetRepositories(user.Name)
	if err != nil {
		logger.Error("dockerhub_repository.listRepositories", "api_error", err)
		return nil, err
	}

	for _, repository := range repositories {
		d.StreamListItem(ctx, repository)
	}

	return nil, nil
}
