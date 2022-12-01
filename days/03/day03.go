package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"fmt"
	"strings"
)

func getMostCommon(lines []string) string {
	var result string
	for i := range lines[0] {
		result += getLeastCommonAt(lines, i)
	}

	return result
}

func getMostCommonAt(lines []string, at int) string {
	counter := make(map[byte]int)
	for _, line := range lines {
		c := line[at]
		v, _ := counter[c]
		counter[c] = v + 1
	}
	if counter['0'] > counter['1'] {
		return "0"
	} else {
		return "1"
	}
}

func getLeastCommonAt(lines []string, at int) string {
	if mostCommon := getMostCommonAt(lines, at); mostCommon == "1" {
		return "0"
	} else {
		return "1"
	}
}

func bitCriteria(input []string, patternFn func([]string, int) string) string {
	validLines := make([]string, len(input))
	copy(validLines, input)

	for bit := 0; len(validLines) > 1; bit++ {
		pattern := patternFn(validLines, bit)
		for i := 0; i < len(validLines); i++ {
			if validLines[i][bit] != pattern[0] {
				validLines = append(validLines[:i], validLines[i+1:]...)
				i--
			}
		}
	}

	return validLines[0]
}

func main() {
	puzzleInput := files.ReadInput()
	lines := strings.Split(puzzleInput, "\n")
	gammaStr := getMostCommon(lines)
	gamma := conversions.MustAtobin(gammaStr)
	epsilon := gamma ^ conversions.MustAtobin(strings.Repeat("1", len(gammaStr)))
	fmt.Println(gamma * epsilon)

	oxygen := bitCriteria(lines, getMostCommonAt)
	co2 := bitCriteria(lines, getLeastCommonAt)
	fmt.Println(conversions.MustAtobin(oxygen) * conversions.MustAtobin(co2))

}
