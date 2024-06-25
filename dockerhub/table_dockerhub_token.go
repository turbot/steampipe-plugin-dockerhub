package dockerhub

import (
	"context"

	"github.com/docker/hub-tool/pkg/hub"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDockerHubToken(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "dockerhub_token",
		Description: "Get details of all the personal access tokens in your DockerHub.",
		List: &plugin.ListConfig{
			Hydrate: listTokens,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("uuid"),
			Hydrate:    getToken,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "uuid",
				Type:        proto.ColumnType_STRING,
				Description: "Universally Unique Identifier (UUID) of the token.",
				Transform:   transform.FromValue().Transform(fetchUUID),
			},
			{
				Name:        "client_id",
				Type:        proto.ColumnType_STRING,
				Description: "Client ID associated with the token.",
				Transform:   transform.FromField("ClientID"),
			},
			{
				Name:        "creator_ip",
				Type:        proto.ColumnType_STRING,
				Description: "IP address of the creator or originator of the token.",
				Transform:   transform.FromField("CreatorIP"),
			},
			{
				Name:        "creator_ua",
				Type:        proto.ColumnType_STRING,
				Description: "User-Agent (UA) string of the creator or originator of the token.",
				Transform:   transform.FromField("CreatorUA"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating when the token was created.",
			},
			{
				Name:        "last_used",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating the last time the token was used.",
			},
			{
				Name:        "generated_by",
				Type:        proto.ColumnType_STRING,
				Description: "Entity that generated the token.",
			},
			{
				Name:        "is_active",
				Type:        proto.ColumnType_BOOL,
				Description: "Boolean value indicating whether the token is active or not.",
			},
			{
				Name:        "token",
				Type:        proto.ColumnType_STRING,
				Description: "Actual token value used for authentication or authorization purposes.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description or additional information about the token.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromValue().Transform(fetchUUID),
			},
		}),
	}
}

func fetchUUID(_ context.Context, d *transform.TransformData) (interface{}, error) {
	token := d.HydrateItem.(hub.Token)
	return token.UUID.String(), nil
}

//// LIST FUNCTION

func listTokens(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("dockerhub_token.listTokens", "connection_error", err)
		return nil, err
	}

	tokens, _, err := client.GetTokens()
	if err != nil {
		logger.Error("dockerhub_token.listTokens", "api_error", err)
		return nil, err
	}

	for _, token := range tokens {
		d.StreamListItem(ctx, token)
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("uuid")

	// Return nil if the id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("dockerhub_token.getToken", "connection_error", err)
		return nil, err
	}

	token, err := client.GetToken(id)
	if err != nil {
		logger.Error("dockerhub_token.getToken", "api_error", err)
		return nil, err
	}

	return *token, nil
}
