package challenge

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt2() {
	puzzle := utils.GetPuzzle()

	locations := strings.Split(puzzle, "\n")

	localtions1 := []int{}
	localtions2 := []int{}

	repeatedLocationsMap := make(map[int]int)

	reg := regexp.MustCompile(`\s+`)

	for _, row := range locations {
		rowSli := reg.Split(row, -1)

		ll, lr := rowSli[0], rowSli[1]

		l1, err := strconv.Atoi(string(ll))
		if err != nil {
			log.Fatal(err)
		}

		l2, err := strconv.Atoi(string(lr))
		if err != nil {
			log.Fatal(err)
		}

		localtions1 = append(localtions1, l1)
		localtions2 = append(localtions2, l2)
	}

	for _, l1 := range localtions1 {
		repeatedLocationsMap[l1] = 0
	}

	for _, l2 := range localtions2 {
		if _, ok := repeatedLocationsMap[l2]; ok {
			repeatedLocationsMap[l2] = repeatedLocationsMap[l2] + 1
		}
	}

	totalDistance := 0

	for k, v := range repeatedLocationsMap {
		totalDistance += k * v
	}

	fmt.Printf("Part 2 -> %d", totalDistance)
}
