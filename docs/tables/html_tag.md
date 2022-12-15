# Table: html_tag

## Examples

## Query attributes of meta tags in a blog page

```
select
  page,
  tag_name,
  tag_attrs
from
  html_tag
where
  page = 'https://steampipe.io/blog/selective-select'
  and tag_name = 'meta'
```

## Query the title of a blog page

```
select
  page,
  tag_name,
  tag_content
from
  html_tag
where
  page = 'https://steampipe.io/blog/selective-select'
  and tag_name = 'title'
```

## Query the citations in a Wikipedia page

```
select
  page,
  tag_name,
  jsonb_pretty(tag_markup) as htmljson
from
  html_tag
where
  page = 'https://en.wikipedia.org/wiki/Screencast'
  and tag_name = 'cite'
```

