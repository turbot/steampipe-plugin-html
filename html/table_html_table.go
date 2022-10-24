package html

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

type HtmlTable struct {
	BaseName string `json:"base_name"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Path     string `json:"path"`
	Columns  string `json:"columns"`
}

func htmlCols() []*plugin.Column {
	return []*plugin.Column{
		{Name: "base_name", Type: proto.ColumnType_STRING, Description: "Base name for CSV tables scraped from a web page."},
		{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of a CSV table scraped from a web page."},
		{Name: "url", Type: proto.ColumnType_STRING, Description: "URL of the scraped page."},
		{Name: "path", Type: proto.ColumnType_STRING, Description: "Path to CSV folder."},
		{Name: "columns", Type: proto.ColumnType_STRING, Description: "Columns of a CSV table."},
	}
}

func tableHtmlTable(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "html_table",
		Description: "HTML table from web page",
		List: &plugin.ListConfig{
			Hydrate: listHtmlTable,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "base_name", Require: plugin.Required},
				{Name: "url", Require: plugin.Required},
			},
		},
		Columns: htmlCols(),
	}
}

func listHtmlTable(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	base_name := d.KeyColumnQuals["base_name"].GetStringValue()
	url := d.KeyColumnQuals["url"].GetStringValue()
	config := GetConfig(d.Connection)

	path := d.KeyColumnQuals["path"].GetStringValue()
	if config.Path != nil {
		path = *config.Path
	}

	plugin.Logger(ctx).Info("listHtmlTable", "base_name", base_name, "url", url, "path", path)
	names, columns := createTables(base_name, url, path)
	plugin.Logger(ctx).Info("listHtmlTable", "names", names, "columns", columns)
	for i := 0; i < len(names); i++ {
		name := names[i]
		cols := columns[i]
		plugin.Logger(ctx).Info("listHtmlTable", "number", i, "name", names, "url", url, "cols", cols)
		item := HtmlTable{base_name, name, url, path, cols}
		d.StreamListItem(ctx, &item)
	}
	return nil, nil
}
