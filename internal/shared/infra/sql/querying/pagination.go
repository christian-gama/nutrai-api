package querying

import (
	"github.com/christian-gama/nutrai-api/internal/shared/domain/queryer"
	"gorm.io/gorm"
)

// Pagination is a struct used to paginate queries.
type Pagination struct {
	Page  int `form:"page"  validate:"omitempty,min=1"`
	Limit int `form:"limit" validate:"omitempty,max=100,min=1"`
}

// GetLimit implements queryer.Paginator.
func (p *Pagination) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Limit > 100 {
		p.Limit = 100
	}

	return p.Limit
}

// GetOffset implements queryer.Paginator.
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetPage implements queryer.Paginator.
func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}

	return p.Page
}

// SetLimit implements queryer.Paginator.
func (p *Pagination) SetLimit(limit int) queryer.Paginator {
	p.Limit = limit
	return p
}

// SetPage implements queryer.Paginator.
func (p *Pagination) SetPage(page int) queryer.Paginator {
	p.Page = page
	return p
}

// PaginationScope is a function that can be used as a GORM scope to paginate queries.
func PaginationScope(paginator queryer.Paginator) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paginator == nil {
			return db
		}

		return db.
			Offset(paginator.GetOffset()).
			Limit(paginator.GetLimit()).
			Select("*, COUNT(*) OVER() AS total")
	}
}
