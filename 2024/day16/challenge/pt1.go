package challenge

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
	START       = "S"
	END         = "E"
	WALL        = "#"
	TILE_COST   = 1
	ROTATE_COST = 1000
)

type Node struct {
	Weight int
	Row    int
	Col    int
	Dir    int
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
		fmt.Printf("\nPart 1 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	reindeerMaze := parsePuzzle(utils.GetPuzzle())

	srow, scol, erow, ecol := findStartAndEnd(reindeerMaze)

	//  up - right - down - left
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	result = dijkstra(srow, scol, erow, ecol, directions, reindeerMaze)
}

func dijkstra(srow, scol, erow, ecol int, directions [][]int, reindeerMaze [][]string) int {
	pq := &PriorityQueue{}

	heap.Init(pq)

	heap.Push(pq, Node{0, srow, scol, RIGHT})

	seen := make(map[string]bool)

	distances := make([][]int, len(reindeerMaze))

	for y := range reindeerMaze {
		distances[y] = make([]int, len(reindeerMaze[y]))
		for x := range reindeerMaze[y] {
			distances[y][x] = math.MaxInt
		}
	}

	distances[srow][scol] = 0

	prev := make(map[string][3]int)

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)

		if seen[makeUniqueMapKey(node.Row, node.Col, node.Dir)] {
			continue
		}

		seen[makeUniqueMapKey(node.Row, node.Col, node.Dir)] = true

		if node.Weight < distances[node.Row][node.Col] {
			distances[node.Row][node.Col] = node.Weight
			prev[makeUniqueMapKey(node.Row, node.Col)] = [3]int{node.Row - directions[node.Dir][0], node.Col - directions[node.Dir][1], node.Dir}
		}

		if reindeerMaze[node.Row][node.Col] == END {
			break
		}

		drow, dcol := directions[node.Dir][0], directions[node.Dir][1]

		nrow, ncol := node.Row+drow, node.Col+dcol

		if reindeerMaze[nrow][ncol] != WALL {
			heap.Push(pq, Node{node.Weight + TILE_COST, nrow, ncol, node.Dir})
		}

		heap.Push(pq, Node{node.Weight + ROTATE_COST, node.Row, node.Col, turnRight(node.Dir)})
		heap.Push(pq, Node{node.Weight + ROTATE_COST, node.Row, node.Col, turnLeft(node.Dir)})
	}

	return distances[erow][ecol]
}

func turnLeft(dir int) int {
	return (dir + 3) % 4
}

func turnRight(dir int) int {
	return (dir + 1) % 4
}

func findStartAndEnd(reindeerMaze [][]string) (int, int, int, int) {
	srow, scol, erow, ecol := -1, -1, -1, -1

	for y, row := range reindeerMaze {
		for x, col := range row {
			switch col {
			case START:
				srow = y
				scol = x
			case END:
				erow = y
				ecol = x
			}
		}
	}

	if srow == -1 || scol == -1 || erow == -1 || ecol == -1 {
		panic("start or end not found")
	}

	return srow, scol, erow, ecol
}

func parsePuzzle(puzzle string) [][]string {
	puzzleSlice := strings.Split(puzzle, "\n")

	matrix := make([][]string, len(puzzleSlice))

	for i, row := range puzzleSlice {
		matrix[i] = make([]string, len(row))

		for y, col := range row {
			matrix[i][y] = string(col)
		}
	}

	return matrix
}

func printPath(reindeerMaze [][]string, erow, ecol int, prev map[string][3]int) {
	curr, ok := prev[makeUniqueMapKey(erow, ecol)]

	for ok && reindeerMaze[curr[0]][curr[1]] != START {
		reindeerMaze[curr[0]][curr[1]] = getDirectionCharByDir(curr[2])

		curr, ok = prev[makeUniqueMapKey(curr[0], curr[1])]
	}

	utils.PrintMatrix(reindeerMaze)
}

func getDirectionCharByDir(dir int) string {
	switch dir {
	case 0:
		return "^"
	case 1:
		return ">"
	case 2:
		return "v"
	default:
		return "<"
	}
}

func makeUniqueMapKey(p ...int) string {
	strValues := make([]string, len(p))
	for i, v := range p {
		strValues[i] = fmt.Sprintf("%v", v)
	}

	return strings.Join(strValues, "|")
}
