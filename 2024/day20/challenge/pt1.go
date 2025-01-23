package challenge

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

type Position struct {
	Row int
	Col int
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

const (
	WALL = "#"
)

func Pt1() {
	result := 0

	defer func(t time.Time) {
		_ = clipboard.WriteAll(fmt.Sprintf("%v", result))
		fmt.Printf("\nPart 1 result -> %v, runned in %s\n", result, time.Since(t))
	}(time.Now())

	raceMap := parsePuzzle(utils.GetPuzzle())

	start, end := findStartAndEnd(raceMap)

	distsFromStart := dijkstra(start, raceMap)
	distsFromEnd := dijkstra(end, raceMap)

	for sn, sv := range distsFromStart {
		for en, ev := range distsFromEnd {
			manhattanDist := manhattanDistance(sn, en)
			if manhattanDist <= 2 && sv+ev+manhattanDist <= distsFromStart[end]-100 {
				result++
			}
		}
	}
}

func manhattanDistance(n1, n2 Position) int {
	return int(math.Abs(float64(n1.Col-n2.Col)) + math.Abs(float64(n1.Row-n2.Row)))
}

func dijkstra(start Position, raceMap [][]string) map[Position]int {
	pq := &PriorityQueue{}

	heap.Init(pq)

	heap.Push(pq, Node{Row: start.Row, Col: start.Col, Weight: 0})

	dists := make(map[Position]int)

	dists[start] = 0

	for pq.Len() > 0 {
		cp := pq.Pop().(Node)

		for _, adj := range adjacents(cp, raceMap) {

			k := Position{Row: adj.Row, Col: adj.Col}

			if _, ok := dists[k]; !ok || cp.Weight+1 < dists[k] {
				dists[k] = cp.Weight + 1
				pq.Push(adj)
			}

		}
	}

	return dists
}

func adjacents(n Node, raceMap [][]string) []Node {
	adjs := []Node{}

	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for _, dir := range directions {

		dr, dc := dir[0], dir[1]

		nr, nc := n.Row+dr, n.Col+dc

		if nr < 0 || nr >= len(raceMap) || nc < 0 || nc >= len(raceMap[0]) || raceMap[nr][nc] == WALL {
			continue
		}

		adjs = append(adjs, Node{Row: nr, Col: nc, Weight: n.Weight + 1})
	}

	return adjs
}

func findStartAndEnd(raceMap [][]string) (Position, Position) {
	var start, end Position

	for y, row := range raceMap {
		for x, col := range row {
			switch col {
			case "S":
				start = Position{Row: y, Col: x}
			case "E":
				end = Position{Row: y, Col: x}
			}
		}
	}

	return start, end
}

func parsePuzzle(puzzle string) [][]string {
	puzzleSlice := strings.Split(puzzle, "\n")

	matrix := make([][]string, len(puzzleSlice))

	for y, row := range puzzleSlice {
		matrix[y] = make([]string, len(row))
		for x, col := range row {
			matrix[y][x] = string(col)
		}

	}

	return matrix
}
