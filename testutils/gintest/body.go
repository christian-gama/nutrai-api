package gintest

import (
	"encoding/json"
)

func GetBody(body string) map[string]any {
	var b map[string]any
	json.Unmarshal([]byte(body), &b)
	return b
}
