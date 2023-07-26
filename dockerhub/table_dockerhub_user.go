package dockerhub

import (
	"context"

	"github.com/docker/hub-tool/pkg/hub"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDockerHubUser(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "dockerhub_user",
		Description: "Get details of the logged in user.",
		List: &plugin.ListConfig{
			Hydrate: listUserInfo,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "ID of the user.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the user.",
			},
			{
				Name:        "full_name",
				Type:        proto.ColumnType_STRING,
				Description: "Full name of the user.",
			},
			{
				Name:        "location",
				Type:        proto.ColumnType_STRING,
				Description: "Location of the user.",
			},
			{
				Name:        "company",
				Type:        proto.ColumnType_STRING,
				Description: "Company associated with the user.",
			},
			{
				Name:        "joined",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Timestamp indicating when the user joined.",
			},
			{
				Name:        "seats",
				Type:        proto.ColumnType_INT,
				Description: "Number of seats consumed.",
				Hydrate:     getUserConsumption,
			},
			{
				Name:        "private_repositories",
				Type:        proto.ColumnType_INT,
				Description: "Number of private repositories consumed.",
				Hydrate:     getUserConsumption,
			},
			{
				Name:        "teams",
				Type:        proto.ColumnType_INT,
				Description: "Number of teams consumed.",
				Hydrate:     getUserConsumption,
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listUserInfo(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	user, err := getUserInfo(ctx, d)
	if err != nil {
		logger.Error("dockerhub_user.listUserInfo", "error", err)
		return nil, err
	}

	d.StreamListItem(ctx, user)

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getUserConsumption(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	userName := h.Item.(*hub.Account).Name

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("dockerhub_user.getUserConsumption", "connection_error", err)
		return nil, err
	}

	consumption, err := client.GetUserConsumption(userName)
	if err != nil {
		logger.Error("dockerhub_user.getUserConsumption", "api_error", err)
		return nil, err
	}

	return consumption, nil
}
