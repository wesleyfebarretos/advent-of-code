package challenge2

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node string

type NodeDirections struct {
	left  Node
	right Node
}

func Run() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	mapp := strings.Split(strings.TrimSpace(string(input)), "\n")

	instructions := mapp[0]

	nodes := mapp[2:]

	reg := regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)

	nodesMap := make(map[Node]NodeDirections)

	var startNodes []Node

	for _, node := range nodes {
		nodeInfo := reg.FindStringSubmatch(node)[1:]

		node, left, right := nodeInfo[0], nodeInfo[1], nodeInfo[2]

		nodesMap[Node(node)] = NodeDirections{Node(left), Node(right)}

		if Node(node[2]) == "A" {
			startNodes = append(startNodes, Node(node))
		}
	}

	steps := make([]int, len(startNodes))

	execInstructions(instructions, nodesMap, startNodes, steps)

	totalSteps := 1

	for _, step := range steps {
		totalSteps = calcLCM(totalSteps, step)
	}

	fmt.Println(totalSteps)
}

func ghostWalk(startNodes []Node, nodesMap map[Node]NodeDirections, instruction rune, steps []int) int {
	endPaths := 0

	for i, node := range startNodes {
		if Node(startNodes[i][2]) == "Z" {
			endPaths++
			continue
		}

		steps[i] += 1

		nodeDirections := nodesMap[node]
		switch string(instruction) {
		case "L":
			startNodes[i] = nodeDirections.left
		default:
			startNodes[i] = nodeDirections.right
		}
	}

	return endPaths
}

func execInstructions(instructions string, nodesMap map[Node]NodeDirections, startNodes []Node, steps []int) {
	finish := false

	for _, instruction := range instructions {
		endPaths := ghostWalk(startNodes, nodesMap, instruction, steps)

		if endPaths == len(startNodes) {
			finish = true
			break
		}
	}

	if !finish {
		execInstructions(instructions, nodesMap, startNodes, steps)
	}
}

// Function to calculate the greatest common divisor (GCD) Euclid's Algorithm
func calcGCD(a, b int) int {
	if b == 0 {
		return a
	}

	return calcGCD(b, a%b)
}

// Function to calculate LCM
func calcLCM(a, b int) int {
	return (a * b) / calcGCD(a, b)
}
