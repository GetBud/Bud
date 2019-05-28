package builder

import "bytes"

// Context is a type that handles the state related to how a query is built; for example,
// controlling the parameters used for the dialect of SQL being used.
type Context struct {
	buf  bytes.Buffer
	args []interface{}
}

// NewContext returns a new Context instance.
func NewContext() *Context {
	return &Context{}
}

// Args ...
func (c *Context) Args() []interface{} {
	return c.args
}

// Write ...
func (c *Context) Write(s string) {
	c.buf.WriteString(s)
}

// String ...
func (c *Context) String() string {
	return c.buf.String()
}
