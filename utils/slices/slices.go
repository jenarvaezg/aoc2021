package slices

import (
	"strconv"
)

func StrSliceToIntSlice(strs []string) []int {
	ints := make([]int, len(strs))

	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}
