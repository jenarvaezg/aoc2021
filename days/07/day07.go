package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"aoc2021/utils/slices"
	"fmt"
	"math"
	"strings"
)

func getOptimalPoint(values []int) int {
	squareSum := 0
	for _, v := range values {
		squareSum += v * v
	}

	return int(math.Round(math.Sqrt(float64(squareSum)) / float64(len(values))))
}

func getTotalMovements(values []int, point int) int {
	total := 0
	for _, v := range values {
		total += intMath.IntAbs(v - point)
	}
	return total
}

func getFuelRaisingCost(values []int, point int) int {
	total := 0
	for _, v := range values {
		diff := intMath.IntAbs(v - point)
		total += int(float64(diff*(diff+1)) / 2)
	}
	return total
}

func main() {
	puzzleInput := files.ReadInput()
	values := slices.StrSliceToIntSlice(strings.Split(puzzleInput, ","))

	fmt.Println(getTotalMovements(values, intMath.IntMedian(values...)))

	fmt.Println(getFuelRaisingCost(values, intMath.IntMean(values...)))

}
