package challenge

import (
	"fmt"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	parsedPuzzle := parsePuzzle(utils.GetPuzzle())

	trailHeadPositions := getTrailHeadsPositions(parsedPuzzle)

	directions := []Position{
		{-1, 0}, // North
		{1, 0},  // South
		{0, 1},  // East
		{0, -1}, // West
	}

	sum := 0

	for _, th := range trailHeadPositions {
		sum += walk2(parsedPuzzle, directions, th, 0, 9)
	}

	fmt.Printf("Part 2 -> %d", sum)
}

func walk2(parsedPuzzle [][]int, directions []Position, currPosition Position, currentNumber, target int) int {
	yLen := len(parsedPuzzle)
	xLen := len(parsedPuzzle[0])

	if currPosition.Y < 0 || currPosition.Y >= yLen || currPosition.X < 0 || currPosition.X >= xLen {
		return 0
	}

	if parsedPuzzle[currPosition.Y][currPosition.X] != currentNumber {
		return 0
	}

	if parsedPuzzle[currPosition.Y][currPosition.X] == target {
		return 1
	}

	points := 0

	for _, direction := range directions {
		nextPosition := Position{
			Y: currPosition.Y + direction.Y,
			X: currPosition.X + direction.X,
		}

		points += walk2(parsedPuzzle, directions, nextPosition, currentNumber+1, target)
	}

	return points
}
