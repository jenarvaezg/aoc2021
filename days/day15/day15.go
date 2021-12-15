package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/grid"
	"aoc2021/utils/intMath"
	"fmt"
	"math"
)

var debugMap = map[grid.Coordinate]bool{
	{X: 9, Y: 8}: true,
	{X: 8, Y: 8}: true,
	{X: 8, Y: 7}: true,
	{X: 8, Y: 6}: true,
	{X: 8, Y: 5}: true,
	{X: 7, Y: 5}: true,
	{X: 7, Y: 4}: true,
	{X: 7, Y: 3}: true,
}

func safestPathTotal(riskMap grid.Grid, current, caller grid.Coordinate, path map[grid.Coordinate]bool, visitedCache map[grid.Coordinate]int) int {
	//fmt.Println("At", current)
	// Path from this cell is known
	if v, ok := visitedCache[current]; ok {
		//fmt.Println("Hey I know this guy", current, v)
		return v
	}

	// reached destination
	if current.X == riskMap.Width-1 && current.Y == riskMap.Height-1 {
		v := riskMap.Cells[current]
		visitedCache[current] = v
		return v
	}

	// This cell is included in path already
	if _, ok := path[current]; ok {
		return math.MaxInt32
	}

	currentValue := riskMap.Cells[current]

	pathCopy := make(map[grid.Coordinate]bool)
	for k, v := range path {
		pathCopy[k] = v
	}
	pathCopy[current] = true

	bestPath := math.MaxInt
	for _, neighbor := range current.Neighboors(false, riskMap.Height, riskMap.Width) {
		bestPath = intMath.IntMin(
			bestPath, safestPathTotal(riskMap, neighbor, current, pathCopy, visitedCache),
		)
	}

	if bestPath > 50000000 {
		fmt.Println("FUCK", current, caller)
	}

	if debugMap[current] {
		fmt.Println("From", current, "the best value is", bestPath, "+", currentValue)
	}
	visitedCache[current] = bestPath + currentValue

	return visitedCache[current]
}

func main() {
	puzzleInput := files.ReadInput()
	riskMap := grid.StringToGrid(puzzleInput)
	start := grid.Coordinate{X: 0, Y: 0}
	end := grid.Coordinate{X: riskMap.Width, Y: riskMap.Width}

	fmt.Println(safestPathTotal(
		riskMap,
		start,
		grid.Coordinate{X: -1, Y: -1},
		map[grid.Coordinate]bool{},
		map[grid.Coordinate]int{end: riskMap.Cells[end]},
	) - riskMap.Cells[start])
}
