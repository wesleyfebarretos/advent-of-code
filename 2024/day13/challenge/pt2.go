package challenge

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %d, runned in %s", result, time.Since(t))
	}(time.Now())

	clawMachines := parsePuzzlePt2(utils.GetTestPuzzle())

	pressTheButtonsPt2(clawMachines)

	result = calcClawMachinesTokens(clawMachines)
}

func pressTheButtonsPt2(clawMachines []ClawMachine) {
outer:
	for i, clawMachine := range clawMachines {
		buttonAPressedCounter := 0
		for {

			//  X = (ButtonA.X)n + (ButtonB.X)m
			//  Y = (ButtonA.Y)n + (ButtonB.Y)m
			//  So im removing from prize the value of amount times of A button was pressed
			//  and divided by buttonB X and Y to check the amount o presses needed to reach the prize
			buttonBXPress := float64((clawMachine.Prize.X - clawMachine.ButtonA.X*buttonAPressedCounter)) / float64(clawMachine.ButtonB.X)
			buttonBYPress := float64((clawMachine.Prize.Y - clawMachine.ButtonA.Y*buttonAPressedCounter)) / float64(clawMachine.ButtonB.Y)

			if buttonBXPress < 1 || buttonBYPress < 1 {
				continue outer
			}

			//  if they equal it means that can reach the amount of needed values to win a prize
			if math.Mod(buttonBXPress, 1) == 0 && math.Mod(buttonBYPress, 1) == 0 && buttonBXPress == buttonBYPress {
				clawMachines[i].ButtonA.Pressed = buttonAPressedCounter
				clawMachines[i].ButtonB.Pressed = int(buttonBXPress)
				continue outer
			}

			buttonAPressedCounter++
		}
	}
}

func parsePuzzlePt2(puzzle string) []ClawMachine {
	puzzleRows := strings.Split(puzzle, "\n\n")

	clawMachines := make([]ClawMachine, len(puzzleRows))

	for i, row := range puzzleRows {
		lines := strings.Split(row, "\n")

		x1, y1 := getXAndYValues(lines[0])
		x2, y2 := getXAndYValues(lines[1])
		x3, y3 := getXAndYValues(lines[2])

		clawMachines[i] = ClawMachine{
			ButtonA: Button{
				X: x1,
				Y: y1,
			},
			ButtonB: Button{
				X: x2,
				Y: y2,
			},
			Prize: Prize{
				X: x3 + 100_000_000_000_00,
				Y: y3 + 100_000_000_000_00,
			},
		}
	}

	return clawMachines
}
