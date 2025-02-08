package cmd

import (
	"iter"
	"slices"
)

const sep = "/"

func shift(args []string) (string, []string) {
	if len(args) == 0 {
		return "", args
	}

	return args[0], args[1:]
}

func presuf(arr iter.Seq[string], prefix, suffix string) []string {
	return slices.Collect(func(yield func(string) bool) {
		for item := range arr {
			if !yield(prefix + item + suffix) {
				break
			}
		}
	})
}

func notNil[T any](all iter.Seq[*T]) iter.Seq[*T] {
	return func(yield func(*T) bool) {
		if all == nil {
			return
		}

		for item := range all {
			if item != nil {
				if !yield(item) {
					break
				}
			}
		}
	}
}
