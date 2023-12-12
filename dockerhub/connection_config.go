package dockerhub

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type dockerHubConfig struct {
	Username      *string `hcl:"username"`
	Password      *string `hcl:"password"`
	TwoFactorCode *string `hcl:"two_factor_code"`
}

func ConfigInstance() interface{} {
	return &dockerHubConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) dockerHubConfig {
	if connection == nil || connection.Config == nil {
		return dockerHubConfig{}
	}
	config, _ := connection.Config.(dockerHubConfig)
	return config
}
