package html

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-html",
		DefaultTransform: transform.FromJSONTag(),
		TableMap: map[string]*plugin.Table{
			"html_table": tableHtmlTable(ctx),
			"html_link":  tableLinkTable(ctx),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
	}
	return p
}
