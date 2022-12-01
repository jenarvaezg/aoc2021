package main

import (
	"aoc2021/utils/files"
	"fmt"
	"sort"
	"strings"
)

func orderSliceValues(values []string) []string {
	for i, v := range values {
		sorted := strings.Split(v, "")
		sort.Strings(sorted)
		values[i] = strings.Join(sorted, "")
	}
	return values
}

func parseInput(input string) ([][]string, [][]string) {
	lines := strings.Split(input, "\n")
	inputs, outputs := make([][]string, len(lines)), make([][]string, len(lines))
	for i, line := range lines {
		splitted := strings.Split(line, " ")
		inputs[i] = orderSliceValues(splitted[:len(splitted)-5])
		outputs[i] = orderSliceValues(splitted[len(splitted)-4:])
	}

	return inputs, outputs
}

func countEasy(values [][]string) int {
	total := 0
	for _, line := range values {
		for _, word := range line {
			if l := len(word); l != 6 && l != 5 {
				total++
			}
		}
	}
	return total
}

func signalsToInt(signal []string, mapping map[string]int) int {
	total := 0
	for _, word := range signal {
		total *= 10
		result, ok := mapping[word]
		if !ok {
			panic("Couldn't find " + word)
		}
		total += result
	}

	return total
}

func stringSetDifference(s, substr string) int {
	missingCount := 0
	for _, c := range substr {
		if !strings.Contains(s, string(c)) {
			missingCount++
		}
	}
	return missingCount
}

func obtainMapping(signal []string) map[string]int {
	signalPositions := make([]string, 10)
	pendingSignals := make([]string, len(signal))
	copy(pendingSignals, signal)

	// First find easy values (1, 4, 7 and 8)
	for i := 0; i < len(pendingSignals); i++ {
		word := 	[i]
		found := true

		if l := len(word); l == 2 {
			signalPositions[1] = word
		} else if l == 3 {
			signalPositions[7] = word
		} else if l == 4 {
			signalPositions[4] = word
		} else if l == 7 {
			signalPositions[8] = word
		} else {
			found = false
		}

		if found {
			pendingSignals = append(pendingSignals[:i], pendingSignals[i+1:]...)
			i--
		}
	}

	// 2, 3, 5, 6, 9 and 0 missing
	// find 3: word with len = 5 and having 7 as substring
	for i := 0; i < len(pendingSignals); i++ {
		word := pendingSignals[i]
		if len(word) == 5 && stringSetDifference(word, signalPositions[7]) == 0 {
			signalPositions[3] = word
			pendingSignals = append(pendingSignals[:i], pendingSignals[i+1:]...)
			break
		}
	}

	// 2, 5, 6, 9 and 0 missing
	// find 2, 6 and 9: words with len = 6 and. 9 has 3 as substring, 0 has 7 as substring, other is 6
	for i := 0; i < len(pendingSignals); i++ {
		word := pendingSignals[i]
		if len(word) == 6 {
			if stringSetDifference(word, signalPositions[3]) == 0 {
				signalPositions[9] = word
			} else if stringSetDifference(word, signalPositions[7]) == 0 {
				signalPositions[0] = word
			} else {
				signalPositions[6] = word
			}
			pendingSignals = append(pendingSignals[:i], pendingSignals[i+1:]...)
			i--
		}
	}

	// 2, and 5 missing
	// 5 is substr of 9, other is 2
	for i := 0; i < len(pendingSignals); i++ {
		word := pendingSignals[i]
		if stringSetDifference(word, signalPositions[9]) == 1 {
			signalPositions[5] = word
		} else {
			signalPositions[2] = word
		}
	}

	// flip position slice into map
	mapping := make(map[string]int)
	for i, v := range signalPositions {
		if v != "" {
			mapping[v] = i
		}

	}

	return mapping

}

func displaysOutputsSum(inputs [][]string, outputs [][]string) int {
	total := 0
	for i := range inputs {
		mapping := obtainMapping(inputs[i])
		total += signalsToInt(outputs[i], mapping)
	}
	return total
}

func main() {
	puzzleInput := files.ReadInput()
	displayInputs, displayOutputs := parseInput(puzzleInput)

	fmt.Println(countEasy(displayOutputs))

	fmt.Println(displaysOutputsSum(displayInputs, displayOutputs))
}
