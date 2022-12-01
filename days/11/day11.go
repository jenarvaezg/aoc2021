package main

import (
	"aoc2021/utils/files"
	"aoc2021/utils/grid"
	"fmt"
)

type octopi grid.Grid

func (o *octopi) increaseValue(coord grid.Coordinate, flashed map[grid.Coordinate]bool) {
	if _, ok := flashed[coord]; ok {
		return
	}

	o.Cells[coord] = o.Cells[coord] + 1
	if o.Cells[coord] > 9 {
		flashed[coord] = true
		o.Cells[coord] = 0
		for _, neighboor := range coord.Neighboors(true, o.Height, o.Width) {
			o.increaseValue(neighboor, flashed)
		}
	}
}

func (o *octopi) step() int {
	flashed := make(map[grid.Coordinate]bool)
	for x := 0; x < o.Width; x++ {
		for y := 0; y < o.Height; y++ {
			o.increaseValue(grid.Coordinate{X: x, Y: y}, flashed)
		}
	}

	return len(flashed)
}

func (o octopi) String() string {
	return grid.Grid(o).String()
}

func main() {
	puzzleInput := files.ReadInput()
	octopi := octopi(grid.StringToGrid(puzzleInput))

	var flashed, i int
	for i = 0; i < 100; i++ {
		flashed += octopi.step()
	}
	fmt.Println(flashed)

	for {
		if octopi.step() == octopi.Height*octopi.Width {
			fmt.Println("FOUND IT", i+1)
			break
		}
		i++
	}

}
