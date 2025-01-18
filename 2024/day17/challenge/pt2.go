package challenge

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	registers, instructions := parsePuzzle(utils.GetPuzzle())

	sliceToString := func(s int, instructions []int) string {
		stringSlice := []string{}

		for _, out := range instructions[len(instructions)-s:] {
			stringSlice = append(stringSlice, strconv.Itoa(out))
		}

		return strings.Join(stringSlice, ",")
	}

	instructionsString := sliceToString(len(instructions), instructions)

	a := 0
	extractFromEnd := 1

	for {
		registers.A = a

		output := findProgramOutput(registers, instructions)

		if output == instructionsString {
			result = a
			break
		} else if output == sliceToString(extractFromEnd, instructions) {
			a <<= 3
			extractFromEnd++
		} else {
			a++
		}
	}
}
