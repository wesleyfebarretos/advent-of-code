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
		fmt.Printf("\nPart 2 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	clawMachines := parsePuzzlePt2(utils.GetPuzzle())

	pressTheButtonsPt2(clawMachines)

	result = calcClawMachinesTokens(clawMachines)
}

func pressTheButtonsPt2(clawMachines []ClawMachine) {
	for i, clawMachine := range clawMachines {
		determinant := (clawMachine.ButtonA.X * clawMachine.ButtonB.Y) - (clawMachine.ButtonA.Y * clawMachine.ButtonB.X)

		//  No way to reach the prize
		if determinant == 0 {
			continue
		}

		buttonAPresses := (float64(clawMachine.ButtonB.Y*clawMachine.Prize.X) - float64(clawMachine.ButtonB.X*clawMachine.Prize.Y)) / float64(determinant)
		buttonBPresses := (float64(clawMachine.ButtonA.X*clawMachine.Prize.Y) - float64(clawMachine.ButtonA.Y*clawMachine.Prize.X)) / float64(determinant)

		if buttonAPresses < 0 || buttonBPresses < 0 || math.Mod(buttonAPresses, 1) != 0 || math.Mod(buttonBPresses, 1) != 0 {
			continue
		}

		clawMachines[i].ButtonA.Pressed = int(math.Ceil(buttonAPresses))
		clawMachines[i].ButtonB.Pressed = int(math.Ceil(buttonBPresses))
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
				X: x3 + 10_000_000_000_000,
				Y: y3 + 10_000_000_000_000,
			},
		}
	}

	return clawMachines
}
