package challenge

import (
	"container/list"
	"fmt"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := ""

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %v, runned in %s\n", result, time.Since(t))
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

	seen := buildSeen()

	list := list.New()

	list.PushBack(Position{0, 0, 0})

	for i := BYTES_TO_FALL + 1; i < len(bytesPosition); i++ {
		found := false

		fallNextByteInMemorySpace(i, memorySpace, bytesPosition)

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
					found = true
				}

				list.PushBack(Position{adjy, adjx, cost + 1})
			}
		}

		list.Init()
		list.PushBack(Position{0, 0, 0})
		seen = buildSeen()

		if !found {
			byte := bytesPosition[3010]
			result = fmt.Sprintf("%d,%d", byte[1], byte[0])
			break
		}
	}
}

func fallNextByteInMemorySpace(index int, memorySpace [][]string, bytesPosition [][2]int) {
	pos := bytesPosition[index]

	y, x := pos[0], pos[1]

	memorySpace[y][x] = "#"
}

func buildSeen() [][]bool {
	seen := make([][]bool, GRID_LEN)

	for i := range seen {
		seen[i] = make([]bool, GRID_LEN)
	}

	return seen
}
