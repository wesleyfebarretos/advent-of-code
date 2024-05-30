package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SecondChallenge() {
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

	time := strings.Join(strings.Split(reg.ReplaceAllString(splitFile[0], ","), ",")[1:], "")
	distance := strings.Join(strings.Split(reg.ReplaceAllString(splitFile[1], ","), ",")[1:], "")

	totalTime, err := strconv.Atoi(time)
	if err != nil {
		panic(err)
	}
	totalDistance, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}

	recordBeatingQty := 0

	for i := 1; i <= totalTime; i++ {
		traveledDistance := i * (totalTime - i)
		if traveledDistance > totalDistance {
			recordBeatingQty++
		}
	}
	fmt.Println(recordBeatingQty)
}
