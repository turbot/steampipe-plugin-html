package html

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type htmlConfig struct {
	Path *string `cty:"path"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"path": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &htmlConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) htmlConfig {
	if connection == nil || connection.Config == nil {
		return htmlConfig{}
	}
	config, _ := connection.Config.(htmlConfig)
	return config
}
