package convert

import (
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var Converters = []copier.TypeConverter{
	{
		SrcType: gorm.DeletedAt{},
		DstType: time.Time{},
		Fn: func(src any) (any, error) {
			return src.(gorm.DeletedAt).Time, nil
		},
	},
	{
		SrcType: time.Time{},
		DstType: gorm.DeletedAt{},
		Fn: func(src any) (any, error) {
			return gorm.DeletedAt{Time: src.(time.Time)}, nil
		},
	},
}
