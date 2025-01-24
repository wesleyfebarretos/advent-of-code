package challenge

import (
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

const (
	GAP = " "
)

// Numeric Keypad
// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//     | 0 | A |
//     +---+---+

//  Directional Keypad
//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

type Position struct {
	Row int
	Col int
}

type PositionDir struct {
	Row int
	Col int
	Dir int
}

type Node struct {
	Row    int
	Col    int
	Weight int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Weight < pq[j].Weight }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func Pt1() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 1 result -> %v, runned in %s\n", result, time.Since(t))
		_ = clipboard.WriteAll(fmt.Sprintf("%v", result))
	}(time.Now())

	codes := parsePuzzle(utils.GetPuzzle())

	numericKeypad := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{GAP, "0", "A"},
	}

	directionalKeypad := [][]string{
		{GAP, "^", "A"},
		{"<", "v", ">"},
	}

	nPadStart := Position{Row: 3, Col: 2}
	dPadStart := Position{Row: 0, Col: 2}

	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	dPadMapByDir := map[int]string{
		0: "^",
		1: ">",
		2: "v",
		3: "<",
	}

	for _, code := range codes {
		initialPosition := nPadStart

		complexity1 := []string{}
		for _, button := range code {
			endPos := findEndPosition(string(button), numericKeypad)

			complexity1 = append(complexity1, dijkstra(initialPosition, endPos, numericKeypad, directions, dPadMapByDir))

			initialPosition = endPos
		}

		complexity1 = strings.Split(strings.Join(complexity1, "A")+"A", "")

		initialPosition = dPadStart

		complexity2 := []string{}

		for _, button := range complexity1 {
			endPos := findEndPosition(string(button), directionalKeypad)

			complexity2 = append(complexity2, dijkstra(initialPosition, endPos, directionalKeypad, directions, dPadMapByDir))

			initialPosition = endPos
		}

		complexity2 = strings.Split(strings.Join(complexity2, "A")+"A", "")

		initialPosition = dPadStart

		complexity3 := []string{}

		for _, button := range complexity2 {
			endPos := findEndPosition(string(button), directionalKeypad)

			complexity3 = append(complexity3, dijkstra(initialPosition, endPos, directionalKeypad, directions, dPadMapByDir))

			initialPosition = endPos
		}

		// reg := regexp.MustCompile(`[^0-9]+`)
		//
		// codeNums := reg.ReplaceAllString(code, "")
		//
		// cn, _ := strconv.Atoi(codeNums)

		complexity3 = strings.Split(strings.Join(complexity3, "A")+"A", "")
		// fmt.Println("----------")
		// fmt.Println(complexity1)
		// fmt.Println(complexity2)
		// fmt.Println(complexity3)
		//
		// fmt.Println(strings.Join(complexity3, ""))
		// fmt.Println(strings.Count(strings.Join(complexity3, ""), "A"))
		// fmt.Println(len(complexity3), cn, calcCodeComplexity(code, len(complexity3)), complexity3)
		result += calcCodeComplexity(code, len(complexity3))

	}
}

func calcCodeComplexity(code string, seqLen int) int {
	reg := regexp.MustCompile(`[^0-9]+`)

	codeNums := reg.ReplaceAllString(code, "")

	cn, _ := strconv.Atoi(codeNums)

	return cn * seqLen
}

func reverseString(s string) string {
	sli := []rune(s)

	for i, j := 0, len(sli)-1; i < j; i, j = i+1, j-1 {
		sli[i], sli[j] = sli[j], sli[i]
	}

	return string(sli)
}

func findEndPosition(value string, keypad [][]string) Position {
	for y := range keypad {
		for x := range keypad[y] {
			if keypad[y][x] == value {
				return Position{Row: y, Col: x}
			}
		}
	}

	panic("end position not found")
}

func dijkstra(start, end Position, keypad [][]string, directions [][2]int, dPadMapByDir map[int]string) string {
	pq := &PriorityQueue{}

	heap.Init(pq)

	// Need to Try all Possibles ways and get min length from them
	heap.Push(pq, Node{Row: start.Row, Col: start.Col, Weight: 0})

	dists := make(map[Position]int)

	visited := make(map[Position]bool)

	dists[start] = 0

	previousDir := make(map[Position]PositionDir)

	for pq.Len() > 0 {
		cp := heap.Pop(pq).(Node)

		if visited[Position{cp.Row, cp.Col}] {
			continue
		}

		visited[Position{cp.Row, cp.Col}] = true

		if cp.Row == end.Row && cp.Col == end.Col {
			break
		}

		for dirIdx, dir := range directions {

			dy, dx := dir[0], dir[1]

			npk := Position{Row: cp.Row + dy, Col: cp.Col + dx}
			np := Node{Row: cp.Row + dy, Col: cp.Col + dx, Weight: cp.Weight + 1}

			if outOfBounds(npk, keypad) || keypad[np.Row][np.Col] == GAP {
				continue
			}

			if _, ok := dists[npk]; !ok || np.Weight < dists[npk] {
				dists[npk] = np.Weight
				previousDir[npk] = PositionDir{Row: cp.Row, Col: cp.Col, Dir: dirIdx}
				heap.Push(pq, np)
			}

		}
	}

	s := strings.Builder{}

	prev, ok := previousDir[end]

	for ok {
		s.WriteString(dPadMapByDir[prev.Dir])
		prev, ok = previousDir[Position{Row: prev.Row, Col: prev.Col}]
	}

	return reverseString(s.String())
}

func outOfBounds(n Position, keypad [][]string) bool {
	return n.Row < 0 || n.Row >= len(keypad) || n.Col < 0 || n.Col >= len(keypad[0])
}

func parsePuzzle(puzzle string) []string {
	return strings.Split(puzzle, "\n")
}
