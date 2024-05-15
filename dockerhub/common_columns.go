package dockerhub

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "account_id",
			Description: "Docker Hub account ID.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getDockerHubAccountId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

func getDockerHubAccountId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	acc, err := getUserInfo(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getDockerHubAccountId", err)
		return nil, nil
	}

	return acc.ID, nil
}
