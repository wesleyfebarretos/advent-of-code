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

	warehouse, directions := parsePuzzle(utils.GetTestPuzzle())

	warehouse = expandWareHouse(warehouse)

	robotPosition := getRobotPosition(warehouse)

	robotWalkPt2(warehouse, robotPosition, directions)

	result = calcBoxGPSCoordinates(warehouse)
}

func robotWalkPt2(warehouse [][]string, initialPosition Position, directions []Directon) {
	currentPos := initialPosition

	for _, direction := range directions {
		utils.PrintMatrix(warehouse)
		nextPos := Position{currentPos.Y + direction.Y, currentPos.X + direction.X}

		if hitAWall(nextPos, warehouse) {
			continue
		}

		if hitAFloor(nextPos, warehouse) || moveBoxesPt2(nextPos, direction, warehouse) {
			warehouse[currentPos.Y][currentPos.X] = warehouse[nextPos.Y][nextPos.X]
			warehouse[nextPos.Y][nextPos.X] = ROBOT
			currentPos = nextPos
		}

		time.Sleep(1 * time.Second / 2)
	}
}

func hitABoxPt2(position Position, warehouse [][]string) bool {
	return strings.Contains(BOX_PT2, warehouse[position.Y][position.X])
}

func moveBoxesPt2(position Position, direction Directon, warehouse [][]string) bool {
	nextPos := Position{position.Y + direction.Y, position.X + direction.X}

	if hitAWall(nextPos, warehouse) {
		return false
	}
	utils.PrintMatrix(warehouse)

	//  FIX:
	//  Something is wrong with logic, its not moving correctly the box's edges
	if hitABoxPt2(nextPos, warehouse) {
		if isVerticalWalk(direction) && !moveBoxesPt2(getNextPositionByBoxEdge(nextPos, warehouse), direction, warehouse) {
			return false
		}

		if !moveBoxesPt2(nextPos, direction, warehouse) {
			return false
		}
	}

	warehouse[nextPos.Y][nextPos.X] = warehouse[position.Y][position.X]
	warehouse[position.Y][position.X] = FLOOR

	return true
}

func getNextPositionByBoxEdge(position Position, warehouse [][]string) Position {
	switch warehouse[position.Y][position.X] {
	case "[":
		return Position{position.Y, position.X + 1}
	case "]":
		return Position{position.Y, position.X - 1}
	default:
		panic("unexpected box edge")
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
				panic("non expected char")
			}
		}
	}

	return expandedWarehouse
}
