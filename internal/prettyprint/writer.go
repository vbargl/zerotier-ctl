package prettyprint

import (
	"io"
	"os"
)

var _ io.Writer = (*writer)(nil)

type writer struct {
	w   io.Writer
	err error
}

func newWriter(w io.Writer) *writer {
	if w == nil {
		w = os.Stdout
	}
	return &writer{w: w}
}

func (w *writer) Write(p []byte) (n int, _ error) {
	if w.err == nil {
		n, w.err = w.w.Write(p)
	}

	return n, nil
}
