package challenge

import (
	"fmt"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

type Position struct {
	y int
	x int
}

func Pt1() {
	antennasPositionMap, puzzle := parsePuzzle(utils.GetPuzzle())
	antiNodesPositionsQty := findAntiNodesUniquePositions(antennasPositionMap, puzzle)

	fmt.Printf("Part 1 -> %d", antiNodesPositionsQty)
}

func findAntiNodesUniquePositions(antennasPositionMap map[string][]Position, puzzle []string) int {
	antiNodePositions := make(map[string]struct{})
	columnLen := len(puzzle)
	rowLen := len(puzzle[0])

	for _, v := range antennasPositionMap {
		for i, antennaPosition := range v {
			for j, nextAntennaPosition := range v {
				if i == j && len(v) <= 1 {
					continue
				}

				x := antennaPosition.x - nextAntennaPosition.x
				y := antennaPosition.y - nextAntennaPosition.y

				antiNodePosition := Position{
					y: antennaPosition.y + y,
					x: antennaPosition.x + x,
				}

				if isWithinTheBounds(antiNodePosition, columnLen, rowLen) {
					antiNodePositionComparable := fmt.Sprintf("%d|%d", antiNodePosition.y, antiNodePosition.x)
					antiNodePositions[antiNodePositionComparable] = struct{}{}
				}
			}
		}
	}

	return len(antiNodePositions)
}

func isWithinTheBounds(pos Position, columnLen, rowLen int) bool {
	return pos.y >= 0 && pos.y < columnLen && pos.x >= 0 && pos.x < rowLen
}

func parsePuzzle(puzzle string) (map[string][]Position, []string) {
	antennasPositionMap := make(map[string][]Position)

	puzzleArr := strings.Split(puzzle, "\n")

	for y := range puzzleArr {
		for x, antenna := range puzzleArr[y] {
			if antenna == '.' {
				continue
			}
			if _, ok := antennasPositionMap[string(antenna)]; !ok {
				antennasPositionMap[string(antenna)] = make([]Position, 0)
			}

			antennasPositionMap[string(antenna)] = append(antennasPositionMap[string(antenna)], Position{y, x})
		}
	}

	return antennasPositionMap, puzzleArr
}
