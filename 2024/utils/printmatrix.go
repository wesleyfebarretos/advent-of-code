package utils

import (
	"fmt"
	"strings"
)

func PrintMatrix[S any](m [][]S) {
	s := strings.Builder{}

	for _, row := range m {
		for _, col := range row {
			s.WriteString(fmt.Sprintf("%v", col))
		}
		s.WriteString("\n")
	}

	fmt.Println(s.String())
}
