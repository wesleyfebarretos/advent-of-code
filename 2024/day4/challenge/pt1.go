package challenge

import (
	"fmt"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

const XMAS = "XMAS"

type Matrix [][]string

type Seen [][]bool

type Directions [][]int

type XPosition struct {
	y int
	x int
}

func Pt1() {
	puzzle := getMatrixOfLetters(utils.GetPuzzle())

	lookDirections := Directions{
		{0, -1},  // weast
		{0, 1},   // east
		{1, 0},   // south
		{-1, 0},  // north
		{-1, -1}, // north-weast
		{-1, 1},  // north-east
		{1, 1},   // south-east
		{1, -1},  // south-weast
	}

	xPositions := getSliceOfXPositions(puzzle)

	xmasQty := findXMASQty(puzzle, xPositions, lookDirections)

	fmt.Printf("Part 1 -> %d", xmasQty)
}

func findXMASQty(puzzle Matrix, xPositions []XPosition, lookDirections Directions) int {
	xmasTotal := 0
	for _, xPosition := range xPositions {
		for _, direction := range lookDirections {
			xmasTotal += walk(puzzle, xPosition, direction, &strings.Builder{})
		}
	}

	return xmasTotal
}

func walk(puzzle Matrix, xPosition XPosition, direction []int, sb *strings.Builder) int {
	if xPosition.y < 0 || xPosition.y >= len(puzzle) {
		return 0
	}

	if xPosition.x < 0 || xPosition.x >= len(puzzle[0]) {
		return 0
	}

	sb.WriteString(puzzle[xPosition.y][xPosition.x])

	word := sb.String()

	if word == XMAS {
		sb.Reset()
		return 1
	}

	for i := range word {
		if string(word[i]) != string(XMAS[i]) {
			return 0
		}
	}

	for i := 0; i < 4; i++ {
		newPosition := XPosition{
			y: xPosition.y + direction[0],
			x: xPosition.x + direction[1],
		}
		if walk(puzzle, newPosition, direction, sb) == 1 {
			return 1
		}
	}

	return 0
}

func getMatrixOfLetters(puzzle string) Matrix {
	lines := strings.Split(puzzle, "\n")

	matrix := make([][]string, len(lines))

	for i := range matrix {
		matrix[i] = make([]string, len(lines[0]))
	}

	for i, line := range lines {
		for j, c := range line {
			matrix[i][j] = string(c)
		}
	}

	return matrix
}

func getSliceOfXPositions(puzzle [][]string) []XPosition {
	xPositions := []XPosition{}

	for i, line := range puzzle {
		for j, c := range line {
			if c == "X" {
				xPositions = append(xPositions, XPosition{y: i, x: j})
			}
		}
	}

	return xPositions
}
