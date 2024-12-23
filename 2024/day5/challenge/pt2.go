package challenge

import (
	"fmt"
	"slices"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	pageOrderingMap, pageNumbers := getPageOrderingAndPageNumbers(utils.GetPuzzle())

	middlePageTotal := 0

	for _, pageNumber := range pageNumbers {
		if !isOrderValid(pageOrderingMap, pageNumber) {
			slices.SortFunc(pageNumber, comparePages(pageOrderingMap))
			middlePageTotal += pageNumber[len(pageNumber)/2]
		}
	}

	fmt.Printf("Part 2 -> %d", middlePageTotal)
}

func comparePages(orderingMap map[int][]int) func(int, int) int {
	return func(a, b int) int {

		if slices.Contains(orderingMap[a], b) {
			return -1
		}

		return 1
	}
}
