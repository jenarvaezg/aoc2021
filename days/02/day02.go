package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"fmt"
	"strings"
)

const (
	UP      = "up"
	DOWN    = "down"
	FORWARD = "forward"
)

type movement struct {
	direction string
	units     int
}

type position struct {
	horizontal int
	depth      int
}

func parseDirections(s string) []movement {
	lines := strings.Split(s, "\n")
	movements := make([]movement, len(lines))
	for i, line := range lines {
		words := strings.Split(line, " ")
		movements[i] = movement{
			direction: words[0], units: conversions.MustAtoi(words[1]),
		}
	}
	return movements

}

func getFinalPosition(movements []movement) position {
	currentPosition := position{}
	for _, m := range movements {
		switch m.direction {
		case string(UP):
			currentPosition.depth -= m.units
		case string(DOWN):
			currentPosition.depth += m.units
		case string(FORWARD):
			currentPosition.horizontal += m.units
		}
	}

	return currentPosition
}

func getFinalPositionWithAim(movements []movement) position {
	var aim int
	currentPosition := position{}
	for _, m := range movements {
		switch m.direction {
		case string(UP):
			aim -= m.units
		case string(DOWN):
			aim += m.units
		case string(FORWARD):
			currentPosition.horizontal += m.units
			currentPosition.depth += aim * m.units
		}
	}

	return currentPosition
}

func main() {
	puzzleInput := files.ReadInput()
	directions := parseDirections(puzzleInput)
	position := getFinalPosition(directions)
	fmt.Println(position.depth * position.horizontal)

	position2 := getFinalPositionWithAim(directions)
	fmt.Println(position2.depth * position2.horizontal)
}
