package challenge

import (
	"fmt"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	antennasPositionMap, puzzle := parsePuzzle(utils.GetPuzzle())
	antiNodesPositionsQty := findAntiNodes(antennasPositionMap, puzzle)

	fmt.Printf("Part 2 -> %d", antiNodesPositionsQty)
}

func findAntiNodes(antennasPositionMap map[string][]Position, puzzle []string) int {
	antiNodePositions := []Position{}

	for _, v := range antennasPositionMap {
		for i, antennaPosition := range v {
			for j, nextAntennaPosition := range v {
				if i == j || len(v) <= 1 {
					continue
				}

				diffX := antennaPosition.x - nextAntennaPosition.x
				diffY := antennaPosition.y - nextAntennaPosition.y

				walk(nextAntennaPosition, diffY, diffX, puzzle, &antiNodePositions)
			}
		}
	}

	antiNodesMap := make(map[string]struct{})

	for _, v := range antiNodePositions {
		antiNodesMap[fmt.Sprintf("y=%dx=%d", v.y, v.x)] = struct{}{}
	}

	return len(antiNodesMap)
}

func walk(currentPosition Position, diffY, diffX int, puzzle []string, antiNodePositions *[]Position) {
	antiNodePosition := Position{
		y: currentPosition.y + diffY,
		x: currentPosition.x + diffX,
	}

	if antiNodePosition.x < 0 || antiNodePosition.x >= len(puzzle[0]) || antiNodePosition.y < 0 || antiNodePosition.y >= len(puzzle) {
		return
	}

	*antiNodePositions = append(*antiNodePositions, antiNodePosition)

	walk(antiNodePosition, diffY, diffX, puzzle, antiNodePositions)
}
