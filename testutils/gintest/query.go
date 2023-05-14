package gintest

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/sql/querying"
)

func BuildScopeQuery(filter querying.Filter, sortBy querying.Sort) string {
	queries := []string{}

	filterStr := ""
	if len(filter) > 0 {
		filterStr = fmt.Sprintf("filter=%s", strings.Join(filter, ","))
	}
	queries = append(queries, filterStr)

	sortByStr := ""
	if len(sortBy) > 0 {
		sortByStr = fmt.Sprintf("sort=%s", strings.Join(sortBy, ","))
	}
	queries = append(queries, sortByStr)

	return strings.Join(queries, "&")
}

func BuildQuery(field string, value any) string {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(value)
		if s.Len() == 0 {
			return ""
		}

		var values []string
		for i := 0; i < s.Len(); i++ {
			values = append(values, fmt.Sprintf("%v", s.Index(i).Interface()))
		}

		result := ""
		for i := 1; i < len(values); i++ {
			result += fmt.Sprintf("&%s=%s", field, values[i])
		}

		result = strings.TrimPrefix(result, "&")

		return result

	default:
		result := fmt.Sprintf("%s=%v", field, value)
		return result
	}
}
