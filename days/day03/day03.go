package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"fmt"
	"strings"
)

func getMostCommon(lines []string) string {
	result := make([]byte, len(lines[0]))
	for i := range lines[0] {
		counter := make(map[byte]int)
		for _, line := range lines {
			c := line[i]
			v, _ := counter[c]
			counter[c] = v + 1
		}
		if counter['0'] > counter['1'] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}

func getLeastCommon(lines []string) string {
	mostCommon := getMostCommon(lines)
	result := ""

	for _, v := range mostCommon {
		if v == '1' {
			result += "0"
		} else {
			result += "1"
		}
	}

	return result
}

func bitCriteria(lines []string, most bool) string {
	validLines := make(map[string]bool)
	for _, l := range lines {
		validLines[l] = true
	}
	for i := 0; ; i++ {
		var pattern string
		ls := make([]string, len(validLines))
		j := 0
		// Get set values
		for k := range validLines {
			// Only add the character at current offset
			ls[j] = string(k[i])
			j++
		}

		if most {
			pattern = getMostCommon(ls)
		} else {
			pattern = getLeastCommon(ls)
		}

		for l := range validLines {
			if l[i] != pattern[0] {
				delete(validLines, l)
			}
			if len(validLines) == 1 {
				return l
			}
		}
	}

}

func main() {
	puzzleInput := files.ReadInput()
	lines := strings.Split(puzzleInput, "\n")
	gammaStr := getMostCommon(lines)
	gamma := conversions.MustAtobin(gammaStr)
	epsilon := gamma ^ conversions.MustAtobin(strings.Repeat("1", len(gammaStr)))
	fmt.Println(gamma * epsilon)

	oxygen := bitCriteria(lines, true)
	co2 := bitCriteria(lines, false)
	fmt.Println(conversions.MustAtobin(oxygen) * conversions.MustAtobin(co2))

}
