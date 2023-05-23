package gintest

import (
	"encoding/json"

	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
)

func GetBody(body string) *response.Body {
	var b *response.Body
	json.Unmarshal([]byte(body), &b)
	return b
}
