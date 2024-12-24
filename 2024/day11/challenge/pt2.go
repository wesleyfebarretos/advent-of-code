package challenge

import (
	"fmt"
	"strconv"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	stones := parsePuzzle(utils.GetPuzzle())

	stonesMap := map[string]int{}

	for _, stone := range stones {
		if _, ok := stonesMap[stone]; !ok {
			stonesMap[stone] = 1
		}
	}

	for i := 0; i < 75; i++ {
		newStonesMap := map[string]int{}

		for stone, value := range stonesMap {
			for _, newStone := range blinkStone(stone) {

				if _, ok := newStonesMap[newStone]; !ok {
					newStonesMap[newStone] = 0
				}

				newStonesMap[newStone] += value
			}
		}

		stonesMap = newStonesMap
	}

	totalOfStones := 0

	for _, value := range stonesMap {
		totalOfStones += value
	}

	fmt.Printf("Part 2 -> %d", totalOfStones)
}

func blinkStone(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}

	if (len(stone) & 1) == 0 {
		middle := len(stone) / 2

		return []string{stone[0:middle], removeZeroesFromLeft(stone[middle:])}
	}

	stoneAsNum, _ := strconv.Atoi(stone)

	return []string{strconv.Itoa(stoneAsNum * 2024)}
}
