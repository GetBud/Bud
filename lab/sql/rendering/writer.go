package rendering

import "bytes"

// Writer ...
type Writer struct {
	buf bytes.Buffer
}

// NewWriter returns a new Writer instance.
func NewWriter() *Writer {
	return &Writer{}
}

// Write ...
func (w *Writer) Write(part string) {
	w.buf.WriteString(part)
}

// String ...
func (w *Writer) String() string {
	return w.buf.String()
}
