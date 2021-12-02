package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)

}

func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func windowIncreaseCount(sweep []int, windowSize int) int {
	val := sum(sweep[:windowSize])
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

func inputToIntSlice(input string) []int {
	strs := strings.Split(input, "\n")
	ints := make([]int, len(input))

	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}

func main() {
	puzzleInput := readInput()
	ints := inputToIntSlice(puzzleInput)

	fmt.Println(windowIncreaseCount(ints, 1))
	fmt.Println(windowIncreaseCount(ints, 3))
}
