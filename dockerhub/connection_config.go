package dockerhub

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type dockerHubConfig struct {
	Username      *string `cty:"username"`
	Password      *string `cty:"password"`
	TwoFactorCode *string `cty:"two_factor_code"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"two_factor_code": {
		Type: schema.TypeString,
	},
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
