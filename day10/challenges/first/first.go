package challenge1

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Pipe struct {
	character string
	right     bool
	left      bool
	top       bool
	bottom    bool
	x         int
	y         int
}

func Run() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rows := getPipes(input)

	var startPipe Pipe

	for _, row := range rows {
		for _, pipe := range row {
			if pipe.character == "S" {
				startPipe = pipe
			}
		}
	}

	directions := map[string]map[string]int{
		"right": {
			"x": 1,
		},
		"left": {
			"x": -1,
		},
		"top": {
			"y": -1,
		},
		"bottom": {
			"y": 1,
		},
	}

	// fmt.Printf("StartPipe: %+v\n", startPipe)
	// Debug pipes
	// for _, row := range rows {
	// 	for _, pipes := range row {
	// 		fmt.Printf("%+v\n", pipes)
	// 	}
	// }

	for direction, coordinates := range directions {
		startDirection := direction
		for coordinate, value := range coordinates {
			field := reflect.ValueOf(startPipe).FieldByName(coordinate)
			startPosition := field.Int() + int64(value)
			fmt.Printf("Direction: %s value: %d\n", startDirection, startPosition)
		}
	}
	// TODO: Criar função para pegar posição inversa da direção atual
	// a função vai receber a direção atual e vai ficar modificando ela
	// até nao encontrar mais caminho e retornar falso pra quebrar o infinite loop
	// pra cada execução do loop inserir um valor a mais no array de movimentos
}

func getPipes(input []byte) [][]Pipe {
	ground := strings.Split(strings.TrimSpace(string(input)), "\n")

	rows := make([][]Pipe, len(ground))

	for i, row := range ground {
		for y, pipeRune := range row {
			pipe := string(pipeRune)
			switch pipe {
			case "|":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     false,
					left:      false,
					top:       true,
					bottom:    true,
					x:         i,
					y:         y,
				})
			case "-":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     true,
					left:      true,
					top:       false,
					bottom:    false,
					x:         i,
					y:         y,
				})
			case "J":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     false,
					left:      true,
					top:       true,
					bottom:    false,
					x:         i,
					y:         y,
				})
			case "L":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     true,
					left:      false,
					top:       true,
					bottom:    false,
					x:         i,
					y:         y,
				})
			case "7":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     false,
					left:      true,
					top:       false,
					bottom:    true,
					x:         i,
					y:         y,
				})
			case "F":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     true,
					left:      false,
					top:       false,
					bottom:    true,
					x:         i,
					y:         y,
				})
			case ".", "S":
				rows[i] = append(rows[i], Pipe{
					character: pipe,
					right:     false,
					left:      false,
					top:       false,
					bottom:    false,
					x:         i,
					y:         y,
				})
			default:
				err := fmt.Sprintf("Unkown pipe error at position: [%d][%d]", i, y)
				panic(err)
			}
		}
	}

	return rows
}
