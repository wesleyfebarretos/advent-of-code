package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FirstChallenge() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	splitFile := strings.Split(input, "\n")

	reg, err := regexp.Compile(`\s+`)
	if err != nil {
		panic(err)
	}

	times := strings.Split(reg.ReplaceAllString(splitFile[0], ","), ",")[1:]
	distances := strings.Split(reg.ReplaceAllString(splitFile[1], ","), ",")[1:]

	records := make([]int, 0)

	for i := range len(times) {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic(err)
		}

		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			panic(err)
		}

		matches := 0

		for i := 1; time > 0; i, time = i+1, time-1 {
			if i*(time-1) > distance {
				matches++
			}
		}

		records = append(records, matches)
	}
	waysToWin := 1

	for _, v := range records {
		waysToWin *= v
	}
	fmt.Println(waysToWin)
}
