package challenge

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		_ = clipboard.WriteAll(fmt.Sprintf("%v", result))
		fmt.Printf("\nPart 2 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	raceMap := parsePuzzle(utils.GetPuzzle())

	start, end := findStartAndEnd(raceMap)

	distsFromStart := dijkstra(start, raceMap)
	distsFromEnd := dijkstra(end, raceMap)

	for sn, sv := range distsFromStart {
		for en, ev := range distsFromEnd {
			manhattanDist := manhattanDistance(sn, en)
			if manhattanDist <= 20 && sv+ev+manhattanDist <= distsFromStart[end]-100 {
				result++
			}
		}
	}
}
