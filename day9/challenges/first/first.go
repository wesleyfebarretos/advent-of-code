package challenge1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	// TODO: TRANSFORMAR O RETORNO DE REPORT E UM [][]int
	// CRIAR AS CAMADAS ADICIONAIS EM CADA INDEX DO ARRAY DE ACORDO COM O ENUNCIADO
	// ESSA CAMADA SÓ É ADICIONADA QUANDO TODOS OS VALORES DO ULTIMO ARRAY SAO 0
	report := getReport()
	for i := range report {
		extrapolateSequences(i, report)
	}
}

func getReport() [][]string {
	report, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	reports := strings.Split(strings.TrimSpace(string(report)), "\n")

	var splitedReports [][]string

	for _, r := range reports {
		splitedReports = append(splitedReports, strings.Split(r, " "))
	}

	return splitedReports
}

func extrapolateSequences(index int, report [][]string) {
	sequence := report[index]
	var extrapolates [][]string
	extrapolates = append(extrapolates, sequence)
	for i := 0; i < len(extrapolates); i++ {
		newSequence, exit := createSequences(extrapolates, i, len(extrapolates[i])-1)
		extrapolates = append(extrapolates, newSequence)

		if exit {
			break
		}
	}
	fmt.Println(extrapolates)
}

func createSequences(extrapolates [][]string, currentIndex, newSequenceLen int) ([]string, bool) {
	currentSequence := extrapolates[currentIndex]
	var newSequence []string

	for i := 0; i < newSequenceLen; i++ {
		currSeqValue, err := strconv.Atoi(currentSequence[i])
		if err != nil {
			panic(err)
		}
		currSeqNextValue, err := strconv.Atoi(currentSequence[i+1])
		if err != nil {
			panic(err)
		}
		currSeqValueDiff := currSeqValue - currSeqNextValue

		if currSeqValueDiff < 0 {
			currSeqValueDiff *= -1
		}
		newSequence = append(newSequence, fmt.Sprintf("%d", currSeqValueDiff))
	}

	if canExit(newSequence) {
		return newSequence, true
	}

	return newSequence, false
}

func canExit(sequence []string) bool {
	for _, v := range sequence {
		if v != "0" {
			return false
		}
	}
	return true
}
