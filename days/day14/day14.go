package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"fmt"
	"math"
	"strings"
)

func parseInput(input string) (string, map[string]string) {
	lines := strings.Split(input, "\n")
	insertions := make(map[string]string)

	for _, l := range lines[2:] {
		splitted := strings.Split(l, " -> ")
		insertions[splitted[0]] = splitted[1]
	}

	return lines[0], insertions
}

func step(template string, insertions map[string]string) string {
	var sb strings.Builder

	//var offset int
	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i : i+2])
		//fmt.Println(pair)
		sb.WriteByte(template[i])
		sb.WriteString(insertions[pair])
	}
	sb.WriteByte(template[len(template)-1])

	return sb.String()
}

func getFrequencyLeastMost(pairs map[string]int) (int, int) {
	counts := make(map[rune]int)
	for k, v := range pairs {
		counts[rune(k[0])] += v
	}

	least, most := math.MaxInt, math.MinInt
	for _, v := range counts {
		most = intMath.IntMax(most, v)
		least = intMath.IntMin(least, v)
	}
	return least, most + 1
}

func stepTimes(template string, rules map[string]string, times int) (int, int) {
	pairs := map[string]int{}

	for i := 0; i < len(template)-1; i++ {
		pairs[string(template[i])+string(template[i+1])]++
	}

	for i := 0; i < times; i++ {
		newPairs := map[string]int{}
		for k, v := range pairs {
			newPairs[string(k[0])+rules[k]] += v
			newPairs[rules[k]+string(k[1])] += v
		}
		pairs = newPairs
	}

	return getFrequencyLeastMost(pairs)
}

func main() {
	puzzleInput := files.ReadInput()
	template, insertions := parseInput(puzzleInput)

	least, most := stepTimes(template, insertions, 10)
	fmt.Println(most-least, template)

	least, most = stepTimes(template, insertions, 40)
	fmt.Println(most - least)
}
