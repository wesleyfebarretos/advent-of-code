package challenge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	puzzle := removeDontIntervals(utils.GetPuzzle())

	reg1 := regexp.MustCompile(`mul\(\d+,\d+\)`)

	mulInstructions := reg1.FindAllString(puzzle, -1)
	total := 0

	reg2 := regexp.MustCompile(`[^\d,]+`)
	for _, mul := range mulInstructions {

		numbers := strings.Split(reg2.ReplaceAllString(mul, ""), ",")

		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])

		total += first * second
	}

	fmt.Printf("Part 2 -> %d", total)

}

func removeDontIntervals(puzzle string) string {
	sb := strings.Builder{}

	dontIdx := strings.Index(puzzle, "don't()")
	doIdx := strings.Index(puzzle, "do()")

	for dontIdx != -1 {

		if dontIdx < doIdx {
			sb.WriteString(puzzle[:dontIdx])
			puzzle = puzzle[doIdx:]
		} else if doIdx == -1 {
			puzzle = ""
		} else if dontIdx > doIdx {
			sb.WriteString(puzzle[:dontIdx])
			puzzle = puzzle[dontIdx:]
		}

		dontIdx = strings.Index(puzzle, "don't()")
		doIdx = strings.Index(puzzle, "do()")
	}

	return sb.String()
}
