package controller

import (
	"errors"
	"regexp"
	"strings"
)

// Path is the path to be used by a controller. It shouldn't be called directly, instead use the
// helper JoinPath, which will automatically join the strings using slash as a separator.
type Path string

// JoinPath is a helper to create a new Path from a list of strings. It will automatically join
// the strings using slash as a separator.
// Example: JoinPath("this", "is", "a", "path") -> "/this/is/a/path".
func JoinPath(path ...string) Path {
	for _, p := range path {
		if strings.Contains(p, "/") {
			panic(
				errors.New(
					"path must not contain '/'. The path will be automatically joined using '/'",
				),
			)
		}
	}

	formattedPath := Path("/" + strings.Join(path, "/"))

	// Look for any sequence of slash such as //, ///, ////, etc.
	reg := regexp.MustCompile(`/{2,}`)

	// Replace any sequence of slash with a single slash.
	formattedPath = Path(reg.ReplaceAllString(string(formattedPath), "/"))

	return formattedPath
}

// String returns the path as a string.
func (p Path) String() string {
	return string(p)
}

// Add is a helper to add a new path.
func (p Path) Add(path string) Path {
	return JoinPath(string(p), path)
}
