package gintest

import (
	"fmt"
	gohttp "net/http"
	"net/http/httptest"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/auth/domain/model/user"
	"github.com/christian-gama/nutrai-api/internal/auth/infra/store"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/controller"
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/response"
	"github.com/christian-gama/nutrai-api/testutils/httputil"
	"github.com/gin-gonic/gin"
)

func MustRequest(handler controller.Controller, opt Option) (ctx *gin.Context) {
	ctx, _ = MustRequestWithBody(handler, opt)
	return ctx
}

func MustRequestWithBody(
	handler controller.Controller,
	opt Option,
) (ctx *gin.Context, body *response.Body) {
	ctx, r, writer := createTestContext()

	handlerPath := handler.Path()

	if len(handler.Params()) > 0 {
		handlerPath = controller.Path(
			fmt.Sprintf("%s/:%s", handlerPath, strings.Join(handler.Params(), "/:")),
		)
	}

	r.Handle(handler.Method().String(), handlerPath.String(), func(ctx *gin.Context) {
		if opt.CurrentUser != nil {
			store.SetUser(ctx, opt.CurrentUser)
		}

		handler.Handle(ctx)
	})

	var err error
	path := handler.Path()
	if len(opt.Params) > 0 {
		path = controller.Path(strings.TrimSuffix(path.String(), "/"))
		path = controller.Path(fmt.Sprintf("%s/%s", path, strings.Join(opt.Params, "/")))
	}

	path += "?page=1&limit=100"
	if opt.Queries != "" {
		path = controller.Path(fmt.Sprintf("%s&%s", path, opt.Queries))
	}

	ctx.Request, err = gohttp.NewRequest(
		string(handler.Method()),
		path.String(),
		strings.NewReader(httputil.Stringify(opt.Data)),
	)
	if err != nil {
		panic(err)
	}

	r.ServeHTTP(ctx.Writer, ctx.Request)
	if ctx.Writer.Status() == 404 {
		panic(fmt.Errorf("path %s not found", path))
	}

	return ctx, GetBody(writer.Body.String())
}

type Option struct {
	// Data is the input of the request.
	Data any

	// Params are the keys of the path params. E.g:
	//
	// []string{"id", "name"}
	//
	// It would be equivalent to /items/:id/:name.
	Params []string

	// Queries are the query params. E.g
	//
	// "id=1&name=foo"
	//
	// It would be equivalent to /items?id=1&name=foo
	Queries string

	// CurrentUser is the user.User that will be set in the context of the request to simulate an
	// authenticated user.
	CurrentUser *user.User
}

func createTestContext() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	writer := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(writer)

	return ctx, r, writer
}
