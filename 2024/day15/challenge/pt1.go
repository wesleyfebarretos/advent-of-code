package challenge

import (
	"fmt"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Directon struct {
	Y int
	X int
}

type Position struct {
	Y int
	X int
}

const (
	ROBOT = "@"
	WALL  = "#"
	BOX   = "O"
	FLOOR = "."
)

func Pt1() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 1 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	warehouse, directions := parsePuzzle(utils.GetPuzzle())

	robotPosition := getRobotPosition(warehouse)

	robotWalk(warehouse, robotPosition, directions)

	result = calcBoxGPSCoordinates(warehouse)
}

func robotWalk(warehouse [][]string, initialPosition Position, directions []Directon) {
	currentPos := initialPosition

	for _, direction := range directions {
		nextPos := Position{currentPos.Y + direction.Y, currentPos.X + direction.X}

		if hitAWall(nextPos, warehouse) {
			continue
		}

		if hitAFloor(nextPos, warehouse) || moveBoxes(nextPos, direction, warehouse) {
			warehouse[currentPos.Y][currentPos.X] = warehouse[nextPos.Y][nextPos.X]
			warehouse[nextPos.Y][nextPos.X] = ROBOT
			currentPos = nextPos
		}
	}
}

func calcBoxGPSCoordinates(warehouse [][]string) int {
	total := 0

	for y, row := range warehouse {
		for x, col := range row {
			if col == BOX {
				total += getGPSCoordinate(y, x)
			}
		}
	}

	return total
}

func getGPSCoordinate(y, x int) int {
	return 100*y + x
}

func moveBoxes(position Position, direction Directon, warehouse [][]string) bool {
	nextPos := Position{position.Y + direction.Y, position.X + direction.X}

	if hitAWall(nextPos, warehouse) {
		return false
	}

	if hitABox(nextPos, warehouse) {
		if !moveBoxes(nextPos, direction, warehouse) {
			return false
		}
	}

	warehouse[nextPos.Y][nextPos.X] = warehouse[position.Y][position.X]
	warehouse[position.Y][position.X] = FLOOR

	return true
}

func hitAFloor(position Position, warehouse [][]string) bool {
	return warehouse[position.Y][position.X] == FLOOR
}

func hitAWall(position Position, warehouse [][]string) bool {
	return warehouse[position.Y][position.X] == WALL
}

func hitABox(position Position, warehouse [][]string) bool {
	return warehouse[position.Y][position.X] == BOX
}

func getRobotPosition(warehouse [][]string) Position {
	for y, row := range warehouse {
		for x, col := range row {
			if col == ROBOT {
				return Position{y, x}
			}
		}
	}

	panic("ROBOT NOT FOUND")
}

func parsePuzzle(puzzle string) ([][]string, []Directon) {
	puzzleSlice := strings.Split(puzzle, "\n\n")

	puzzleMap, puzzleDirections := puzzleSlice[0], puzzleSlice[1]

	puzzleMapSlice := strings.Split(puzzleMap, "\n")

	puzzleDirectionsSlice := strings.Split(puzzleDirections, "\n")

	matrix := make([][]string, len(puzzleMapSlice))

	for i, row := range puzzleMapSlice {
		matrix[i] = make([]string, (len(row)))
		for y, symbol := range row {
			matrix[i][y] = string(symbol)
		}
	}

	directions := []Directon{}

	for _, row := range puzzleDirectionsSlice {
		for _, col := range row {
			var dir Directon

			switch string(col) {
			case "<":
				dir = Directon{0, -1}
			case ">":
				dir = Directon{0, 1}
			case "v":
				dir = Directon{1, 0}
			case "^":
				dir = Directon{-1, 0}
			}

			directions = append(directions, dir)
		}
	}

	return matrix, directions
}
