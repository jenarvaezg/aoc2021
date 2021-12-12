package grid

import (
	"aoc2021/utils/conversions"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

var orthogonalNeighboorsDeltas = []Coordinate{
	{-1, 0},
	{0, 1},
	{0, -1},
	{1, 0},
}

var diagonalNeighboorsDeltas = []Coordinate{
	{-1, 1},
	{-1, -1},
	{1, 1},
	{1, -1},
}

func (c *Coordinate) sum(other Coordinate) Coordinate {
	return Coordinate{
		c.X + other.X,
		c.Y + other.Y,
	}
}

func (c *Coordinate) Neighboors(diagonal bool, height, width int) []Coordinate {
	var neighboors []Coordinate

	deltas := orthogonalNeighboorsDeltas
	if diagonal {
		deltas = append(deltas, diagonalNeighboorsDeltas...)
	}

	for _, delta := range deltas {
		neighboor := c.sum(delta)
		if neighboor.X < 0 || neighboor.Y < 0 || neighboor.X >= width || neighboor.Y >= height {
			continue
		}
		neighboors = append(neighboors, neighboor)
	}

	//fmt.Println("I am ", c, "neighboors are", neighboors)

	return neighboors

}

type Grid struct {
	Cells         map[Coordinate]int
	Height, Width int
}

func (g Grid) String() string {
	var sb strings.Builder
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			sb.WriteString(strconv.Itoa(g.Cells[Coordinate{X: x, Y: y}]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func StringToGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make(map[Coordinate]int)

	for i, l := range lines {
		horizontal := make([]int, len(l))
		for j, c := range l {
			horizontal[j] = conversions.MustAtoi(string(c))
			grid[Coordinate{j, i}] = horizontal[j]

		}
	}

	return Grid{grid, len(lines), len(lines)}
}
