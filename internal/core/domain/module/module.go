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

// String implements the fmt.Stringer interface.
func (m Module) String() string {
	return m.name
}

// Initialize performs the initialization of a module, logging the process and calling the given
// function.
func Initialize(log logger.Logger, callback func() (*Module, func())) {
	module, init := callback()
	log.Infof("Initializing '%s' module", module)
	init()
}
