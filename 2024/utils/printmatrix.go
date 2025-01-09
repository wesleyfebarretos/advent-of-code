package utils

import (
	"fmt"
	"strings"
)

func PrintMatrix[S any](space [][]S) {
	s := strings.Builder{}

	for _, row := range space {
		for _, num := range row {
			s.WriteString(fmt.Sprintf("%v", num))
		}
		s.WriteString("\n")
	}

	fmt.Println(s.String())
}
