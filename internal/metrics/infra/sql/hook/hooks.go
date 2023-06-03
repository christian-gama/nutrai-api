package hook

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/conn"
	"github.com/christian-gama/nutrai-api/internal/core/infra/sql/hook"
	"github.com/christian-gama/nutrai-api/internal/metrics/infra/metrics/sql"
	"gorm.io/gorm"
)

func Register() {
	hook.Register(conn.GetPsql(), UpdateHook, CreateHook, DeleteHook, QueryHook)
}

func UpdateHook(db *gorm.DB) {
	db.Callback().
		Update().
		After(after("update")).
		Register(register("update"), func(d *gorm.DB) {
			sql.QueriesTotal.Inc()
		})
}

func CreateHook(db *gorm.DB) {
	db.Callback().
		Create().
		After(after("create")).
		Register(register("create"), func(d *gorm.DB) {
			sql.QueriesTotal.Inc()
		})
}

func DeleteHook(db *gorm.DB) {
	db.Callback().
		Delete().
		After(after("delete")).
		Register(register("delete"), func(d *gorm.DB) {
			sql.QueriesTotal.Inc()
		})
}

func QueryHook(db *gorm.DB) {
	db.Callback().
		Query().
		After(after("query")).
		Register(register("query"), func(d *gorm.DB) {
			sql.QueriesTotal.Inc()
		})
}

func after(name string) string {
	return fmt.Sprintf("gorm:%s", name)
}

func register(name string) string {
	return fmt.Sprintf("metric:after_%s", name)
}
