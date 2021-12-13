package intMath

import (
	"math"
	"sort"
)

func IntAbs(v int) int {
	if v > 0 {
		return v
	}

	return -v
}

func IntMax(values ...int) int {
	max := math.MinInt

	for _, v := range values {
		if max < v {
			max = v
		}
	}

	return max
}

func IntMin(values ...int) int {
	min := math.MaxInt

	for _, v := range values {
		if min > v {
			min = v
		}
	}

	return min
}

func IntMedian(values ...int) int {
	sort.Ints(values)

	return values[len(values)/2]
}

func IntMean(values ...int) int {
	sum := IntSum(values...)

	return sum / len(values)
}

func IntSum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func IntProduct(values ...int) int {
	total := 1
	for _, v := range values {
		total *= v
	}
	return total
}
