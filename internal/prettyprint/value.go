package prettyprint

import (
	"iter"
	"slices"
)

type Value struct {
	single *string
	slice  *[]string
}

func Single(value string) Value {
	return Value{single: &value}
}

func Slice(all []string) Value {
	return Value{slice: &all}
}

func Seq(all iter.Seq[string]) Value {
	allSlice := slices.AppendSeq([]string(nil), all)
	return Value{slice: &allSlice}
}

func (v *Value) Get(row int) string {
	switch {
	case v.single != nil && row == 0:
		return *v.single

	case v.slice != nil && row < len(*v.slice):
		return (*v.slice)[row]

	default:
		return ""
	}
}

func (v *Value) Width() int {
	switch {
	case v.single != nil:
		return len(*v.single)

	case v.slice != nil:
		max := 0
		for _, s := range *v.slice {
			if len(s) > max {
				max = len(s)
			}
		}
		return max

	default:
		return 0
	}
}

func (v *Value) Len() int {
	switch {
	case v.single != nil:
		return 1

	case v.slice != nil:
		return len(*v.slice)

	default:
		return 0
	}
}
