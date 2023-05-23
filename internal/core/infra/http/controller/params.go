package controller

import (
	"fmt"
	"strings"
)

// Params is the list of params to be used by a controller.
// Example: Params{"id"} -> localhost:8080/resource/:id.
type Params []string

// Params is a helper to create a new Params. It will automatically add a colon as a prefix.
// Example: AddParams("id") -> :id.
func AddParams(params string) Params {
	params = strings.TrimSpace(params)
	params = strings.TrimPrefix(params, ":")
	return Params{params}
}

// Add is a helper to add a new param.
// Example: AddParams("id").Add("name").
func (p Params) Add(param string) Params {
	return append(p, param)
}

// Slice returns the params as a slice of strings.
func (p Params) Slice() []string {
	return p
}

// ToPath returns the params as a path.
// Example: AddParams("id").ToPath("/resource") -> "/resource/:id".
func (p Params) ToPath(path Path) Path {
	return Path(
		fmt.Sprintf("%s/:%s", strings.TrimSuffix(path.String(), "/"), strings.Join(p, "/:")),
	)
}
