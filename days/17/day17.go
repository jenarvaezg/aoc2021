package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"aoc2021/utils/grid"
	"aoc2021/utils/intMath"
	"fmt"
	"regexp"
)

type bounds struct {
	xMin, xMax int
	yMin, yMax int
}

var re = regexp.MustCompile(`(-?\d+)`) //target area: x=(.?\d*)\.\.(.?\d*), y=(.?\d*)\.\.(.?\d*)`)

func parseInput(input string) bounds {
	matches := re.FindAllStringSubmatch(input, -1)
	return bounds{
		conversions.MustAtoi(matches[0][0]), conversions.MustAtoi(matches[1][0]),
		conversions.MustAtoi(matches[2][0]), conversions.MustAtoi(matches[3][0]),
	}
}

func isValidLaunch(position, delta grid.Coordinate, area bounds) bool {
	newPosition := position.Sum(delta)
	if newPosition.X <= area.xMax && newPosition.X >= area.xMin && newPosition.Y <= area.yMax && newPosition.Y >= area.yMin {
		return true
	}
	if (delta.X >= 0 && newPosition.X > area.xMax) || (delta.X <= 0 && newPosition.X < area.xMin) {
		return false
	}

	if newPosition.Y < area.yMin {
		return false
	}

	var newDeltaX, newDeltaY int
	if delta.X > 0 {
		newDeltaX = delta.X - 1
	} else if delta.X < 0 {
		newDeltaX = delta.X + 1
	}
	newDeltaY = delta.Y - 1

	return isValidLaunch(newPosition, grid.Coordinate{X: newDeltaX, Y: newDeltaY}, area)
}

func possibleShots(area bounds) []grid.Coordinate {
	var validShots []grid.Coordinate
	for x := -500; x < 500; x++ {
		for y := -500; y < 500; y++ {
			coord := grid.Coordinate{X: x, Y: y}
			if isValidLaunch(grid.Coordinate{X: 0, Y: 0}, coord, area) {
				validShots = append(validShots, coord)
			}
		}
	}
	return validShots
}

func highestY(shots []grid.Coordinate) int {
	var ys []int
	for _, coord := range shots {
		ys = append(ys, coord.Y)
	}
	maxY := intMath.IntMax(ys...)
	return maxY * (maxY + 1) / 2
}

func main() {
	puzzleInput := files.ReadInput()
	area := parseInput(puzzleInput)
	shots := possibleShots(area)

	fmt.Println(highestY(shots))

	fmt.Println(len(shots))

}
