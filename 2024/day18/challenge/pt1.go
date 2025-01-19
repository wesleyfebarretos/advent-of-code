package challenge

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

const (
	GRID_LEN      = 71
	BYTES_TO_FALL = 1024
	END_X         = 70
	END_Y         = 70
)

type Position struct {
	Y    int
	X    int
	Cost int
}

func Pt1() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 1 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	bytesPosition := parseInput(utils.GetPuzzle())

	memorySpace := buildMatrix()

	for i, row := range bytesPosition {
		if i > BYTES_TO_FALL {
			break
		}

		y, x := row[0], row[1]

		memorySpace[y][x] = "#"
	}

	seen := make([][]bool, GRID_LEN)

	for i := range seen {
		seen[i] = make([]bool, GRID_LEN)
	}

	list := list.New()

	list.PushBack(Position{0, 0, 0})

	shortPath := math.MaxInt

	for list.Len() > 0 {
		pos := list.Remove(list.Front()).(Position)

		y, x, cost := pos.Y, pos.X, pos.Cost

		for _, adj := range neighbours(y, x, memorySpace) {
			adjy, adjx := adj[0], adj[1]

			if seen[adjy][adjx] {
				continue
			}

			seen[adjy][adjx] = true

			if adjy == END_Y && adjx == END_X {
				shortPath = int(math.Min(float64(shortPath), float64(cost+1)))
				break
			}

			list.PushBack(Position{adjy, adjx, cost + 1})
		}
	}

	result = shortPath
}

func neighbours(y, x int, memorySpace [][]string) [][2]int {
	neighbours := [][2]int{}

	for _, dir := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		dy, dx := dir[0], dir[1]

		ny, nx := y+dy, x+dx

		if ny < 0 || ny >= len(memorySpace) || nx < 0 || nx >= len(memorySpace[0]) || memorySpace[ny][nx] == "#" {
			continue
		}

		neighbours = append(neighbours, [2]int{ny, nx})
	}

	return neighbours
}

func buildMatrix() [][]string {
	matrix := make([][]string, GRID_LEN)

	for i := range matrix {
		matrix[i] = make([]string, GRID_LEN)
		for y := range matrix[i] {
			matrix[i][y] = "."
		}
	}

	return matrix
}

func parseInput(puzzle string) [][2]int {
	sliceInput := strings.Split(puzzle, "\n")

	bytePositions := make([][2]int, len(sliceInput))

	for i, bp := range sliceInput {
		bpSlice := strings.Split(bp, ",")

		x, _ := strconv.Atoi(bpSlice[0])
		y, _ := strconv.Atoi(bpSlice[1])

		bytePositions[i][0] = y
		bytePositions[i][1] = x
	}

	return bytePositions
}
