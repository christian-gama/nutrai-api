package http

import (
	"net/http"
	"strings"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

func GetAuthorizationHeader(request *http.Request) (string, error) {
	authorization := request.Header.Get("Authorization")
	if authorization == "" {
		return "", errors.Unauthorized("you are not authorized to access this resource")
	}

	parts := strings.Split(authorization, " ")
	if len(parts) != 2 {
		return "", errors.Unauthorized("you are not authorized to access this resource")
	}

	return parts[1], nil
}
