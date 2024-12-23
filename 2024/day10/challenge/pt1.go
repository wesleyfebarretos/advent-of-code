package challenge

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Position struct {
	Y int
	X int
}

func Pt1() {
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
		seen := makeSeenMatrix(parsedPuzzle)
		sum += walk(parsedPuzzle, directions, seen, th, 0, 9)
	}

	fmt.Printf("Part 1 -> %d", sum)
}

func walk(parsedPuzzle [][]int, directions []Position, seen [][]bool, currPosition Position, currentNumber, target int) int {
	yLen := len(parsedPuzzle)
	xLen := len(parsedPuzzle[0])

	if currPosition.Y < 0 || currPosition.Y >= yLen || currPosition.X < 0 || currPosition.X >= xLen {
		return 0
	}

	if seen[currPosition.Y][currPosition.X] {
		return 0
	}

	if parsedPuzzle[currPosition.Y][currPosition.X] != currentNumber {
		return 0
	}

	if parsedPuzzle[currPosition.Y][currPosition.X] == target {
		seen[currPosition.Y][currPosition.X] = true
		return 1
	}

	points := 0

	for _, direction := range directions {
		nextPosition := Position{
			Y: currPosition.Y + direction.Y,
			X: currPosition.X + direction.X,
		}

		points += walk(parsedPuzzle, directions, seen, nextPosition, currentNumber+1, target)
	}

	return points
}

func makeSeenMatrix(parsedPuzzle [][]int) [][]bool {
	seen := make([][]bool, len(parsedPuzzle))
	for i := range seen {
		seen[i] = make([]bool, len(parsedPuzzle[0]))
	}

	return seen
}

func getTrailHeadsPositions(parsedPuzzle [][]int) []Position {
	trailHeadPositions := []Position{}

	for y, line := range parsedPuzzle {
		for x, num := range line {
			if num == 0 {
				trailHeadPositions = append(trailHeadPositions, Position{y, x})
			}
		}
	}

	return trailHeadPositions
}

func parsePuzzle(puzzle string) [][]int {
	puzzleArr := strings.Split(puzzle, "\n")

	parsedPuzzle := make([][]int, len(puzzleArr))

	for i, line := range puzzleArr {
		parsedPuzzle[i] = make([]int, len(line))
		for y, num := range line {
			n, _ := strconv.Atoi(string(num))
			parsedPuzzle[i][y] = n
		}
	}

	return parsedPuzzle
}

func printPuzzle(puzzle string, pos Position) {
	puzzleArr := strings.Split(puzzle, "\n")

	parsedPuzzle := make([][]string, len(puzzleArr))

	for i, line := range puzzleArr {
		parsedPuzzle[i] = make([]string, len(line))
		for y, num := range line {
			parsedPuzzle[i][y] = string(num)
		}
	}

	parsedPuzzle[pos.Y][pos.X] = "*"

	s := strings.Builder{}

	for _, line := range puzzleArr {
		for _, num := range line {
			s.WriteRune(num)
		}
		s.WriteString("\n")
	}

	fmt.Println(s.String())
}
