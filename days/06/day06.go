package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"fmt"
	"strconv"
	"strings"
)

type fish struct {
	timer int
}

func (f *fish) step() *fish {
	if f.timer == 0 {
		f.timer = 6
		return &fish{timer: 8}
	}
	f.timer--
	return nil
}

func parseFishes(s string) []*fish {
	values := strings.Split(s, ",")
	fishes := make([]*fish, len(values))

	for i, v := range values {
		fishes[i] = &fish{
			timer: conversions.MustAtoi(v),
		}
	}
	return fishes
}

func fishesStr(fishes []*fish) string {
	var sb strings.Builder
	for _, f := range fishes {
		sb.WriteString(strconv.FormatInt(int64(f.timer), 10))
		sb.WriteString(",")
	}

	return sb.String()
}

func simulateSlow(fishes []*fish, days int) []*fish {

	for i := 0; i < days; i++ {
		for _, f := range fishes {
			if newFish := f.step(); newFish != nil {
				fishes = append(fishes, newFish)
			}
		}
	}

	return fishes
}

func simulate(fishes []*fish, days int) int {

	daysMap := make([]int, 9)
	for _, f := range fishes {
		daysMap[f.timer]++
	}

	for i := 0; i < days; i++ {
		first := daysMap[0]
		for i := 0; i < 8; i++ {
			daysMap[i] = daysMap[i+1]
		}
		daysMap[6] += first
		daysMap[8] = first
	}

	sum := 0
	for _, f := range daysMap {
		sum += f
	}

	return sum
}

func main() {
	puzzleInput := files.ReadInput()
	initialFishes := parseFishes(puzzleInput)

	fishes := simulateSlow(initialFishes, 80)
	fmt.Println(len(fishes))

	fmt.Println(simulate(parseFishes(puzzleInput), 256))

}
