---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d1a"
display_name: html
name: html
description: Steampipe plugin for HTML
---

# html

[html]() is 


## Installation

Build and use locally.


## Tables

```
> .inspect html
+--------------------+-----------------------------------------------------------------------------+
| TABLE              | DESCRIPTION                                                                 |
+--------------------+-----------------------------------------------------------------------------+
| html               | Query web pages                                                             |
+--------------------+-----------------------------------------------------------------------------+
```

## Credentials

None.

## Connection Configuration

Put this into ~/.steampipe/config/html.spc`.

```
connection "html" {
  plugin    = "html"
}
```
