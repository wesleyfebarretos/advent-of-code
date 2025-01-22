package challenge

import (
	"fmt"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	patterns, designs := parsePuzzle(utils.GetPuzzle())

	possibleArrangements := make(map[string]int)

	for _, d := range designs {
		result += countPossiblesArrangements(d, patterns, possibleArrangements)
	}
}

func countPossiblesArrangements(design string, patterns []string, possibleArrangements map[string]int) int {
	//  DP Memoization
	//  the patterns could walk through the same way a lot of times
	//  so as the algorithm suggest, just cache it and return the value always that reach that points
	if _, ok := possibleArrangements[design]; ok {
		return possibleArrangements[design]
	}

	if len(design) == 0 {
		return 1
	}

	result := 0
	for _, p := range patterns {

		if len(p) > len(design) {
			continue
		}

		if strings.HasPrefix(design, p) {
			result += countPossiblesArrangements(strings.TrimPrefix(design, p), patterns, possibleArrangements)
		}

	}

	possibleArrangements[design] = result

	return result
}
