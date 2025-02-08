package prettyprint

import (
	"fmt"
	"io"
	"iter"
)

type Table[T any] struct {
	columns []string
	values  []func(T) Value
}

func NewTable[T any]() *Table[T] {
	return &Table[T]{}
}

func (t *Table[T]) AddColumn(name string, f func(T) Value) {
	t.columns = append(t.columns, name)
	t.values = append(t.values, f)
}

func (t *Table[T]) Print(cfg TablePrintConfig, all iter.Seq[T]) error {
	w := newWriter(cfg.Output)
	sep := " "

	var count int
	cells := make([][]Value, len(t.columns)) // value for cell - [col][row]
	widths := make([]int, 0, len(t.columns)) // max width of column - [col]
	sublines := make([]int, 0)               // max number of lines of row - [row]

	for _, col := range t.columns {
		widths = append(widths, len(col))
	}

	{
		row := 0
		for item := range all {
			sublines = append(sublines, 1)
			for col, f := range t.values {
				value := f(item)
				cells[col] = append(cells[col], value)
				widths[col] = max(widths[col], value.Width())
				sublines[row] = max(sublines[row], value.Len())
			}
			row += 1
		}
		count = row
	}

	if !cfg.NoHeader {
		for i, name := range t.columns {
			_, _ = fmt.Fprintf(w, "%-*s", widths[i], name)
			if i < len(t.columns)-1 {
				_, _ = fmt.Fprint(w, sep)
			}
		}
		_, _ = fmt.Fprintln(w)
	}

	for row := range count {
		for subrow := range sublines[row] {
			for col := range len(t.columns) {
				value := cells[col][row]
				_, _ = fmt.Fprintf(w, "%-*s", widths[col], value.Get(subrow))
				if col < len(t.columns)-1 {
					_, _ = fmt.Fprint(w, sep)
				}
			}
			_, _ = fmt.Fprintln(w)
		}
	}

	return w.err
}

type TablePrintConfig struct {
	NoHeader bool
	Output   io.Writer
}
