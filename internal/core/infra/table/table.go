package table

import (
	"fmt"

	"github.com/christian-gama/nutrai-api/internal/core/domain/module"
)

// Name returns a formatted name for a table.
func Name(module *module.Module, table string) string {
	return fmt.Sprintf("%s_%s", module, table)
}
