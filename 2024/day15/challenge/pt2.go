package challenge

import (
	"fmt"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

const (
	BOX_PT2 = "[]"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	warehouse, directions := parsePuzzle(utils.GetPuzzle())

	warehouse = expandWareHouse(warehouse)

	robotPosition := getRobotPosition(warehouse)

	robotWalkPt2(warehouse, robotPosition, directions)

	utils.PrintMatrix(warehouse)

	result = calcBoxGPSCoordinatesPt2(warehouse)
}

func robotWalkPt2(warehouse [][]string, initialPosition Position, directions []Directon) {
	currentPos := initialPosition

	for i, direction := range directions {
		fmt.Println("INDEX ", i)
		utils.PrintMatrix(warehouse)
		nextPos := Position{currentPos.Y + direction.Y, currentPos.X + direction.X}

		if moveBoxesPt2(nextPos, direction, warehouse, true) {
			warehouse[currentPos.Y][currentPos.X] = warehouse[nextPos.Y][nextPos.X]
			warehouse[nextPos.Y][nextPos.X] = ROBOT
			currentPos = nextPos
		}

		//  FIX: At some point near this index a bug is happen, i need to find a way to wait
		//  all checks from all recursions before change positions
		if i > 725 {
			time.Sleep(1 * time.Second / 4)
		}

	}
}

func calcBoxGPSCoordinatesPt2(warehouse [][]string) int {
	total := 0

	for y, row := range warehouse {
		for x, col := range row {
			if col == "[" {
				total += getGPSCoordinate(y, x)
			}
		}
	}

	return total
}

func hitABoxPt2(position Position, warehouse [][]string) bool {
	return strings.Contains(BOX_PT2, warehouse[position.Y][position.X])
}

func moveBoxesPt2(position Position, direction Directon, warehouse [][]string, checkAdjacent bool) bool {
	nextPos := Position{position.Y + direction.Y, position.X + direction.X}

	if hitAFloor(position, warehouse) {
		return true
	}

	if hitAWall(position, warehouse) || hitABox(nextPos, warehouse) {
		return false
	}

	edge := warehouse[position.Y][position.X]

	if isVerticalWalk(direction) {
		if !moveBoxesPt2(nextPos, direction, warehouse, true) {
			return false
		}

		if checkAdjacent {
			adjacentBoxEdge := getAdjacentBoxEdge(edge, position)
			if !moveBoxesPt2(adjacentBoxEdge, direction, warehouse, false) {
				return false
			}
		}
	} else if !moveBoxesPt2(nextPos, direction, warehouse, true) {
		return false
	}

	warehouse[nextPos.Y][nextPos.X] = warehouse[position.Y][position.X]
	warehouse[position.Y][position.X] = FLOOR
	return true
}

func getAdjacentBoxEdge(edge string, position Position) Position {
	switch edge {
	case "[":
		return Position{position.Y, position.X + 1}
	case "]":
		return Position{position.Y, position.X - 1}
	default:
		return position
	}
}

func isVerticalWalk(dir Directon) bool {
	return dir.Y != 0
}

func expandWareHouse(warehouse [][]string) [][]string {
	expandedWarehouse := make([][]string, len(warehouse))

	for i, row := range warehouse {
		expandedWarehouse[i] = []string{}
		for _, col := range row {
			switch col {
			case WALL:
				expandedWarehouse[i] = append(expandedWarehouse[i], "#", "#")
			case BOX:
				expandedWarehouse[i] = append(expandedWarehouse[i], "[", "]")
			case FLOOR:
				expandedWarehouse[i] = append(expandedWarehouse[i], ".", ".")
			case ROBOT:
				expandedWarehouse[i] = append(expandedWarehouse[i], "@", ".")
			default:
				panic("unexpected char")
			}
		}
	}

	return expandedWarehouse
}
