package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"aoc2021/utils/intMath"
	"fmt"
	"regexp"
	"strings"
)

type coord struct {
	x int
	y int
}

type line struct {
	from coord
	to   coord
}

func parseLines(input string) []line {
	re := regexp.MustCompile(`\d+`)
	lines := make([]line, strings.Count(input, "\n")+1)
	for i, l := range strings.Split(input, "\n") {
		matches := re.FindAllString(l, -1)
		lines[i] = line{
			from: coord{
				x: conversions.MustAtoi(matches[0]),
				y: conversions.MustAtoi(matches[1]),
			},
			to: coord{
				x: conversions.MustAtoi(matches[2]),
				y: conversions.MustAtoi(matches[3]),
			},
		}
	}
	return lines
}

func drawMap(lines []line, diagonals bool) map[coord]int {
	drawingMap := make(map[coord]int)

	for _, l := range lines {
		if l.from.x != l.to.x && l.from.y != l.to.y {
			if !diagonals {
				continue
			}
		}

		var xSign, ySign int
		if l.from.x > l.to.x {
			xSign = -1
		} else if l.from.x < l.to.x {
			xSign = 1
		}
		if l.from.y > l.to.y {
			ySign = -1
		} else if l.from.y < l.to.y {
			ySign = 1
		}

		steps := intMath.IntMax(
			intMath.IntAbs(l.from.y-l.to.y),
			intMath.IntAbs(l.from.x-l.to.x),
		) + 1

		for i := 0; i < steps; i++ {
			drawingMap[coord{
				x: l.from.x + i*xSign,
				y: l.from.y + i*ySign,
			}] += 1
		}

	}

	return drawingMap
}

func countOverlaps(drawnMap map[coord]int) int {
	total := 0

	for _, v := range drawnMap {
		if v > 1 {
			total += 1
		}
	}

	return total
}

func main() {
	puzzleInput := files.ReadInput()
	lines := parseLines(puzzleInput)

	drawnMap := drawMap(lines, false)
	fmt.Println(countOverlaps(drawnMap))

	drawnMapDiagonals := drawMap(lines, true)
	fmt.Println(countOverlaps(drawnMapDiagonals))
}
