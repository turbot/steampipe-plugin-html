package html

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

type HtmlLink struct {
	Page  string `json:"page"`
	Url   string `json:"url"`
	Label string `json:"label"`
}

func linkCols() []*plugin.Column {
	return []*plugin.Column{
		{Name: "page", Type: proto.ColumnType_STRING, Description: "The page containing links."},
		{Name: "url", Type: proto.ColumnType_STRING, Description: "The URL of the link."},
		{Name: "label", Type: proto.ColumnType_STRING, Description: "The label of the link."},
	}
}

func tableHtmlLink(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "html_link",
		Description: "Links found in a web page.",
		List: &plugin.ListConfig{
			Hydrate: listHtmlLink,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "page", Require: plugin.Required},
			},
		},
		Columns: linkCols(),
	}
}

func listHtmlLink(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	page := d.KeyColumnQuals["page"].GetStringValue()
	labels, urls := findLinks(page)
	plugin.Logger(ctx).Info("listHtmlTable", "labels", labels, "urls", urls)
	for i := 0; i < len(labels); i++ {
		label := labels[i]
		url := urls[i]
		item := HtmlLink{page, url, label}
		d.StreamListItem(ctx, item)
	}
	return nil, nil
}
