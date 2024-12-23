package challenge

import (
	"fmt"
	"strconv"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	puzzle := parsePuzzle(utils.GetPuzzle())

	result := int64(0)

	for y := range puzzle {
		target := puzzle[y][0]
		startValue := puzzle[y][1]
		startIdx := 1
		if findTruthyEquation2(startValue, target, startIdx, puzzle[y]) {
			result += puzzle[y][0]
		}
	}

	fmt.Println("Part 2 -> ", result)
}

func findTruthyEquation2(currentValue, target int64, currentIdx int, numbers []int64) bool {
	if currentIdx == len(numbers)-1 {
		return currentValue == target
	}

	if currentValue > target {
		return false
	}

	return findTruthyEquation2(currentValue*numbers[currentIdx+1], target, currentIdx+1, numbers) || findTruthyEquation2(currentValue+numbers[currentIdx+1], target, currentIdx+1, numbers) || findTruthyEquation2(concat(currentValue, numbers[currentIdx+1]), target, currentIdx+1, numbers)
}

func concat(currentValue, nextValue int64) int64 {
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, nextValue))
	return int64(concat)
}
