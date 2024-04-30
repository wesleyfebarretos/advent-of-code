package challenge1

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	ground := strings.Split(strings.TrimSpace(string(input)), "\n")
	var steps [4]int

	startDepth, startPosition := getStartPosition(ground)

	ways := makeTheWay(startDepth, startPosition)

	for i := 0; i < len(ways); {
		currDepth, currPosition := ways[i][0], ways[i][1]
		exit := walk(ground, steps, currDepth, currPosition)
		if exit {
			i++
		}
	}
}

func walk(ground []string, steps [4]int, currDepth, currPosition int) (exit bool) {
	fmt.Println(currDepth, currPosition)
	return true
}

func verticalWalk(groundLen int, currDepth *int) bool {
	if *currDepth == 0 {
		*currDepth += 1
	}
	if *currDepth == 0 {
		*currDepth += 1
	}
	return true
}

func makeTheWay(startDepth, startPosition int) [4][2]int {
	var ways [4][2]int
	ways[0][0], ways[0][1] = startDepth-1, startPosition
	ways[1][0], ways[1][1] = startDepth+1, startPosition
	ways[2][0], ways[2][1] = startDepth, startPosition-1
	ways[3][0], ways[3][1] = startDepth, startPosition+1
	return ways
}

func getStartPosition(ground []string) (startDepth, startPosition int) {
	for i := range ground {
		startDepth = i
		for i2, v2 := range ground[i] {
			if v2 == rune('S') {
				startPosition = i2
				return
			}
		}
	}
	panic("start position not found")
}
