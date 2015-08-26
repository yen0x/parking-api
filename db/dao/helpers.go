package dao

import (
	"strings"
)

// TODO(remy): comment
func insertFields(tableName string, fields string) string {
	return strings.Replace(fields, "\""+tableName+"\".", "", -1)
}
