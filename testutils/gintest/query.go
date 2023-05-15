package gintest

import (
	"fmt"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/querying"
)

func BuildScopeQuery(filter querying.Filter, sortBy querying.Sort, preload querying.Preload) string {
	queries := []string{}

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
