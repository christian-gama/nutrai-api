package path

import (
	"os"
	"path/filepath"

	"github.com/christian-gama/nutrai-api/pkg/errutil/errors"
)

// Root searches for the 'go.mod' file from the current working directory upwards.
func Root() string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(errors.InternalServerError("go.mod not found"))
		}
		currentDir = parent
	}

	return currentDir
}
