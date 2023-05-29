# Queries

The Querying Feature of the Nutrai API allows users to make complex queries to the database with ease and efficiency. It includes a filter, sort, preload and pagination systems.

## Filtering

The `Filter` system allows for complex querying of the database using different operators. A new filter can be created using `AddFilter(field string, operator string, value any) queryer.Filterer`.

For example, `AddFilter("name", "eq", "John")` would create a filter that looks for records where the name field equals "John".

Operators allowed in filters are:

- `eq` (equal)
- `neq` (not equal)
- `gt` (greater than)
- `gte` (greater than or equal)
- `lt` (less than)
- `lte` (less than or equal)
- `like` (like, used for string matching)
- `in` (in, used for array matching)

## Sorting

The `Sort` system allows for ordering of the returned records. A new sorter can be created using `AddSort(field string, isDesc bool) queryer.Sorter`.

For example, `AddSort("name", true)` would create a sorter that sorts the records by the name field in descending order.

## Preloading

The `Preload` system allows for preloading associated records. A new preloader can be created using `AddPreload(field string) queryer.Preloader`.

For example, `AddPreload("orders")` would preload the orders for each returned record.

## Pagination

The `Pagination` system allows for limiting and offsetting the returned records. A new paginator can be created using the `Pagination` struct and its methods.

For example, a `Pagination{Page: 2, Limit: 20}` would limit the returned records to 20 and offset them by 20 (i.e. return the second page of records).

## Validators

The API includes validators for filters to ensure that only valid queries are processed. These validators are used in the application input, using the go-playground/validator package, passing the validator as a struct tag.

## How to Use
```go
type Input struct {
	querying.Pagination `faker:"-"`
	querying.Filter     `form:"filter" validate:"query,filter=name email"`
	querying.Sort       `form:"sort" validate:"query,sort=id name"`
	querying.Preload    `form:"preload" validate:"query,preload=user"`
}
```

In the URL, the query parameters would look like this:

```
GET .../resource?filter=field=name,op=like,value=john&sort=id:desc,name:asc&preload=user&limit=20&page=2
```

Splitting the query above, we have:
```
filter=field=name,op=like,value=john
sort=id:desc,name:asc
preload=user
limit=20
page=2
```

By passing the params in the query, the API will automatically parse the query and create the respective filters, sorters, preloaders and paginator, if you're using a [Manager](../internal/core/infra/sql/manager/manager.go). If you're manually creating a query, you must use the gin Scope methods to get the query params and pass them to the respective functions.