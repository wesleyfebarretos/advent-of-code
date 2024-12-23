package challenge

import (
	"fmt"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	puzzle := parsePuzzle(utils.GetPuzzle())

	seen := make([][]bool, len(puzzle))

	for i := range seen {
		seen[i] = make([]bool, len(puzzle[0]))
	}

	position := findGuardPosition(puzzle)

	puzzle[position.y][position.x] = "."

	stuckLoops := findStuckLoops(puzzle, seen, position, '^')

	fmt.Printf("Part 2 -> %d", stuckLoops)
}

func findStuckLoops(puzzle [][]string, seen [][]bool, position Position, guardDireciton rune) int {
	walk(puzzle, seen, position, guardDireciton)
	total := 0

	for y := range seen {
		for x := range seen[y] {

			if !seen[y][x] || y == position.y && x == position.x {
				continue
			}

			puzzle[y][x] = "#"

			obstructionMap := make(map[string]struct{})

			total += walkAndFindStuckLoop(puzzle, position, guardDireciton, obstructionMap)

			puzzle[y][x] = "."
		}
	}

	return total
}

func walkAndFindStuckLoop(puzzle [][]string, position Position, guardDireciton rune, obstructionMap map[string]struct{}) int {
	if position.y < 0 || position.y >= len(puzzle) {
		return 0
	}

	if position.x < 0 || position.x >= len(puzzle[0]) {
		return 0
	}

	// if puzzle[position.y][position.x] != "#" {
	// 	printWay(puzzle, position, guardDireciton)
	// }

	if puzzle[position.y][position.x] == "#" {
		obstructionPosition := fmt.Sprintf("%d%d%c", position.y, position.x, guardDireciton)

		if _, ok := obstructionMap[obstructionPosition]; ok {
			return 1
		}
		obstructionMap[obstructionPosition] = struct{}{}
	}

	nextPosition, newGuardDirection := findNextPosition(puzzle, position, guardDireciton)

	return walkAndFindStuckLoop(puzzle, nextPosition, newGuardDirection, obstructionMap)
}
