package builder

import (
	"bytes"
	"strings"
)

// Context is a type that handles the state related to how a query is built; for example,
// controlling the parameters used for the dialect of SQL being used.
type Context struct {
	buf   bytes.Buffer
	args  []interface{}
	depth int
}

// NewContext returns a new Context instance.
func NewContext() *Context {
	return &Context{}
}

// AddArgs ...
func (c *Context) AddArgs(args ...interface{}) {
	c.args = append(c.args, args...)
}

// Args ...
func (c *Context) Args() []interface{} {
	return c.args
}

// Depth ..
func (c *Context) Depth(n int) int {
	c.depth += n
	return c.depth
}

// Indent ...
func (c *Context) Indent() string {
	return strings.Repeat("  ", c.depth)
}

// Write ...
func (c *Context) Write(s string) {
	c.buf.WriteString(s)
}

// String ...
func (c *Context) String() string {
	return c.buf.String()
}
