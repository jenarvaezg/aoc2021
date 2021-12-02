package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	UP      direction = "up"
	DOWN    direction = "down"
	FORWARD direction = "forward"
)

type movement struct {
	direction string
	units     int64
}

type position struct {
	horizontal int64
	depth      int64
}

func readInput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(dat)

}

func parseDirections(s string) []movement {
	lines := strings.Split(s, "\n")
	movements := make([]movement, len(lines))
	for i, line := range lines {
		words := strings.Split(line, " ")
		units, _ := strconv.ParseInt(words[1], 10, 64)
		movements[i] = movement{
			direction: words[0], units: units,
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
	var aim int64
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
	puzzleInput := readInput()
	directions := parseDirections(puzzleInput)
	position := getFinalPosition(directions)
	fmt.Println(position.depth * position.horizontal)

	position2 := getFinalPositionWithAim(directions)
	fmt.Println(position2.depth * position2.horizontal)
}
