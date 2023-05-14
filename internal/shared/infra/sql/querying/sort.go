package querying

import (
	"fmt"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/queryer"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Sort is a slice of strings that represents a sort.
type Sort []string

// AddSort returns a new Sort with the given field and order.
func AddSort(field string, isDesc bool) queryer.Sorter {
	return Sort{}.Add(field, isDesc)
}

// Add implements Sorter.
func (s Sort) Add(field string, isDesc bool) queryer.Sorter {
	if isDesc {
		s = append(s, fmt.Sprintf("%s:desc", field))
	} else {
		s = append(s, fmt.Sprintf("%s:asc", field))
	}

	return s
}

// Field implements Sorter.
func (s Sort) Field(idx int) string {
	if s[idx] == "" {
		return ""
	}

	parts := strings.Split(s[idx], ":")
	return parts[0]
}

// Order implements Sorter.
func (s Sort) IsDesc(idx int) bool {
	if s[idx] == "" {
		return false
	}

	parts := strings.Split(s[idx], ":")
	return parts[1] == "desc"
}

// Slice implements Sorter.
func (s Sort) Slice() []string {
	return s
}

// SortScope returns a GORM scope that applies sorting to the query.
func SortScope(sorter queryer.Sorter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if sorter == nil {
			return db
		}

		for i := range sorter.Slice() {
			if sorter.Field(i) == "" {
				continue
			}

			columName := schema.NamingStrategy{}.ColumnName("", sorter.Field(i))

			db = db.Order(clause.OrderByColumn{
				Column: clause.Column{Name: columName},
				Desc:   sorter.IsDesc(i),
			})
		}

		return db
	}
}
