package challenge

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	puzzle := utils.GetPuzzle()

	reports := parsePuzzle(puzzle)

	safeReports := 0

	for _, report := range reports {

		isSafe := findReportViolations(report) == 0
		problemDampenerTry := true

		if !isSafe && problemDampenerTry {
			isSafe = problemDampenerHandler(report)
			problemDampenerTry = false
		}

		if isSafe {
			safeReports++
		}
	}

	fmt.Printf("Part 2 -> %d", safeReports)
}

func findReportViolations(report []int) int {
	violations := 0

	current, next := report[0], report[1]

	isIncrementing := current < next

	for i := 0; i < len(report)-1; i++ {

		current := report[i]
		next := report[i+1]

		difference := int(math.Abs(float64(current - next)))

		if isUnsafe(isIncrementing, current, next, difference) {
			violations++
		}
	}

	return violations
}

func problemDampenerHandler(report []int) bool {
	removeIdx := -1
	violations := math.MaxInt64

	for i := range report {

		for j := 0; j < len(report)-1; j++ {
			subReportArr := slices.Concat(report[:i], report[i+1:])

			v := findReportViolations(subReportArr)

			if v < violations {
				violations = v
				removeIdx = i
			}
		}

	}

	return findReportViolations(slices.Concat(report[:removeIdx], report[removeIdx+1:])) == 0
}

func parsePuzzle(puzzle string) [][]int {
	puzzleReports := strings.Split(puzzle, "\n")

	reports := make([][]int, len(puzzleReports))

	for i, report := range puzzleReports {
		levels := strings.Fields(report)

		reports[i] = make([]int, len(levels))

		for j, v := range levels {
			num, err := strconv.Atoi(v)

			if err != nil {
				log.Fatal(err)
			}

			reports[i][j] = num
		}
	}

	return reports
}

func isUnsafe(isIncrementing bool, current, next, difference int) bool {
	return (isIncrementing && current >= next) || (!isIncrementing && current <= next) || (difference < 1 || difference > 3)
}
