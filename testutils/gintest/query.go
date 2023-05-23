package gintest

import (
	"fmt"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
)

type QueryOption struct {
	Filter  querying.Filter
	Sort    querying.Sort
	Preload querying.Preload
}

func FilterOption(filter querying.Filter) QueryOption {
	return QueryOption{Filter: filter}
}

func SortOption(sort querying.Sort) QueryOption {
	return QueryOption{Sort: sort}
}

func PreloadOption(preload querying.Preload) QueryOption {
	return QueryOption{Preload: preload}
}

func BuildScopeQuery(options ...QueryOption) string {
	if len(options) == 0 {
		return ""
	}

	queries := []string{}
	var filter querying.Filter
	var sortBy querying.Sort
	var preload querying.Preload

	for _, option := range options {
		if option.Filter != nil {
			filter = append(filter, option.Filter...)
		}

		if option.Sort != nil {
			sortBy = append(sortBy, option.Sort...)
		}

		if option.Preload != nil {
			preload = append(preload, option.Preload...)
		}
	}

	filterStr := ""
	if len(filter) > 0 {
		filterStr = fmt.Sprintf("filter=%s", strings.Join(filter, "&filter="))
	}
	queries = append(queries, filterStr)

	sortByStr := ""
	if len(sortBy) > 0 {
		sortByStr = fmt.Sprintf("sort=%s", strings.Join(sortBy, "&sort="))
	}
	queries = append(queries, sortByStr)

	preloadStr := ""
	if len(preload) > 0 {
		preloadStr = fmt.Sprintf("preload=%s", strings.Join(preload, "&preload="))
	}
	queries = append(queries, preloadStr)

	return strings.Join(queries, "&")
}
