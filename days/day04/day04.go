package main

import (
	"aoc2021/utils/conversions"
	"aoc2021/utils/files"
	"fmt"
	"strings"
)

type bingoBoard struct {
	grid [5][5]string
}

func (b *bingoBoard) mark(draw string) bool {
	// Add mark
	for y, line := range b.grid {
		for x := range line {
			if b.grid[y][x] == draw {
				b.grid[y][x] = "X"
			}
		}
	}

	marksArray := [5]string{"X", "X", "X", "X", "X"}
	// Check if winner horizontally
	for _, line := range b.grid {
		if line == marksArray {
			return true
		}
	}

	// Check if winner vertically
	for x := range b.grid {
		if b.grid[0][x] == "X" && b.grid[1][x] == "X" && b.grid[2][x] == "X" && b.grid[3][x] == "X" && b.grid[4][x] == "X" {
			return true
		}
	}

	return false
}

func (b bingoBoard) sumUnmarked() int {
	total := 0
	for _, line := range b.grid {
		for _, cell := range line {
			if cell != "X" {
				total += conversions.MustAtoi(cell)
			}
		}
	}
	return total
}

func newBingoBoardFromLines(lines []string) *bingoBoard {
	var grid [5][5]string
	for i, line := range lines {
		for j, cell := range strings.Split(strings.ReplaceAll(strings.TrimLeft(line, " "), "  ", " "), " ") {
			grid[i][j] = cell
		}
	}
	return &bingoBoard{
		grid: grid,
	}
}

func parseInput(input string) ([]*bingoBoard, []string) {
	lines := strings.Split(input, "\n")
	draws := strings.Split(lines[0], ",")

	bingoBoards := make([]*bingoBoard, len(lines[1:])/6)

	for i := 1; i < len(lines); i += 6 {
		bingoBoards[(i-1)/6] = newBingoBoardFromLines(lines[i+1 : i+6])
	}

	return bingoBoards, draws

}

func playBingo(boards []*bingoBoard, draws []string) (*bingoBoard, string) {
	for _, draw := range draws {
		for _, b := range boards {
			if winner := b.mark(draw); winner {
				return b, draw
			}
		}
	}
	panic("Bingo shouldn't end without winner!")
}

func playBingoUntilLast(boards []*bingoBoard, draws []string) (*bingoBoard, string) {
	playingBoards := make([]*bingoBoard, len(boards))
	copy(playingBoards, boards)

	for _, draw := range draws {
		for i := 0; i < len(playingBoards); i++ {
			b := playingBoards[i]
			if winner := b.mark(draw); winner {
				playingBoards = append(playingBoards[:i], playingBoards[i+1:]...)
				i--
				if len(playingBoards) == 0 {
					return b, draw
				}
			}
		}
	}
	panic("Bingo shouldn't end without winner!")
}

func main() {
	puzzleInput := files.ReadInput()
	boards, draws := parseInput(puzzleInput)

	winner, lastDraw := playBingo(boards, draws)
	fmt.Println(winner.sumUnmarked() * conversions.MustAtoi(lastDraw))

	// Can reuse boards as bingo order is the same
	lastWinner, lastLastDraw := playBingoUntilLast(boards, draws)
	fmt.Println(lastWinner.sumUnmarked() * conversions.MustAtoi(lastLastDraw))

}
