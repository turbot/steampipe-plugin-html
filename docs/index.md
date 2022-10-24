---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d1a"
display_name: html
name: html
description: Steampipe plugin for HTML
---

# HTML + Steampipe

Web pages often contain data in HTML tables. This plugin downloads that data for use with Steampipe.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

## Installation

Build and use locally.

## Tables

```
> .inspect html
+--------------------+-----------------------------------------------------------------------------+
| TABLE              | DESCRIPTION                                                                 |
+--------------------+-----------------------------------------------------------------------------+
| html               | Query web pages for HTML tables, download them to CSV files                 |
+--------------------+-----------------------------------------------------------------------------+
```

## Credentials

None.

## Connection Configuration

Put this into ~/.steampipe/config/html.spc`.

```
connection "html" {
  plugin    = "html"

  path      = "/home/jon/csv" # or wherever your Steampipe CSV plugin looks for CSV files
}
```
