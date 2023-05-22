package module

import (
	"github.com/christian-gama/nutrai-api/internal/core/domain/logger"
)

// Module is a struct that represents a module.
type Module struct {
	name string
}

// New creates a new module with the given name.
func New(name string) *Module {
	return &Module{name: name}
}

// Name returns the module's name.
func (m *Module) Name() string {
	return m.name
}

// Init is a function that initializes the module with a logger and a function that contains the
// module's logic.

func (m *Module) Init(log logger.Logger, f func()) {
	log.Infof("Initializing '%s' module", m.name)
	f()
}
