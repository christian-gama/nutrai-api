package querying

import (
	"strings"

	"github.com/christian-gama/nutrai-api/internal/shared/domain/querying"
	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

// Preload is a slice of strings that represents a Preload.
type Preload []string

// AddPreload returns a new Preload with the given field and order.
func AddPreload(field string) querying.Preloader {
	return Preload{}.Add(field)
}

// Add implements Preloader.
func (s Preload) Add(field string) querying.Preloader {
	s = append(s, field)

	return s
}

// Slice implements Preloader.
func (s Preload) Slice() []string {
	return s
}

// PreloadScope returns a GORM scope that applies preloading to the query.
func PreloadScope(preload querying.Preloader) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if preload != nil {
			for _, field := range preload.Slice() {
				if field == "" {
					continue
				}

				splittedByDot := strings.Split(field, ".")
				output := ""

				for i, s := range splittedByDot {
					if i != 0 {
						output += "."
					}

					output += strcase.ToCamel(s)
				}

				db = db.Preload(output)
			}
		}

		return db
	}
}
