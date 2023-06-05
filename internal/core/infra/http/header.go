package http

import (
	"fmt"
	"net/http"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

func CheckAuthorizationHeader(request *http.Request, secret string) (string, error) {
	authorization := request.Header.Get("Authorization")
	if authorization != fmt.Sprintf("Bearer %s", secret) {
		return "", errors.Unauthorized("you are not authorized to access this resource")
	}

	return authorization, nil
}
