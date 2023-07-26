package dockerhub

import (
	"context"
	"errors"
	"os"

	"github.com/docker/hub-tool/pkg/hub"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func getClient(ctx context.Context, d *plugin.QueryData) (*hub.Client, error) {
	conn, err := GetNewClientCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*hub.Client), nil
}

var GetNewClientCached = plugin.HydrateFunc(GetNewClientUncached).Memoize()

func GetNewClientUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	// Default to using env vars (#2)
	username := os.Getenv("DOCKER_HUB_USERNAME")
	password := os.Getenv("DOCKER_HUB_PASSWORD")
	twoFactorCode := ""

	// But prefer the config (#1)
	dockerHubConfig := GetConfig(d.Connection)
	if dockerHubConfig.Username != nil {
		username = *dockerHubConfig.Username
	}
	if dockerHubConfig.Password != nil {
		password = *dockerHubConfig.Password
	}
	if dockerHubConfig.TwoFactorCode != nil {
		twoFactorCode = *dockerHubConfig.TwoFactorCode
	}

	// Error if the minimum config is not set
	if username == "" || password == "" {
		return nil, errors.New("username and password must be configured")
	}

	client, err := hub.NewClient(hub.WithHubAccount(username), hub.WithPassword(password))
	if err != nil {
		return nil, err
	}
	token, _, err := client.Login(username, password, func() (string, error) {
		return twoFactorCode, nil
	})
	if err != nil {
		return nil, err
	}

	hubclient, err := hub.NewClient(hub.WithHubToken(token))
	if err != nil {
		return nil, err
	}

	return hubclient, nil
}

func getUserInfo(ctx context.Context, d *plugin.QueryData) (*hub.Account, error) {
	conn, err := GetUserInfoCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}

	return conn.(*hub.Account), nil
}

var GetUserInfoCached = plugin.HydrateFunc(GetUserInfoUncached).Memoize()

func GetUserInfoUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("GetUserInfoUncached", "connection_error", err)
		return nil, err
	}

	user, err := client.GetUserInfo()
	if err != nil {
		logger.Error("GetUserInfoUncached", "api_error", err)
		return nil, err
	}

	return user, nil
}
