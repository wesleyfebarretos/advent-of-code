package challenge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Registers struct {
	A                  int
	B                  int
	C                  int
	InstructionPointer int
}

func Pt1() {
	result := ""

	defer func(t time.Time) {
		fmt.Printf("\nPart 1 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	registers, instructions := parsePuzzle(utils.GetPuzzle())

	result = findProgramOutput(registers, instructions)
}

func findProgramOutput(registers Registers, instructions []int) string {
	outputs := []int{}

	opMap := operationsMap(&registers, &outputs)

	for registers.InstructionPointer < len(instructions) || registers.A > 0 {
		opMap[instructions[registers.InstructionPointer]](instructions[registers.InstructionPointer+1])
	}

	outputsStr := []string{}

	for _, out := range outputs {
		outputsStr = append(outputsStr, strconv.Itoa(out))
	}

	return strings.Join(outputsStr, ",")
}

func operationsMap(registers *Registers, outputs *[]int) map[int]func(int) {
	opMap := make(map[int]func(int))

	opMap[0] = func(op int) {
		registers.A >>= getComboOperator(*registers, op)
		registers.InstructionPointer += 2
	}

	opMap[1] = func(op int) {
		registers.B ^= op
		registers.InstructionPointer += 2
	}

	opMap[2] = func(op int) {
		registers.B = getComboOperator(*registers, op) % 8
		registers.InstructionPointer += 2
	}

	opMap[3] = func(op int) {
		if registers.A == 0 {
			registers.InstructionPointer += 2
			return
		}

		if registers.InstructionPointer != op {
			registers.InstructionPointer = op
		} else {
			registers.InstructionPointer += 2
		}
	}

	opMap[4] = func(op int) {
		registers.B ^= registers.C
		registers.InstructionPointer += 2
	}

	opMap[5] = func(op int) {
		*outputs = append(*outputs, getComboOperator(*registers, op)%8)
		registers.InstructionPointer += 2
	}

	opMap[6] = func(op int) {
		registers.B = registers.A >> getComboOperator(*registers, op)
		registers.InstructionPointer += 2
	}

	opMap[7] = func(op int) {
		registers.C = registers.A >> getComboOperator(*registers, op)

		registers.InstructionPointer += 2
	}

	return opMap
}

func getComboOperator(registers Registers, op int) int {
	switch op {
	case 7:
		panic("not an valid operator")
	case 4:
		return registers.A
	case 5:
		return registers.B
	case 6:
		return registers.C
	default:
		return op
	}
}

func parsePuzzle(puzzle string) (Registers, []int) {
	puzzleSlice := strings.Split(puzzle, "\n\n")

	registersPuzzle, instructions := puzzleSlice[0], puzzleSlice[1]

	registersSlice := strings.Split(registersPuzzle, "\n")

	getRegisterValue := func(register string) int {
		reg := regexp.MustCompile(`[^0-9]+`)

		value := reg.ReplaceAllString(register, "")

		intv, _ := strconv.Atoi(value)

		return intv
	}

	registers := Registers{}

	registers.A = getRegisterValue(registersSlice[0])
	registers.B = getRegisterValue(registersSlice[1])
	registers.C = getRegisterValue(registersSlice[2])
	registers.InstructionPointer = 0

	reg := regexp.MustCompile(`[0-9]+`)

	instructionsSlice := reg.FindAllString(instructions, -1)

	intStructions := []int{}

	for _, instruction := range instructionsSlice {
		intStruction, _ := strconv.Atoi(instruction)

		intStructions = append(intStructions, intStruction)

	}

	return registers, intStructions
}
