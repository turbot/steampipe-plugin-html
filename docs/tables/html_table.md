# Table: html_table

## Examples

### Download an HTML table from a Wikipedia page

```
select
  base_name,
  name,
  path,
  columns
from
  html_table
where
  url = 'https://simple.wikipedia.org/wiki/List_of_U.S._states_by_population'
  and base_name = 'wiki'
```

In this example the plugin found one table on the page, and downloaded it as `/home/jon/csv/wiki_0.csv`.

Here is a query of that table.

```
with data as (
  select
    "State" as state,
    "Rankinstates&territories,2010"::int as rank2010,
    "Rankinstates&territories,2019"::int as rank2019,
    replace("Percentchange,20102019[note1]",'%','')::numeric as pct_change,
    replace("PercentofthetotalU.S.population,2018[note3]",'%','')::numeric as pct_of_us_pop
  from
    wiki_0
)
select
	state,
  rank2010,
  rank2019,
  - (rank2019 - rank2010) as rank_change,
  pct_change,
  pct_of_us_pop
from
  data
order by
  pct_change desc
```

```
+---------------+----------+----------+-------------+------------+---------------+
| state         | rank2010 | rank2019 | rank_change | pct_change | pct_of_us_pop |
+---------------+----------+----------+-------------+------------+---------------+
| Utah          | 35       | 30       | 5           | 16.0       | 0.96          |
| Texas         | 2        | 2        | 0           | 15.3       | 8.68          |
| Colorado      | 22       | 21       | 1           | 14.5       | 1.72          |
| NewYork       | 3        | 3        | 0           | 14.2       | 6.44          |
| Nevada        | 36       | 33       | 3           | 14.1       | 0.92          |
| Idaho         | 40       | 39       | 1           | 14.0       | 0.53          |
| Arizona       | 16       | 14       | 2           | 13.9       | 2.17          |
| NorthDakota   | 49       | 48       | 1           | 13.3       | 0.23          |
```


