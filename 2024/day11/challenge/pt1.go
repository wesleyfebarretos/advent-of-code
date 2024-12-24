package challenge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt1() {
	stones := parsePuzzle(utils.GetPuzzle())

	for i := 0; i < 25; i++ {
		innerStones := []string{}

		for _, stone := range stones {
			if stone == "0" {
				innerStones = append(innerStones, "1")
			} else if (len(stone) & 1) == 0 {
				middle := len(stone) / 2

				innerStones = append(innerStones, stone[0:middle])

				innerStones = append(innerStones, removeZeroesFromLeft(stone[middle:]))

			} else {
				stoneAsNum, _ := strconv.Atoi(stone)

				innerStones = append(innerStones, strconv.Itoa(stoneAsNum*2024))
			}
		}

		stones = innerStones
	}

	fmt.Printf("Part 1 -> %d", len(stones))
}

func removeZeroesFromLeft(num string) string {
	numWithoutZeroes := strings.TrimLeft(num, "0")

	if numWithoutZeroes == "" {
		return "0"
	}

	return numWithoutZeroes
}

func parsePuzzle(puzzle string) []string {
	reg := regexp.MustCompile(`\s+`)
	return strings.Split(reg.ReplaceAllString(puzzle, " "), " ")
}
