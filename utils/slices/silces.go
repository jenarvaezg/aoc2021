package slices

import (
	"strconv"
	"strings"
)

func StrSliceToIntSlice(str string) []int {
	strs := strings.Split(str, "\n")
	ints := make([]int, len(str))

	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}

func SumIntSlice(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
