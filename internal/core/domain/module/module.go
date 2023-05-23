package module

import "github.com/christian-gama/nutrai-api/internal/core/domain/logger"

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

// String implements the Stringer interface. It returns the module's name.
func (m *Module) String() string {
	return m.name
}

func Initialize(log logger.Logger, fn func() (*Module, func())) {
	module, init := fn()
	log.Infof("Initializing '%s' module", module)
	init()
}
