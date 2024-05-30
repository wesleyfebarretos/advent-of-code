package challenge1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	report := getReport()
	for i := range report {
		extrapolateSequences(i, &report)
	}
	fmt.Println(sumReportsExtrapolation(report))
}

func getReport() [][]int {
	report, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	reports := strings.Split(strings.TrimSpace(string(report)), "\n")

	var splitedReports [][]string

	for _, r := range reports {
		splitedReports = append(splitedReports, strings.Split(r, " "))
	}

	var intReports [][]int

	for _, report := range splitedReports {
		var intReport []int
		for _, r := range report {
			intr, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}
			intReport = append(intReport, intr)
		}
		intReports = append(intReports, intReport)
	}

	return intReports
}

func extrapolateSequences(index int, report *[][]int) {
	sequence := (*report)[index]

	var sequences [][]int
	sequences = append(sequences, sequence)
	for i := 0; i < len(sequences); i++ {
		exit := createSequences(&sequences, i, len(sequences[i])-1)
		if exit {
			extrapolate(&sequences)
			break
		}
	}

	for _, v := range sequences {
		(*report)[index] = v
		break
	}
}

func createSequences(sequences *[][]int, currentIndex, newSequenceLen int) bool {
	currentSequence := (*sequences)[currentIndex]
	var newSequence []int

	for i := 0; i < newSequenceLen; i++ {
		currSeqValue := currentSequence[i]
		currSeqNextValue := currentSequence[i+1]
		currSeqValueDiff := currSeqNextValue - currSeqValue

		newSequence = append(newSequence, currSeqValueDiff)
	}

	*sequences = append((*sequences), newSequence)

	return canExit(newSequence)
}

func extrapolate(sequences *[][]int) {
	for i := len(*sequences) - 1; i >= 0; i-- {
		if i == len(*sequences)-1 {
			(*sequences)[i] = append((*sequences)[i], 0)
			continue
		}
		beforeSeq := (*sequences)[i+1]
		beforeSeqLastValue := beforeSeq[len(beforeSeq)-1]
		currSeq := (*sequences)[i]
		currSeqLastValue := currSeq[len(currSeq)-1]
		(*sequences)[i] = append((*sequences)[i], beforeSeqLastValue+currSeqLastValue)
	}
}

func sumReportsExtrapolation(report [][]int) int {
	sum := 0
	for _, r := range report {
		sum += r[len(r)-1]
	}
	return sum
}

func canExit(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}
	return true
}
