package gintest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/christian-gama/nutrai-api/internal/auth/app/query"
	"github.com/christian-gama/nutrai-api/internal/auth/app/service"
	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	value "github.com/christian-gama/nutrai-api/internal/auth/domain/value/jwt"
	fake "github.com/christian-gama/nutrai-api/testutils/fake/auth/domain/model/user"
)

// AuthUser returns a user with a valid access token injected in the request header (Authorization).
// It will persist the user in the database.
func AuthUser(request *http.Request) *user.User {
	ctx := context.Background()

	user := fake.User()
	registerOutput, err := service.MakeRegisterHandler().
		Handle(ctx, &service.RegisterInput{
			Email:    user.Email,
			Name:     user.Name,
			Password: user.Password,
		})
	if err != nil {
		panic(err)
	}

	findByEmailOutput, err := query.MakeFindByEmailHandler().
		Handle(ctx, &query.FindByEmailInput{
			Email: user.Email,
		})
	if err != nil {
		panic(err)
	}

	SetAccessToken(request, registerOutput.Access)

	return user.SetID(findByEmailOutput.ID)
}

func SetAccessToken(request *http.Request, token value.Token) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}
