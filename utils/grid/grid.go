package grid

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/intMath"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

var UP = Coordinate{X: 0, Y: -1}
var DOWN = Coordinate{X: 0, Y: 1}
var LEFT = Coordinate{X: -1, Y: 0}
var RIGHT = Coordinate{X: 1, Y: 0}

var LEFTDOWN = Coordinate{X: -1, Y: 1}
var LEFTUP = Coordinate{X: -1, Y: -1}
var RIGHDOWN = Coordinate{X: 1, Y: 1}
var RIGHTUP = Coordinate{X: 1, Y: -1}

var orthogonalNeighboorsDeltas = []Coordinate{
	RIGHT,
	DOWN,
	UP,
	LEFT,
}

var diagonalNeighboorsDeltas = []Coordinate{
	LEFTDOWN,
	LEFTUP,
	RIGHDOWN,
	RIGHTUP,
}

func (c *Coordinate) Sum(other Coordinate) Coordinate {
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
		neighboor := c.Sum(Coordinate(delta))
		if neighboor.X < 0 || neighboor.Y < 0 || neighboor.X >= width || neighboor.Y >= height {
			continue
		}
		neighboors = append(neighboors, neighboor)
	}

	return neighboors
}

func (c *Coordinate) ManhattanDistance(other Coordinate) float64 {
	return float64(intMath.IntAbs(other.X-c.X) + intMath.IntAbs(other.Y-c.Y))
}

type Grid struct {
	Cells         map[Coordinate]int
	Height, Width int
}

func (g Grid) String() string {
	var sb strings.Builder
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			sb.WriteString(strconv.Itoa(g.Cells[Coordinate{x, y}]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func StringToGrid(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make(map[Coordinate]int)
	var width int

	for i, l := range lines {
		width = len(l)
		horizontal := make([]int, len(l))
		for j, c := range l {
			horizontal[j] = conversions.MustAtoi(string(c))
			grid[Coordinate{j, i}] = horizontal[j]

		}
	}

	return Grid{grid, len(lines), width}
}
