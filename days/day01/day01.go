package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"aoc2021/utils/slices"
	"fmt"
	"strings"
)

func windowIncreaseCount(sweep []int, windowSize int) int {
	val := intMath.IntSum(sweep[:windowSize]...)
	increases := 0

	for i := windowSize; i < len(sweep); i++ {
		newVal := val - sweep[i-windowSize] + sweep[i]
		if newVal > val {
			increases++
		}
		val = newVal
	}
	return increases
}

func main() {
	puzzleInput := files.ReadInput()
	ints := slices.StrSliceToIntSlice(strings.Split(puzzleInput, "\n"))

	fmt.Println(windowIncreaseCount(ints, 1))
	fmt.Println(windowIncreaseCount(ints, 3))
}
