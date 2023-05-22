package module

type Module struct {
	name string
}

func New(name string) *Module {
	return &Module{name: name}
}

func (m *Module) Name() string {
	return m.name
}
