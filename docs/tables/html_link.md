# Table: html_link

## Examples

### Query links in the NY Times home page

```
> select
  *
from
  html_link
where
  url = 'https://www.nytimes.com'
```

```
+-------------------------+-------------------------------------------------------------------------------------------------------+-----------------------------------------+
| page                    | url                                                                                                   | label                                   |
+-------------------------+-------------------------------------------------------------------------------------------------------+-----------------------------------------+
| https://www.nytimes.com | https://www.nytimes.com/subscription?campaignId=37WXW                                                 | Subscriptions                           |
| https://www.nytimes.com | /section/technology                                                                                   | Tech                                    |
| https://www.nytimes.com | https://cooking.nytimes.com/                                                                          | Cooking                                 |
| https://www.nytimes.com | https://cn.nytimes.com                                                                                | 中文                                    |
| https://www.nytimes.com | https://www.nytimes.com/section/todayspaper                                                           | Today’s Paper                           |
```

### Check links on a blog page

```
> with urls as (
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

```
+----------------------+--------------------------------------------------------------------------+--------------------------+
| response_status_code | url                                                                      | label                    |
+----------------------+--------------------------------------------------------------------------+--------------------------+
| 200                  | https://steampipe.io/blog                                                | Blog                     |
| 200                  | https://steampipe.io/community/join                                      | Join our Slack Community |
| 200                  | https://steampipe.io/docs/reference/env-vars/steampipe_cache_max_size_mb | STEAMPIPE_MAX_CACHE_SIZE |
| 200                  | https://steampipe.io/#major-memory-reduction                             | Major memory reduction   |
| 200                  | https://hub.steampipe.io/plugins/turbot/github                           | GitHub                   |
```