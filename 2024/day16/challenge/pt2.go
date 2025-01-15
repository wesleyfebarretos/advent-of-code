package challenge

import (
	"container/heap"
	"fmt"
	"math"
	"time"

	"github.com/wesleyfebarretos/advent-of-code/2024/utils"
)

func Pt2() {
	result := 0

	defer func(t time.Time) {
		fmt.Printf("\nPart 2 result -> %d, runned in %s\n", result, time.Since(t))
	}(time.Now())

	reindeerMaze := parsePuzzle(utils.GetPuzzle())

	srow, scol, erow, ecol := findStartAndEnd(reindeerMaze)

	//  up - right - down - left
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	result = dijkstraPt2(srow, scol, erow, ecol, directions, reindeerMaze)
}

func dijkstraPt2(srow, scol, erow, ecol int, directions [][]int, reindeerMaze [][]string) int {
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

	distFromStart := make(map[string]int)

	distances[srow][scol] = 0

	prev := make(map[string][3]int)

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)

		if seen[makeUniqueMapKey(node.Row, node.Col, node.Dir)] {
			continue
		}

		seen[makeUniqueMapKey(node.Row, node.Col, node.Dir)] = true

		if _, ok := distFromStart[makeUniqueMapKey(node.Row, node.Col, node.Dir)]; !ok ||
			distFromStart[makeUniqueMapKey(node.Row, node.Col, node.Dir)] > node.Weight {
			distFromStart[makeUniqueMapKey(node.Row, node.Col, node.Dir)] = node.Weight
		}

		if node.Weight < distances[node.Row][node.Col] {
			distances[node.Row][node.Col] = node.Weight
			prev[makeUniqueMapKey(node.Row, node.Col)] = [3]int{node.Row - directions[node.Dir][0], node.Col - directions[node.Dir][1], node.Dir}
		}

		drow, dcol := directions[node.Dir][0], directions[node.Dir][1]

		nrow, ncol := node.Row+drow, node.Col+dcol

		if reindeerMaze[nrow][ncol] != WALL {
			heap.Push(pq, Node{node.Weight + TILE_COST, nrow, ncol, node.Dir})
		}

		heap.Push(pq, Node{node.Weight + ROTATE_COST, node.Row, node.Col, turnRight(node.Dir)})
		heap.Push(pq, Node{node.Weight + ROTATE_COST, node.Row, node.Col, turnLeft(node.Dir)})
	}

	seen = make(map[string]bool)
	distFromEnd := make(map[string]int)

	pq = &PriorityQueue{}

	heap.Init(pq)

	for i := range directions {
		heap.Push(pq, Node{0, erow, ecol, i})
	}

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)

		if seen[makeUniqueMapKey(node.Row, node.Col, node.Dir)] {
			continue
		}

		seen[makeUniqueMapKey(node.Row, node.Col, node.Dir)] = true

		if _, ok := distFromEnd[makeUniqueMapKey(node.Row, node.Col, turnBack(node.Dir))]; !ok ||
			distFromEnd[makeUniqueMapKey(node.Row, node.Col, turnBack(node.Dir))] > node.Weight {
			distFromEnd[makeUniqueMapKey(node.Row, node.Col, turnBack(node.Dir))] = node.Weight
		}

		drow, dcol := directions[node.Dir][0], directions[node.Dir][1]

		nrow, ncol := node.Row+drow, node.Col+dcol

		if reindeerMaze[nrow][ncol] != WALL {
			heap.Push(pq, Node{node.Weight + TILE_COST, nrow, ncol, node.Dir})
		}

		heap.Push(pq, Node{node.Weight + ROTATE_COST, node.Row, node.Col, turnRight(node.Dir)})
		heap.Push(pq, Node{node.Weight + ROTATE_COST, node.Row, node.Col, turnLeft(node.Dir)})
	}

	uniqueTilesSet := make(map[string]struct{})

	for y, row := range reindeerMaze {
		for x := range row {
			for dir := range directions {
				mapKey := makeUniqueMapKey(y, x, dir)

				//  i keep de keys of each position and direction but now with the weight starting from end so if start + end is equal is it part of the best paths
				//  i mean lets suppose that best path is 8000, so the end weight starting from beginnig is 8000 and starting from the end is 0, so start + end is = 8000
				//  and it keeps happen for each position
				if distFromStart[mapKey]+distFromEnd[mapKey] == distances[erow][ecol] {
					uniqueTilesSet[makeUniqueMapKey(y, x)] = struct{}{}
				}
			}
		}
	}

	return len(uniqueTilesSet)
}

func turnBack(dir int) int {
	return (dir + 2) % 4
}
