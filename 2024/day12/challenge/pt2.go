package challenge

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("Part 2 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	garden := parsePuzzle(utils.GetPuzzle())

	plantMapPositions := initializePlantMapPositions(garden)

	directions := []Position{
		{-1, 0},
		{0, 1},
		{1, 0},
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
			result += len(region) * calcRegionSides(region, directions, garden)
		}
	}
}

func calcRegionSides(region []Position, directions []Position, garden [][]string) int {
	plant := garden[region[0].Y][region[0].X]

	sides := []Position{}

	rowsMap := make(map[int][]Position)

	for _, pos := range region {
		if rowsMap[pos.Y] == nil {
			rowsMap[pos.Y] = []Position{}
		}

		rowsMap[pos.Y] = append(rowsMap[pos.Y], pos)
	}

	fenchMap := make(map[string][]int)

	for _, row := range rowsMap {
		for _, pos := range row {
			for _, direction := range directions {

				edge := Position{
					Y: pos.Y + direction.Y,
					X: pos.X + direction.X,
				}

				if outOfBounds(edge, garden) || garden[edge.Y][edge.X] != plant {
					compar := getEdgeComparable(edge, pos)

					if fenchMap[compar] == nil {
						fenchMap[compar] = []int{}
					}

					if direction.Y != 0 {
						fenchMap[compar] = append(fenchMap[compar], edge.X)
					} else {
						fenchMap[compar] = append(fenchMap[compar], edge.Y)
					}
				}
			}
		}
	}

	sortAndRemoveDuplicates(fenchMap)

	for key, fenchs := range fenchMap {
		validFenchs := []int{}

		for len(fenchs) > 0 {
			fench := fenchs[0]

			fenchs = fenchs[1:]

			if len(validFenchs) == 0 {
				validFenchs = append(validFenchs, fench)
				continue
			}

			previous := validFenchs[len(validFenchs)-1]

			//  remove adjacent fench
			if fench-previous == 1 {
				validFenchs = validFenchs[:len(validFenchs)-1]
			}

			// put the new fench in place of old one
			validFenchs = append(validFenchs, fench)
		}

		for _, fench := range validFenchs {
			startComparIdx := strings.Index(key, "=")
			endComparIdx := strings.Index(key, "|")

			numb, _ := strconv.Atoi(key[startComparIdx+1 : endComparIdx])

			if strings.Contains(key, "col") {
				sides = append(sides, Position{Y: fench, X: numb})
			} else {
				sides = append(sides, Position{Y: numb, X: fench})
			}
		}
	}

	return len(sides)
}

func sortAndRemoveDuplicates(fenchMap map[string][]int) {
	for k, fenchs := range fenchMap {
		seem := make(map[int]struct{})
		for _, fench := range fenchs {
			seem[fench] = struct{}{}
		}

		uniqueFenchSlice := []int{}

		for k := range seem {
			uniqueFenchSlice = append(uniqueFenchSlice, k)
		}

		slices.Sort(uniqueFenchSlice)

		fenchMap[k] = uniqueFenchSlice
	}
}

func outOfBounds(edge Position, garden [][]string) bool {
	return edge.Y < 0 || edge.Y >= len(garden) || edge.X < 0 || edge.X >= len(garden[0])
}

func getEdgeComparable(edge, pos Position) string {
	comp := ""
	if edge.Y != pos.Y {
		comp = fmt.Sprintf("row=%d|%d", edge.Y, pos.Y)
	} else {
		comp = fmt.Sprintf("col=%d|%d", edge.X, pos.X)
	}

	return comp
}
