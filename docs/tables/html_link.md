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
| https://www.nytimes.com | /section/us                                                                                           | U.S.                                    |
| https://www.nytimes.com | /section/world                                                                                        | World                                   |
| https://www.nytimes.com | /international/?action=click&region=Editions&pgtype=Homepage                                          | International                           |
| https://www.nytimes.com | /section/nyregion                                                                                     | N.Y.                                    |


