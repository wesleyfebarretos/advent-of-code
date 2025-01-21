package challenge

import (
	"container/list"
	"fmt"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt1() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 1 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	patterns, designs := parsePuzzle(utils.GetPuzzle())

	list := list.New()

	for _, d := range designs {
		list.PushBack([]string{d, d})
	}

	validDesigns := make(map[string]bool)

	for list.Len() > 0 {
		cdrd := list.Remove(list.Front()).([]string)

		cd, rd := cdrd[0], cdrd[1]

		for _, p := range patterns {

			if cd == p {
				validDesigns[rd] = true
				continue
			}

			if strings.Index(cd, p) == 0 {
				list.PushBack([]string{cd[len(p):], rd})
			}
		}
	}

	result = len(validDesigns)
}

func parsePuzzle(puzzle string) ([]string, []string) {
	puzzleSlice := strings.Split(puzzle, "\n\n")

	towelPatterns, designs := puzzleSlice[0], puzzleSlice[1]

	towelPatternsSlice := strings.Split(towelPatterns, ", ")

	designsSlice := strings.Split(designs, "\n")

	return towelPatternsSlice, designsSlice
}
