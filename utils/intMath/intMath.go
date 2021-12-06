package math

import "math"

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
