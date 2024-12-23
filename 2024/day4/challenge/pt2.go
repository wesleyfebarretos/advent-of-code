package challenge

import (
	"fmt"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

type APosition struct {
	y int
	x int
}

func Pt2() {
	puzzle := getMatrixOfLetters(utils.GetPuzzle())

	lookDirections := Directions{
		{-1, -1}, // north-weast
		{-1, 1},  // noth-east
		{1, 1},   // south-east
		{1, -1},  // south-weast
	}

	aPositions := getSliceOfAPositions(puzzle)

	masInXQty := findMASInXQty(puzzle, aPositions, lookDirections)

	fmt.Printf("Part 2 -> %d", masInXQty)
}

func findMASInXQty(puzzle Matrix, aPositions []APosition, lookDirections Directions) int {
	masInXTotal := 0
	for _, aPosition := range aPositions {
		masInXTotal += validMasXFormat(puzzle, aPosition, lookDirections)
	}

	return masInXTotal
}

func validMasXFormat(puzzle Matrix, aPosition APosition, lookDirections [][]int) int {

	letterMap := map[string]int{
		"M": 0,
		"S": 0,
	}

	for _, direction := range lookDirections {

		position := aPosition

		position.y += direction[0]
		position.x += direction[1]

		if position.y < 0 || position.y >= len(puzzle) {
			return 0
		}

		if position.x < 0 || position.x >= len(puzzle[0]) {
			return 0
		}

		letter := puzzle[position.y][position.x]

		//  check opposite diagonal
		if aPosition.y-direction[0] >= 0 && aPosition.y-direction[0] < len(puzzle) && aPosition.x-direction[1] >= 0 && aPosition.x-direction[1] < len(puzzle[0]) {
			oppositeLetter := puzzle[aPosition.y-direction[0]][aPosition.x-direction[1]]
			if letter == oppositeLetter {
				return 0
			}
		}

		if _, ok := letterMap[letter]; !ok {
			return 0
		}

		letterMap[letter]++
	}

	if letterMap["M"] == 2 && letterMap["S"] == 2 {
		return 1
	}

	return 0
}

func getSliceOfAPositions(puzzle [][]string) []APosition {
	aPositions := []APosition{}

	for i, line := range puzzle {
		for j, c := range line {
			if c == "A" {
				aPositions = append(aPositions, APosition{y: i, x: j})
			}
		}
	}

	return aPositions
}
