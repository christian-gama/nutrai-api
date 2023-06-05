package routes

import (
	"github.com/christian-gama/nutrai-api/internal/core/infra/http/router"
	"github.com/christian-gama/nutrai-api/pkg/slice"
)

// Api initializes a new instance of routes with a given group with a "api" prefix.
func Api(group ...string) *routes {
	return &routes{
		group: router.Router.Group("api").Group(slice.FirstElementOrDefault(group)),
	}
}

// Root initializes a new instance of routes with a given group and a root path (no prefix).
func Root(group ...string) *routes {
	return &routes{
		group: router.Router.Group(slice.FirstElementOrDefault(group)),
	}
}

// Internal initializes a new instance of routes with a given group with a "internal" prefix.
func Internal(group ...string) *routes {
	return &routes{
		group: router.Router.Group("internal").Group(slice.FirstElementOrDefault(group)),
	}
}
