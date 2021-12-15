package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/grid"
	"fmt"

	"github.com/beefsack/go-astar"
)

type tile struct {
	coord grid.Coordinate
	grid  *grid.Grid
}

func (t tile) PathNeighbors() []astar.Pather {
	var neighboors []astar.Pather
	for _, neighboor := range t.coord.Neighboors(false, t.grid.Height, t.grid.Width) {
		neighboors = append(neighboors, tile{neighboor, t.grid})
	}
	return neighboors
}

func (t tile) PathNeighborCost(to astar.Pather) float64 {
	toTile := to.(tile)
	v := t.grid.Cells[toTile.coord]

	return float64(v)
}

func (t tile) PathEstimatedCost(to astar.Pather) float64 {
	toTile := to.(tile)
	return t.coord.ManhattanDistance(toTile.coord)
}

func expandGrid(riskMap grid.Grid, times int) grid.Grid {
	newCells := map[grid.Coordinate]int{}
	for baseCoord, v := range riskMap.Cells {
		for xTimes := 0; xTimes < times; xTimes++ {
			for yTimes := 0; yTimes < times; yTimes++ {
				coord := grid.Coordinate{
					X: baseCoord.X + (riskMap.Width * xTimes),
					Y: baseCoord.Y + (riskMap.Height * yTimes),
				}
				newVal := v + xTimes + yTimes
				if newVal > 9 {
					newCells[coord] = (newVal % 9)
				} else {
					newCells[coord] = newVal
				}

			}
		}

	}

	return grid.Grid{Cells: newCells, Height: riskMap.Height * times, Width: riskMap.Width * times}
}

func main() {
	puzzleInput := files.ReadInput()

	riskMap := grid.StringToGrid(puzzleInput)
	start, end := grid.Coordinate{X: 0, Y: 0}, grid.Coordinate{X: riskMap.Width - 1, Y: riskMap.Height - 1}
	_, distance, _ := astar.Path(tile{coord: start, grid: &riskMap}, tile{coord: end, grid: &riskMap})
	fmt.Println(distance)

	expandedGrid := expandGrid(riskMap, 5)
	start, end = grid.Coordinate{X: 0, Y: 0}, grid.Coordinate{X: expandedGrid.Width - 1, Y: expandedGrid.Height - 1}
	_, distance, _ = astar.Path(tile{coord: start, grid: &expandedGrid}, tile{coord: end, grid: &expandedGrid})
	fmt.Println(distance)

}
