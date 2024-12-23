package challenge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/wesleyfebarretos/advent-of-code/utils"
)

func Pt1() {
	puzzle := utils.GetPuzzle()

	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)

	mulInstructions := reg.FindAllString(puzzle, -1)
	total := 0

	reg2 := regexp.MustCompile(`[^\d,]+`)
	for _, mul := range mulInstructions {

		numbers := strings.Split(reg2.ReplaceAllString(mul, ""), ",")

		first, _ := strconv.Atoi(numbers[0])
		second, _ := strconv.Atoi(numbers[1])

		total += first * second
	}

	fmt.Printf("Part 1 -> %d", total)

}
