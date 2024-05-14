package challenge1

import (
	"fmt"
	"math"
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

	for _direction, coordinates := range directions {
		direction := _direction
		for _coordinate, _coordinateValue := range coordinates {
			coordinate := _coordinate
			coordinateValue := _coordinateValue
			pipe := startPipe
			for {
				matchDirection, nextDirection, nextPipe := walk(rows, direction, coordinate, coordinateValue, pipe)

				if !matchDirection {
					index++
					break
				}

				for key, value := range directions[nextDirection] {
					coordinate = key
					coordinateValue = value
				}

				pipe = nextPipe
				direction = nextDirection
				steps[index]++
			}
		}
	}

	maxSteps := slices.Max(steps)
	fmt.Printf("Steps: %v", math.Ceil(float64(maxSteps)/2))
}

func walk(rows [][]Pipe, direction, coordinate string, coordinateValue int, pipe Pipe) (bool, string, Pipe) {
	x, y := pipe.x, pipe.y

	if coordinate == "x" {
		x += coordinateValue
	} else {
		y += coordinateValue
	}

	if x < 0 || y < 0 {
		return false, direction, pipe
	}
	nextPipe := rows[y][x]
	// fmt.Printf("Direction: %s\nPipe: %+v\nNextPipe: %+v\nX: %d\nY: %d\n", direction, pipe, nextPipe, x, y)

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
