package dockerhub

import (
	"context"
	"github.com/docker/hub-tool/pkg/hub"
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
			KeyColumns: []*plugin.KeyColumn{
				{Name: "namespace", Require: plugin.Optional, Operators: []string{"="}},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace of the repository.",
				Transform:   transform.From(fetchNamespaceFromRepository),
			},
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

	// Get namespace from "Quals" if available otherwise use the authenticated user's namespace
	namespace, err := getNamespace(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("dockerhub_repository.listRepositories", "connection_error", err)
		return nil, err
	}

	repositories, _, err := client.GetRepositories(namespace)
	if err != nil {
		logger.Error("dockerhub_repository.listRepositories", "api_error", err)
		return nil, err
	}

	for _, repository := range repositories {
		d.StreamListItem(ctx, repository)
	}

	return nil, nil
}

func getNamespace(ctx context.Context, d *plugin.QueryData) (string, error) {
	logger := plugin.Logger(ctx)

	if d.Quals["namespace"] != nil {
		for _, q := range d.Quals["namespace"].Quals {
			if q.Operator == "=" {
				return q.Value.GetStringValue(), nil
			}
		}
	}

	user, err := getUserInfo(ctx, d)
	if err != nil {
		logger.Error("dockerhub_repository.getUserInfo", "error", err)
		return "", err
	}

	return user.Name, nil
}

func fetchNamespaceFromRepository(_ context.Context, d *transform.TransformData) (interface{}, error) {
	repository := d.HydrateItem.(hub.Repository)
	namespace, _ := splitRepositoryName(repository.Name)
	return namespace, nil
}
