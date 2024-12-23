package challenge

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt1() {
	puzzle := utils.GetPuzzle()

	puzzleReports := strings.Split(puzzle, "\n")

	safeReports := 0

	for _, report := range puzzleReports {

		reportValues := strings.Split(report, " ")

		isSafe := true

		prev, err := strconv.Atoi(reportValues[0])
		if err != nil {
			log.Fatal(err)
		}

		next, err := strconv.Atoi(reportValues[1])
		if err != nil {
			log.Fatal(err)
		}

		isIncrementing := prev < next

		for i := 1; i < len(reportValues); i++ {
			value := reportValues[i]

			current, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			difference := int(math.Abs(float64(prev - current)))

			if (isIncrementing && prev >= current) || (!isIncrementing && prev <= current) || (difference < 1 || difference > 3) {
				isSafe = false
				break
			}

			prev = current
		}

		if isSafe {
			safeReports++
		}
	}

	fmt.Printf("Part 1 -> %d", safeReports)
}
