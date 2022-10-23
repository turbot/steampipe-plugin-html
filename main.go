package main

import (
	"github.com/judell/steampipe-plugin-html/html"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: html.Plugin})
}
