package challenge

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

const (
	CANVAS_ZERO_VALUE  = "."
	CANVAS_ROBOT_VALUE = "*"
)

func Pt2() {
	result := 1

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	space := makeGrid()

	canvas := makePrintterGrid()

	robots := parsePuzzle(utils.GetPuzzle())

	fillSpaceWithRobotsPosition(robots, space)

	result = walkWithRobotsPt2(robots, space, canvas)
}

func walkWithRobotsPt2(robots []Robot, space [][]int, canvas [][]string) int {
	for seconds := 1; ; seconds++ {
		for i := range robots {
			robots[i].Position = getNextPosition(robots[i])
		}
		if findChristmansTree(robots) {
			// Just to Print The Tree
			resetSpace(space)
			fillSpaceWithRobotsPosition(robots, space)
			drawInCanvasBySpacePositions(space, canvas)
			printSpace(canvas)
			// End
			return seconds
		}
	}
}

func findChristmansTree(robots []Robot) bool {
	robotsByColumnMap := make(map[int][]Robot)

	minNeighbors := 10

	for _, robot := range robots {
		if robotsByColumnMap[robot.Position.Y] == nil {
			robotsByColumnMap[robot.Position.Y] = []Robot{}
		}

		robotsByColumnMap[robot.Position.Y] = append(robotsByColumnMap[robot.Position.Y], robot)
	}

	for _, robots := range robotsByColumnMap {
		if len(robots) < minNeighbors {
			continue
		}

		xPositions := []int{}

		for _, robot := range robots {
			xPositions = append(xPositions, robot.Position.X)
		}

		slices.Sort(xPositions)

		neighborsQty := 1

		for x := 0; x < len(xPositions)-1; x++ {
			if xPositions[x+1]-xPositions[x] == 1 {
				neighborsQty++
			} else {
				neighborsQty = 1
			}
		}

		if neighborsQty >= minNeighbors {
			return true
		}
	}

	return false
}

func resetSpace(space [][]int) {
	for y := range space {
		for x := range space[y] {
			space[y][x] = 0
		}
	}
}

func makePrintterGrid() [][]string {
	grid := make([][]string, GRID_HEIGHT)

	for i := range grid {
		grid[i] = make([]string, GRID_WIDTH)
	}

	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = CANVAS_ZERO_VALUE
		}
	}

	return grid
}

func drawInCanvasBySpacePositions(space [][]int, canvas [][]string) {
	for y := range space {
		for x := range space[y] {
			if space[y][x] > 0 {
				canvas[y][x] = CANVAS_ROBOT_VALUE
			} else {
				canvas[y][x] = CANVAS_ZERO_VALUE
			}
		}
	}
}

func getNextPosition(robot Robot) Location {
	x := wrap(robot.Position.X+robot.Direction.X, GRID_WIDTH)
	y := wrap(robot.Position.Y+robot.Direction.Y, GRID_HEIGHT)

	return Location{X: x, Y: y}
}

func wrap(value, max int) int {
	if value < 0 {
		return value + max
	}

	if value >= max {
		return value - max
	}

	return value
}

func printSpace[S any](space [][]S) {
	s := strings.Builder{}

	for _, row := range space {
		for _, num := range row {
			s.WriteString(fmt.Sprintf("%v", num))
		}
		s.WriteString("\n")
	}

	fmt.Println(s.String())
}
