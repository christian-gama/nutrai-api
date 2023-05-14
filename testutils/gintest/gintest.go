package gintest

import (
	"fmt"
	gohttp "net/http"
	"net/http/httptest"
	"strings"

	"github.com/christian-gama/nutrai-api/internal/shared/infra/http"
	"github.com/christian-gama/nutrai-api/testutils/httputil"
	"github.com/gin-gonic/gin"
)

func MustRequest(handler http.Controller, opt Option) (ctx *gin.Context) {
	ctx, _ = MustRequestWithBody(handler, opt)
	return ctx
}

func MustRequestWithBody(
	handler http.Controller,
	opt Option,
) (ctx *gin.Context, body *http.ResponseBody) {
	ctx, r, writer := createTestContext()
	handlerPath := handler.Path()
	if len(handler.Params()) > 0 {
		handlerPath = fmt.Sprintf("%s/:%s", handlerPath, strings.Join(handler.Params(), "/:"))
	}

	r.Handle(handler.Method(), handlerPath, func(ctx *gin.Context) {
		handler.Handle(ctx)
	})

	var err error
	path := handler.Path()
	if len(opt.Params) > 0 {
		path = strings.TrimSuffix(path, "/")
		path = fmt.Sprintf("%s/%s", path, strings.Join(opt.Params, "/"))
	}

	path += "?page=1&limit=100"
	if opt.Queries != "" {
		path = fmt.Sprintf("%s&%s", path, opt.Queries)
	}

	ctx.Request, err = gohttp.NewRequest(
		handler.Method(),
		path,
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
	// Data is the payload of the request.
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
}

func createTestContext() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	writer := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(writer)

	return ctx, r, writer
}
