package challenge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Robot struct {
	Position  Location
	Direction Location
}

type Location struct {
	X int
	Y int
}

const (
	GRID_LEN     = 103
	GRID_ROW_LEN = 101
	SECONDS      = 100
)

func Pt1() {
	result := 1

	defer func(t time.Time) {
		fmt.Printf("\nPart 1 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	space := makeGrid()

	robots := parsePuzzle(utils.GetPuzzle())

	walkWithRobots(robots)

	fillSpaceWithRobotsPosition(robots, space)

	quads := cutQuads(space)

	for _, quad := range quads {
		quadSum := 0
		for _, num := range quad {
			quadSum += num
		}

		result *= quadSum
	}
}

func cutQuads(space [][]int) [][]int {
	northWest := []int{}
	northEast := []int{}
	southWest := []int{}
	southEast := []int{}

	midX := len(space[0]) / 2
	midY := len(space) / 2

	for y, row := range space {

		if y == midY {
			continue
		}

		for x := range row {

			if x == midX {
				continue
			}

			if y < midY {
				if x < midX {
					northWest = append(northWest, row[x])
				} else {
					northEast = append(northEast, row[x])
				}

				continue
			}

			if x < midX {
				southWest = append(southWest, row[x])
			} else {
				southEast = append(southEast, row[x])
			}
		}
	}

	return [][]int{northWest, northEast, southWest, southEast}
}

func fillSpaceWithRobotsPosition(robots []Robot, space [][]int) {
	for _, robot := range robots {
		space[robot.Position.Y][robot.Position.X]++
	}
}

func walkWithRobots(robots []Robot) {
	for i := range robots {
		for range SECONDS {
			robots[i].Position = getNextPosition(robots[i])
		}
	}
}

func getNextPosition(robot Robot) Location {
	x := wrap(robot.Position.X+robot.Direction.X, GRID_ROW_LEN)
	y := wrap(robot.Position.Y+robot.Direction.Y, GRID_LEN)

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

func makeGrid() [][]int {
	grid := make([][]int, GRID_LEN)

	for i := range grid {
		grid[i] = make([]int, GRID_ROW_LEN)
	}

	return grid
}

func printSpace(space [][]int) {
	s := strings.Builder{}

	for _, row := range space {
		for _, num := range row {
			s.WriteString(fmt.Sprintf("%d", num))
		}
		s.WriteString("\n")
	}

	fmt.Println(s.String())
}

func parsePuzzle(puzzle string) []Robot {
	puzzleSlice := strings.Split(puzzle, "\n")

	robots := make([]Robot, len(puzzleSlice))

	reg := regexp.MustCompile(`[^0-9,\s-]`)

	for i := range robots {
		robotInfo := strings.Split(reg.ReplaceAllString(puzzleSlice[i], ""), " ")

		initialPosition := strings.Split(robotInfo[0], ",")
		direction := strings.Split(robotInfo[1], ",")

		x1, _ := strconv.Atoi(initialPosition[0])
		y1, _ := strconv.Atoi(initialPosition[1])
		x2, _ := strconv.Atoi(direction[0])
		y2, _ := strconv.Atoi(direction[1])

		robots[i] = Robot{
			Position: Location{
				X: x1,
				Y: y1,
			},
			Direction: Location{
				X: x2,
				Y: y2,
			},
		}
	}

	return robots
}
