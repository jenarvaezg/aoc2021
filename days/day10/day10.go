package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"fmt"
	"strings"
)

var syntaxErrorScoreMapping = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompleteScoreMapping = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

var parensMatching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func syntaxErrorScore(line string) int {
	var parensStack []rune
	for _, c := range line {
		if v, ok := syntaxErrorScoreMapping[c]; !ok {
			parensStack = append([]rune{c}, parensStack...)
		} else {
			if expectedMatch := parensMatching[c]; parensStack[0] != expectedMatch {
				return v
			}
			parensStack = parensStack[1:]
		}
	}
	return 0
}

func autocompleteScore(line string) int {
	// assume bad syntax filtered out already
	var parensStack []rune
	var score int
	for _, c := range line {
		if _, ok := syntaxErrorScoreMapping[c]; !ok {
			parensStack = append([]rune{c}, parensStack...)
		} else {
			parensStack = parensStack[1:]
		}
	}

	for _, c := range parensStack {
		v := autocompleteScoreMapping[c]
		score = score*5 + v
	}
	return score
}

func splitChunksScores(input []string) (int, []int) {
	var syntaxErrorScoreTotal int
	var autocompleteScores []int
	for _, line := range input {
		if score := syntaxErrorScore(line); score > 0 {
			syntaxErrorScoreTotal += score
		} else {
			autocompleteScores = append(autocompleteScores, autocompleteScore(line))
		}
	}
	return syntaxErrorScoreTotal, autocompleteScores
}

func main() {
	puzzleInput := files.ReadInput()
	lines := strings.Split(puzzleInput, "\n")

	syntaxErrorScore, autocompleteScores := splitChunksScores(lines)
	fmt.Println(syntaxErrorScore)
	fmt.Println(intMath.IntMedian(autocompleteScores...))
}
