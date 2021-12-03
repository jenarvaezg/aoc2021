package conversions

import "strconv"

func MustAtoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func MustAtobin(s string) int {
	val, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(val)
}
