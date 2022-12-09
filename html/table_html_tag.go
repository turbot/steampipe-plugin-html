package html

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func tagCols() []*plugin.Column {
	return []*plugin.Column{
		{Name: "page", Type: proto.ColumnType_STRING, Description: "The page containing tags."},
		{Name: "tag_name", Type: proto.ColumnType_STRING, Description: "The name of the tag."},
		{Name: "tag_content", Type: proto.ColumnType_STRING, Description: "The content of the tag."},
		{Name: "tag_attrs", Type: proto.ColumnType_JSON, Description: "The tag's attributes."},
	}
}

func tableHtmlTag(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "html_tag",
		Description: "Tags found in a web page.",
		List: &plugin.ListConfig{
			Hydrate: listTags,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "page", Require: plugin.Required},
				{Name: "tag_name", Require: plugin.Required},
			},
		},
		Columns: tagCols(),
	}
}

func listTags(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	page := d.KeyColumnQuals["page"].GetStringValue()
	tag_name := d.KeyColumnQuals["tag_name"].GetStringValue()
	tag_results := findTags(page, tag_name)
	for i := 0; i < len(tag_results); i++ {
		tag_result := tag_results[i]
		d.StreamListItem(ctx, tag_result)
	}
	return nil, nil
}
