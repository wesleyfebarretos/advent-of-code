package challenge

import (
	"container/heap"
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
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

	test := "029A"

	initialPosition := nPadStart

	numPath := []string{}
	for _, button := range test {
		endPos := findEndPosition(string(button), numericKeypad)

		numPath = append(numPath, dijkstra(initialPosition, endPos, numericKeypad, directions, dPadMapByDir))

		initialPosition = endPos
	}

	numPath = strings.Split(strings.Join(numPath, "A")+"A", "")

	initialPosition = dPadStart

	dirPath := []string{}
	fmt.Println("alo")
	for _, button := range numPath {
		endPos := findEndPosition(string(button), directionalKeypad)

		dirPath = append(dirPath, dijkstra(initialPosition, endPos, directionalKeypad, directions, dPadMapByDir))

		initialPosition = endPos
	}

	dirPath = strings.Split(strings.Join(dirPath, "A")+"A", "")

	initialPosition = dPadStart
	dirPath2 := []string{}

	fmt.Println("22")
	for _, button := range dirPath {
		endPos := findEndPosition(string(button), directionalKeypad)

		dirPath2 = append(dirPath2, dijkstra(initialPosition, endPos, directionalKeypad, directions, dPadMapByDir))

		initialPosition = endPos
	}

	dirPath2 = strings.Split(strings.Join(dirPath, "A")+"A", "")

	fmt.Println(numPath, dirPath, dirPath2, len(dirPath), len(dirPath2))
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

	if !ok {
		fmt.Println(end, start)
		panic("end not found")
	}

	// fmt.Println(end)
	for ok {
		s.WriteString(dPadMapByDir[prev.Dir])
		// fmt.Println(keypad[start.Row][start.Col], s.String(), keypad[end.Row][end.Col], prev)
		prev, ok = previousDir[Position{Row: prev.Row, Col: prev.Col}]
	}
	// fmt.Println(previousDir)
	// fmt.Println(dists)

	return s.String()
}

func outOfBounds(n Position, keypad [][]string) bool {
	return n.Row < 0 || n.Row >= len(keypad) || n.Col < 0 || n.Col >= len(keypad[0])
}
