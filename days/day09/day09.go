package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"fmt"
	"sort"
	"strings"
)

type coordinate struct {
	x, y int
}

type heightMap struct {
	grid          map[coordinate]int
	height, width int
}

func (h heightMap) localMinimums() []coordinate {
	var minimums []coordinate

	for x := 0; x < h.width; x++ {
		for y := 0; y < h.height; y++ {
			coord := coordinate{x, y}
			cell := h.grid[coord]
			left, up, right, down := x-1, y-1, x+1, y+1

			if left >= 0 && cell >= h.grid[coordinate{left, y}] {
				continue
			}

			if up >= 0 && cell >= h.grid[coordinate{x, up}] {
				continue
			}

			if right < h.width && cell >= h.grid[coordinate{right, y}] {
				continue
			}

			if down < h.height && cell >= h.grid[coordinate{x, down}] {
				continue
			}

			minimums = append(minimums, coordinate{x, y})
		}
	}

	return minimums
}

func (h heightMap) riskLevelSum() int {
	total := 0
	for _, coord := range h.localMinimums() {
		total += h.grid[coord] + 1
	}
	return total
}

func (h heightMap) basinSize(coord coordinate, visited map[coordinate]bool) int {
	if _, ok := visited[coord]; ok {
		return 0
	} else if v, ok := h.grid[coord]; v == 9 || !ok {
		return 0
	}

	visited[coord] = true
	size := 1

	size += h.basinSize(coordinate{coord.x - 1, coord.y}, visited) // left
	size += h.basinSize(coordinate{coord.x, coord.y - 1}, visited) // up
	size += h.basinSize(coordinate{coord.x + 1, coord.y}, visited) // right
	size += h.basinSize(coordinate{coord.x, coord.y + 1}, visited) // down

	return size
}

func (h heightMap) basinSizes() []int {
	var sizes []int

	for _, minimum := range h.localMinimums() {
		sizes = append(sizes, h.basinSize(minimum, map[coordinate]bool{}))
	}

	sort.Ints(sizes)

	return sizes
}

func parseInput(input string) heightMap {
	lines := strings.Split(input, "\n")
	grid := make(map[coordinate]int)

	for i, l := range lines {
		horizontal := make([]int, len(l))
		for j, c := range l {
			horizontal[j] = conversions.MustAtoi(string(c))
			grid[coordinate{j, i}] = horizontal[j]

		}
	}

	return heightMap{grid, len(lines), len(lines)}

}

func main() {
	puzzleInput := files.ReadInput()
	heightMap := parseInput(puzzleInput)

	fmt.Println(heightMap.riskLevelSum())

	basinSizes := heightMap.basinSizes()

	fmt.Println(intMath.IntProduct(basinSizes[len(basinSizes)-3:]...))

}
