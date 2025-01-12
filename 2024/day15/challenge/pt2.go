package challenge

import (
	"container/list"
	"fmt"
	"slices"
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

	robotWalkBFSApproach(warehouse, robotPosition, directions)

	// robotWalkDFSApproachPt2(warehouse, robotPosition, directions)

	result = calcBoxGPSCoordinatesPt2(warehouse)
}

func robotWalkBFSApproach(warehouse [][]string, initialPosition Position, directions []Directon) {
	robotPosition := initialPosition

	for _, dir := range directions {
		queue := list.New()

		queue.PushBack(robotPosition)

		seen := [][2]int{}

		for queue.Len() > 0 {
			currPos := queue.Remove(queue.Front()).(Position)

			seen = append(seen, [2]int{currPos.Y, currPos.X})

			nextPos := Position{currPos.Y + dir.Y, currPos.X + dir.X}

			if warehouse[nextPos.Y][nextPos.X] == WALL {
				seen = [][2]int{}
				break
			}

			if warehouse[nextPos.Y][nextPos.X] == FLOOR {
				continue
			}

			if strings.Contains(BOX_PT2, warehouse[nextPos.Y][nextPos.X]) {
				queue.PushBack(nextPos)

				if isVerticalWalk(dir) {
					if warehouse[nextPos.Y][nextPos.X] == "[" {
						queue.PushBack(Position{nextPos.Y, nextPos.X + 1})
					} else {
						queue.PushBack(Position{nextPos.Y, nextPos.X - 1})
					}
				}
			}
		}

		if len(seen) > 0 {
			robotPosition.Y = robotPosition.Y + dir.Y
			robotPosition.X = robotPosition.X + dir.X
		}

		slices.Reverse(seen)

		duplicates := make(map[string]bool)

		for _, row := range seen {
			y, x := row[0], row[1]

			if duplicates[fmt.Sprintf("%d|%d", y, x)] {
				continue
			}

			duplicates[fmt.Sprintf("%d|%d", y, x)] = true

			warehouse[y+dir.Y][x+dir.X] = warehouse[y][x]
			warehouse[y][x] = FLOOR
		}
	}
}

func robotWalkDFSApproachPt2(warehouse [][]string, initialPosition Position, directions []Directon) {
	currentPos := initialPosition

	for _, direction := range directions {
		nextPos := Position{currentPos.Y + direction.Y, currentPos.X + direction.X}

		walkSlice := [][2]int{}

		if moveBoxesPt2(nextPos, direction, warehouse, true, &walkSlice) {
			seen := make(map[string]struct{})

			for _, row := range walkSlice {
				y, x := row[0], row[1]

				if _, ok := seen[fmt.Sprintf("%d|%d", y, x)]; ok {
					continue
				}

				seen[fmt.Sprintf("%d|%d", y, x)] = struct{}{}

				warehouse[y+direction.Y][x+direction.X] = warehouse[y][x]
				warehouse[y][x] = "."
			}

			warehouse[currentPos.Y][currentPos.X] = warehouse[nextPos.Y][nextPos.X]
			warehouse[nextPos.Y][nextPos.X] = ROBOT
			currentPos = nextPos
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

func moveBoxesPt2(position Position, direction Directon, warehouse [][]string, checkAdjacent bool, walkSlice *[][2]int) bool {
	nextPos := Position{position.Y + direction.Y, position.X + direction.X}

	if hitAFloor(position, warehouse) {
		return true
	}

	if hitAWall(position, warehouse) || hitABox(nextPos, warehouse) {
		return false
	}

	edge := warehouse[position.Y][position.X]

	if isVerticalWalk(direction) {
		if !moveBoxesPt2(nextPos, direction, warehouse, true, walkSlice) {
			return false
		}

		if checkAdjacent {
			adjacentBoxEdge := getAdjacentBoxEdge(edge, position)
			if !moveBoxesPt2(adjacentBoxEdge, direction, warehouse, false, walkSlice) {
				return false
			}
		}
	} else if !moveBoxesPt2(nextPos, direction, warehouse, true, walkSlice) {
		return false
	}

	*walkSlice = append(*walkSlice, [2]int{position.Y, position.X})
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
