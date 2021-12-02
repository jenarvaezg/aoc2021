package files

import "os"

func ReadInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)

}
