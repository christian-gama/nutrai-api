package value

// Context represents the mail context.
type Context struct {
	Body string `json:"body" faker:"sentence"`
	Head string `json:"head" faker:"sentence"`
}

// NewContext creates a new Context.
func NewContext() *Context {
	return &Context{}
}

// SetBody sets the Body field.
func (c *Context) SetBody(body string) *Context {
	c.Body = body
	return c
}

// SetHead sets the Head field.
func (c *Context) SetHead(head string) *Context {
	c.Head = head
	return c
}
