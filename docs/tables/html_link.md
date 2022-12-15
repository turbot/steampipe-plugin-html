# Table: html_link

## Examples

### Query links in the NY Times home page

```
select
  *
from
  html_link
where
  page = 'https://www.nytimes.com'
```

### Check links on a blog page

```
with urls as (
    select
      page,
      case 
        when url ~ '^/' then 'https://steampipe.io' || url
        else url
      end as url,
      label
    from
      html_link
    where
      page = 'https://steampipe.io/blog/release-0-16-0'
  )
  select
    n.response_status_code,
    u.url,
    u.label
  from
    net_http_request n
  join
    urls u
  on 
    n.url = u.url
  where 
    n.url in ( select url from urls )
```
