package challenge

import (
	"fmt"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Position struct {
	Y int
	X int
}

func Pt1() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("Part 1 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	garden := parsePuzzle(utils.GetPuzzle())

	plantMapPositions := initializePlantMapPositions(garden)

	directions := []Position{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	regions := [][]Position{}

	for plant, plantPositions := range plantMapPositions {
		visited := initializeVisited(garden)

		for _, position := range plantPositions {
			findRegions(garden, visited, &regions, directions, plant, position)
		}
	}

	for _, region := range regions {
		if len(region) > 0 {
			result += len(region) * calcRegionPerimeter(region, garden, directions)
		}
	}
}

func calcRegionPerimeter(region []Position, garden [][]string, directions []Position) int {
	perimeter := 0

	plant := garden[region[0].Y][region[0].X]

	for _, pos := range region {
		for _, direction := range directions {
			around := Position{
				Y: pos.Y + direction.Y,
				X: pos.X + direction.X,
			}

			if around.Y < 0 || around.Y >= len(garden) || around.X < 0 || around.X >= len(garden[0]) || garden[around.Y][around.X] != plant {
				perimeter++
			}
		}
	}

	return perimeter
}

func findRegions(garden [][]string, visited [][]bool, regions *[][]Position, directions []Position, plant string, currPos Position) {
	region := []Position{}

	walk(garden, visited, &region, directions, plant, currPos)

	*regions = append(*regions, region)
}

func walk(garden [][]string, visited [][]bool, region *[]Position, directions []Position, plant string, currPos Position) {
	if currPos.Y < 0 || currPos.Y >= len(garden) || currPos.X < 0 || currPos.X >= len(garden[0]) {
		return
	}

	if visited[currPos.Y][currPos.X] {
		return
	}

	if garden[currPos.Y][currPos.X] != plant {
		return
	}

	visited[currPos.Y][currPos.X] = true

	*region = append(*region, Position{currPos.Y, currPos.X})

	for _, direction := range directions {
		nextPosition := Position{
			Y: currPos.Y + direction.Y,
			X: currPos.X + direction.X,
		}

		walk(garden, visited, region, directions, plant, nextPosition)
	}
}

func initializeVisited(garden [][]string) [][]bool {
	visited := make([][]bool, len(garden))

	for row := range visited {
		visited[row] = make([]bool, len(garden[0]))
	}

	return visited
}

func initializePlantMapPositions(garden [][]string) map[string][]Position {
	plantMapPositions := make(map[string][]Position)

	for y, gardenRow := range garden {
		for x, gardenCol := range gardenRow {
			if _, ok := plantMapPositions[gardenCol]; !ok {
				plantMapPositions[gardenCol] = []Position{}
			}

			plantMapPositions[gardenCol] = append(plantMapPositions[gardenCol], Position{y, x})
		}
	}

	return plantMapPositions
}

func initializeGardenPlotRegionsMap(garden [][]string) map[string][][]Position {
	gardenPlotRegionsMap := make(map[string][][]Position)

	for _, gardenRow := range garden {
		for _, gardenCol := range gardenRow {
			if _, ok := gardenPlotRegionsMap[gardenCol]; !ok {
				gardenPlotRegionsMap[gardenCol] = [][]Position{}
			}
		}
	}

	return gardenPlotRegionsMap
}

func parsePuzzle(puzzle string) [][]string {
	puzzleArr := strings.Split(puzzle, "\n")

	parsedPuzzle := make([][]string, len(puzzleArr))

	for i := range parsedPuzzle {
		parsedPuzzle[i] = make([]string, len(puzzleArr[0]))
	}

	for row := range puzzleArr {
		for col, plant := range puzzleArr[row] {
			parsedPuzzle[row][col] = string(plant)
		}
	}

	return parsedPuzzle
}
