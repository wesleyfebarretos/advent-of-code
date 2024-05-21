package challenge2

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
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

	pipes := getPipes(input)

	var startPipe Pipe

	for _, row := range pipes {
		for _, pipe := range row {
			if pipe.Character == "S" {
				startPipe = pipe
			}
		}
	}

	var maybeS []string

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

	for direction, dValue := range directions {
		for coordinate, cValue := range dValue {
			x, y := startPipe.x, startPipe.y

			if coordinate == "x" {
				x += cValue
			} else {
				y += cValue
			}

			reverseDirection := getReverseDirection(direction)
			nextPipe := pipes[y][x]
			matchDirection := reflect.ValueOf(nextPipe).FieldByName(reverseDirection)

			if matchDirection.Bool() {
				maybeS = append(maybeS, direction)
			}
		}
	}

	var truthyS string

	switch strings.Join(maybeS, "") {
	case "TopDown", "DownTop":
		truthyS = "|"
	case "LeftRight", "RightLeft":
		truthyS = "-"
	case "DownRight", "RightDown":
		truthyS = "F"
	case "DownLeft", "LeftDown":
		truthyS = "7"
	case "TopLeft", "LeftTop":
		truthyS = "J"
	case "TopRight", "RightTop":
		truthyS = "L"
	}

	reg := regexp.MustCompile(`[S]+`)
	rows := strings.Split(strings.TrimSpace(reg.ReplaceAllString(string(input), truthyS)), "\n")

	// TODO: Retirar caracteres que não fazem parte do loop principal e após isso aplicar ray casting algorithm
	fmt.Println(strings.Join(rows, "\n"))
}

func getPipes(input []byte) [][]Pipe {
	ground := strings.Split(strings.TrimSpace(string(input)), "\n")

	rows := make([][]Pipe, len(ground))

	for y, row := range ground {
		for x, pipeRune := range row {
			pipe := string(pipeRune)
			switch pipe {
			case "|":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     false,
					Left:      false,
					Top:       true,
					Down:      true,
					x:         x,
					y:         y,
				})
			case "-":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     true,
					Left:      true,
					Top:       false,
					Down:      false,
					x:         x,
					y:         y,
				})
			case "J":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     false,
					Left:      true,
					Top:       true,
					Down:      false,
					x:         x,
					y:         y,
				})
			case "L":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     true,
					Left:      false,
					Top:       true,
					Down:      false,
					x:         x,
					y:         y,
				})
			case "7":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     false,
					Left:      true,
					Top:       false,
					Down:      true,
					x:         x,
					y:         y,
				})
			case "F":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     true,
					Left:      false,
					Top:       false,
					Down:      true,
					x:         x,
					y:         y,
				})
			case ".", "S":
				rows[y] = append(rows[y], Pipe{
					Character: pipe,
					Right:     false,
					Left:      false,
					Top:       false,
					Down:      false,
					x:         x,
					y:         y,
				})
			default:
				err := fmt.Sprintf("Unkown pipe error at position: [%d][%d]", y, x)
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
