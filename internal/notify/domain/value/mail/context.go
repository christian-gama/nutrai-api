package value

type Context struct {
	Body string
	Head string
}

// NewContext creates a new Context.
func NewContext(body string, head string) *Context {
	return &Context{
		Body: body,
		Head: head,
	}
}
