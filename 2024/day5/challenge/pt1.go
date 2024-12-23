package challenge

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt1() {
	pageOrderingMap, pageNumbers := getPageOrderingAndPageNumbers(utils.GetPuzzle())

	middlePageTotal := 0

	for _, pageNumber := range pageNumbers {
		if isOrderValid(pageOrderingMap, pageNumber) {
			middlePageTotal += pageNumber[len(pageNumber)/2]
		}
	}

	fmt.Printf("Part 1 -> %d", middlePageTotal)
}

func isOrderValid(pageOrderingMap map[int][]int, nums []int) bool {
	for j, num := range nums {

		if j+1 >= len(nums) {
			break
		}

		if !slices.Contains(pageOrderingMap[num], nums[j+1]) {
			return false
		}
	}

	return true
}

func getPageOrderingAndPageNumbers(puzzle string) (map[int][]int, [][]int) {
	parsedPuzzle := strings.Split(puzzle, "\n\n")

	pageOrderingMap := map[int][]int{}

	for _, line := range strings.Split(parsedPuzzle[0], "\n") {
		numbers := strings.Split(line, "|")

		number1, _ := strconv.Atoi(numbers[0])
		number2, _ := strconv.Atoi(numbers[1])

		if _, ok := pageOrderingMap[number1]; !ok {
			pageOrderingMap[number1] = []int{}
		}

		pageOrderingMap[number1] = append(pageOrderingMap[number1], number2)
	}

	puzzlePgNumbers := strings.Split(parsedPuzzle[1], "\n")

	pageNumbers := make([][]int, len(puzzlePgNumbers))

	for i, pgNumbers := range puzzlePgNumbers {
		numbers := strings.Split(pgNumbers, ",")

		pageNumbers[i] = make([]int, len(numbers))

		for j, n := range numbers {
			intn, _ := strconv.Atoi(n)

			pageNumbers[i][j] = intn
		}
	}

	return pageOrderingMap, pageNumbers
}
