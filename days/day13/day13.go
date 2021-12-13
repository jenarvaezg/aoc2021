package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"aoc2021/utils/grid"
	"aoc2021/utils/intMath"
	"fmt"
	"strings"
)

type paperFold struct {
	axis string
	line int
}

type paperMap grid.Grid

func (o paperMap) String() string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			grid.Grid(o).String(),
			"1",
			"#",
		),
		"0",
		".",
	)
}

func (o paperMap) countDots() int {
	return len(o.Cells)
}

func (o *paperMap) fold(fold paperFold) {
	if fold.axis == "x" {
		for c := range o.Cells {
			if c.X < fold.line {
				continue
			}
			if c.X > fold.line {
				if o.Cells[c] == 1 {
					o.Cells[grid.Coordinate{X: fold.line - (c.X - fold.line), Y: c.Y}] = 1
				}
			}
			delete(o.Cells, c)
		}
		o.Width = fold.line
	} else {
		for c := range o.Cells {
			if c.Y < fold.line {
				continue
			}
			if c.Y > fold.line {
				if o.Cells[c] == 1 {
					o.Cells[grid.Coordinate{X: c.X, Y: fold.line - (c.Y - fold.line)}] = 1
				}
			}
			delete(o.Cells, c)
		}
		o.Height = fold.line
	}
}

func parsePaper(input string) (paperMap, []paperFold) {
	lines := strings.Split(input, "\n")
	g := make(map[grid.Coordinate]int)
	var folds []paperFold
	var width, height int

	var i int
	for i = 0; lines[i] != ""; i++ {
		coordsSplit := strings.Split(lines[i], ",")
		x, y := conversions.MustAtoi(coordsSplit[0]), conversions.MustAtoi(coordsSplit[1])
		width = intMath.IntMax(width, x)
		height = intMath.IntMax(height, y)
		g[grid.Coordinate{X: x, Y: y}] = 1
	}

	for i = i + 1; i < len(lines); i++ {
		fmt.Println(i, lines[i])
		word := strings.Split(lines[i], " ")[2]
		foldSplit := strings.Split(word, "=")
		folds = append(folds, paperFold{foldSplit[0], conversions.MustAtoi(foldSplit[1])})
	}

	return paperMap(grid.Grid{Cells: g, Height: height + 1, Width: width + 1}), folds
}

func main() {
	puzzleInput := files.ReadInput()
	paper, folds := parsePaper(puzzleInput)

	paper.fold(folds[0])
	fmt.Println(paper.countDots())

	for i := 1; i < len(folds); i++ {
		paper.fold(folds[i])
	}

	fmt.Println(paper)
}
