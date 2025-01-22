package challenge

import (
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

	for _, d := range designs {
		seen := map[string]bool{}

		if isAValidDesign(d, patterns, seen) {
			result += 1
		}
	}
}

func isAValidDesign(design string, patterns []string, seen map[string]bool) bool {
	if seen[design] {
		return false
	}

	seen[design] = true

	if len(design) == 0 {
		return true
	}

	for _, p := range patterns {
		if len(p) > len(design) {
			continue
		}

		if design[:len(p)] == p {
			if isAValidDesign(design[len(p):], patterns, seen) {
				return true
			}
		}
	}

	return false
}

func parsePuzzle(puzzle string) ([]string, []string) {
	puzzleSlice := strings.Split(puzzle, "\n\n")

	towelPatterns, designs := puzzleSlice[0], puzzleSlice[1]

	towelPatternsSlice := strings.Split(towelPatterns, ", ")

	designsSlice := strings.Split(designs, "\n")

	return towelPatternsSlice, designsSlice
}
