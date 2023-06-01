package sqlerr

import (
	"strings"

	"github.com/iancoleman/strcase"
)

// getFriendlyTableName returns a friendly table name.
//
// A table is composed by the following parts: namespace_tablename
//
// This function will return the last part of the table name, which is the table name itself,
// removing the namespace part.
func getFriendlyTableName(tableName string) string {
	if strings.Contains(tableName, "_") {
		splittedTableName := strings.Split(tableName, "_")
		return strcase.ToLowerCamel(splittedTableName[len(splittedTableName)-1])
	}

	return strcase.ToLowerCamel(tableName)
}
