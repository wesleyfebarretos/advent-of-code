package challenge

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt1() {
	puzzle := utils.GetPuzzle()

	locations := strings.Split(puzzle, "\n")

	localtions1 := []int{}
	localtions2 := []int{}

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

	sort.Slice(localtions1, func(i, j int) bool {
		return localtions1[i] < localtions1[j]
	})

	sort.Slice(localtions2, func(i, j int) bool {
		return localtions2[i] < localtions2[j]
	})

	totalDistance := 0

	for i := 0; i < len(localtions1); i++ {
		totalDistance += int(math.Abs(float64(localtions1[i] - localtions2[i])))
	}

	fmt.Printf("Part 1 -> %d", totalDistance)
}
