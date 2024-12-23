package challenge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt1() {
	puzzle := parsePuzzle(utils.GetPuzzle())

	result := int64(0)

	for y := range puzzle {
		target := puzzle[y][0]
		startValue := puzzle[y][1]
		startIdx := 1
		if findTruthyEquation(startValue, target, startIdx, puzzle[y]) {
			result += puzzle[y][0]
		}
	}

	fmt.Println("Part 1 -> ", result)
}

func findTruthyEquation(currentValue, target int64, currentIdx int, numbers []int64) bool {
	if currentIdx == len(numbers)-1 {
		return currentValue == target
	}

	if currentValue > target {
		return false
	}

	return findTruthyEquation(currentValue*numbers[currentIdx+1], target, currentIdx+1, numbers) || findTruthyEquation(currentValue+numbers[currentIdx+1], target, currentIdx+1, numbers)
}

func parsePuzzle(puzzle string) [][]int64 {
	parsedPuzzle := strings.Split(puzzle, "\n")

	matrix := make([][]int64, len(parsedPuzzle))

	reg := regexp.MustCompile(`[^0-9]+`)

	for y := range parsedPuzzle {

		numbers := strings.Split(reg.ReplaceAllString(parsedPuzzle[y], ","), ",")

		matrix[y] = make([]int64, len(numbers))

		for i, n := range numbers {
			intn, _ := strconv.Atoi(n)
			matrix[y][i] = int64(intn)
		}

	}

	return matrix
}

// Other approach
func combinatorialAnalysis(puzzle [][]int64) {
	sum := int64(0)

	for _, row := range puzzle {
		expected := row[0]
		combinations := []int64{row[1]}

		for y := range row {
			if y < 2 {
				continue
			}

			newCombinations := []int64{}

			num := row[y]

			for _, cnum := range combinations {
				sum := num + cnum
				if sum <= expected {
					newCombinations = append(newCombinations, sum)
				}

				mult := num * cnum

				if mult <= expected {
					newCombinations = append(newCombinations, mult)
				}

			}

			combinations = newCombinations
		}

		for _, cnum := range combinations {
			if cnum == expected {
				sum += cnum
				break
			}
		}
	}

	fmt.Println(sum)
}
