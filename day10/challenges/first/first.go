package challenge1

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
)

type Pipe struct {
	Character string
	Right     bool
	Left      bool
	Top       bool
	Down      bool
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
			if pipe.Character == "S" {
				startPipe = pipe
			}
		}
	}

	directions := map[string]map[string]int{
		"Right": {
			"x": 1,
		},
		"Left": {
			"x": -1,
		},
		"Top": {
			"y": -1,
		},
		"Down": {
			"y": 1,
		},
	}

	steps := make([]int, 4)
	index := 0

	// fmt.Printf("Initial Coordinates:\nX: %d\nY: %d\n", startPipe.x, startPipe.y)
	for direction, coordinates := range directions {
		startDirection := direction
		for coordinate, value := range coordinates {
			x, y := startPipe.x, startPipe.y
			if coordinate == "x" {
				x += value
			} else {
				y += value
			}

			for {
				fmt.Printf("PrevPipe: %+v\n", rows[y][x])
				matchDirection, nextDirection, pipe := walk(rows, startDirection, x, y)

				if !matchDirection {
					index++
					break
				}

				x, y = pipe.x, pipe.y

				for coordinate, value := range directions[nextDirection] {
					if coordinate == "x" {
						x += value
					} else {
						y += value
					}
				}

				startDirection = nextDirection
				steps[index]++
			}
		}
	}

	fmt.Printf("Steps: %v", steps)
}

func walk(rows [][]Pipe, direction string, x, y int) (bool, string, Pipe) {
	// TODO: Verificar o motivo de não estar continuando o seguimento dos passos
	// A função está terminando muito cedo
	// O nextPipe está com o mesmo valor do PipeCurrent da sessão, entender bem e ajustar fluxo de dados
	fmt.Printf("Direction: %s\nPipe: %+v\nX: %d\nY: %d\n", direction, rows[y][x], x, y)
	nextPipe := rows[y][x]

	nextDirection := direction

	directions := []string{
		"Right",
		"Left",
		"Top",
		"Down",
	}

	reverseDirection := getReverseDirection(direction)

	matchDirection := reflect.ValueOf(nextPipe).FieldByName(reverseDirection)

	if !matchDirection.Bool() {
		return false, nextDirection, nextPipe
	}

	v := reflect.ValueOf(nextPipe)
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := typ.Field(i).Name
		if slices.Index(directions, key) == -1 || key == reverseDirection {
			continue
		}
		value := v.Field(i)
		if value.Bool() {
			nextDirection = key
			break
		}
	}

	return matchDirection.Bool(), nextDirection, nextPipe
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
					Character: pipe,
					Right:     false,
					Left:      false,
					Top:       true,
					Down:      true,
					x:         i,
					y:         y,
				})
			case "-":
				rows[i] = append(rows[i], Pipe{
					Character: pipe,
					Right:     true,
					Left:      true,
					Top:       false,
					Down:      false,
					x:         i,
					y:         y,
				})
			case "J":
				rows[i] = append(rows[i], Pipe{
					Character: pipe,
					Right:     false,
					Left:      true,
					Top:       true,
					Down:      false,
					x:         i,
					y:         y,
				})
			case "L":
				rows[i] = append(rows[i], Pipe{
					Character: pipe,
					Right:     true,
					Left:      false,
					Top:       true,
					Down:      false,
					x:         i,
					y:         y,
				})
			case "7":
				rows[i] = append(rows[i], Pipe{
					Character: pipe,
					Right:     false,
					Left:      true,
					Top:       false,
					Down:      true,
					x:         i,
					y:         y,
				})
			case "F":
				rows[i] = append(rows[i], Pipe{
					Character: pipe,
					Right:     true,
					Left:      false,
					Top:       false,
					Down:      true,
					x:         i,
					y:         y,
				})
			case ".", "S":
				rows[i] = append(rows[i], Pipe{
					Character: pipe,
					Right:     false,
					Left:      false,
					Top:       false,
					Down:      false,
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

func getReverseDirection(direction string) string {
	switch direction {
	case "Right":
		return "Left"
	case "Left":
		return "Right"
	case "Down":
		return "Top"
	case "Top":
		return "Down"
	default:
		panic("invalid direction")
	}
}
