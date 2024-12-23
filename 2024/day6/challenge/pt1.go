package challenge

import (
	"fmt"
	"log"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Position struct {
	y int
	x int
}

func Pt1() {
	puzzle := parsePuzzle(utils.GetPuzzle())

	seen := make([][]bool, len(puzzle))

	for i := range seen {
		seen[i] = make([]bool, len(puzzle[0]))
	}

	position := findGuardPosition(puzzle)

	puzzle[position.y][position.x] = "."

	distinctPositions := walk(puzzle, seen, position, '^')

	fmt.Printf("Part 1 -> %d", distinctPositions)
}

func walk(puzzle [][]string, seen [][]bool, position Position, guardDireciton rune) int {
	if position.y < 0 || position.y >= len(puzzle) {
		return 0
	}

	if position.x < 0 || position.x >= len(puzzle[0]) {
		return 0
	}

	nextPosition, newGuardDirection := findNextPosition(puzzle, position, guardDireciton)

	if seen[position.y][position.x] || puzzle[position.y][position.x] != "." {
		return 0 + walk(puzzle, seen, nextPosition, newGuardDirection)
	}

	// printWay(puzzle, position, guardDireciton)

	seen[position.y][position.x] = true

	return 1 + walk(puzzle, seen, nextPosition, newGuardDirection)
}

func findNextPosition(puzzle [][]string, position Position, guardDireciton rune) (Position, rune) {
	if puzzle[position.y][position.x] == "#" {
		return changeDirectionHandle(position, guardDireciton)
	}

	nextPosition := Position{
		y: position.y,
		x: position.x,
	}

	switch guardDireciton {
	case '>':
		nextPosition.x += 1
	case '^':
		nextPosition.y -= 1
	case 'v':
		nextPosition.y += 1
	case '<':
		nextPosition.x -= 1
	}

	return nextPosition, guardDireciton
}

func changeDirectionHandle(position Position, guardDireciton rune) (Position, rune) {
	changeDirection := Position{
		y: position.y,
		x: position.x,
	}

	var newGuardDirection rune

	switch guardDireciton {
	case '>':
		changeDirection.x -= 1
		changeDirection.y += 1
		newGuardDirection = 'v'
	case '^':
		changeDirection.y += 1
		changeDirection.x += 1
		newGuardDirection = '>'
	case 'v':
		changeDirection.y -= 1
		changeDirection.x -= 1
		newGuardDirection = '<'
	case '<':
		changeDirection.x += 1
		changeDirection.y -= 1
		newGuardDirection = '^'
	}

	return changeDirection, newGuardDirection
}

func findGuardPosition(puzzle [][]string) Position {
	for y, line := range puzzle {
		for x, character := range line {
			if character == "^" {
				return Position{y, x}
			}
		}
	}

	log.Fatal("guard not found")

	return Position{}
}

func parsePuzzle(puzzle string) [][]string {
	lines := strings.Split(puzzle, "\n")
	matrix := make([][]string, len(lines))

	for i := range matrix {
		matrix[i] = make([]string, len(lines[0]))
	}

	for i := range lines {
		for j, v := range lines[i] {
			matrix[i][j] = string(v)
		}
	}

	return matrix
}

func printWay(puzzle [][]string, position Position, guardDireciton rune) {
	debug := ""
	// time.Sleep(1 * time.Second / 8)
	for i := range puzzle {
		for j := range puzzle[i] {
			if i == position.y && j == position.x {
				debug += string(guardDireciton)
			}
			debug += puzzle[i][j]
		}
		debug += "\n"
	}
	fmt.Println(debug)
}
