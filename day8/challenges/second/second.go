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

const (
	startNode  Node = Node("AAA")
	finishNode Node = Node("ZZZ")
)

func Run() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	mapp := strings.Split(strings.TrimSpace(string(input)), "\n")

	instructions := mapp[0]

	nodes := mapp[2:]

	reg := regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

	nodesMap := make(map[Node]NodeDirections)

	for _, node := range nodes {
		nodeInfo := reg.FindStringSubmatch(node)[1:]

		node, left, right := nodeInfo[0], nodeInfo[1], nodeInfo[2]

		nodesMap[Node(node)] = NodeDirections{Node(left), Node(right)}
	}

	stepsToFinish := execInstructions(instructions, nodesMap, startNode, 0)

	fmt.Println(stepsToFinish)
}

func walk(nodeDirections NodeDirections, instruction rune) Node {
	switch string(instruction) {
	case "L":
		return Node(nodeDirections.left)
	default:
		return Node(nodeDirections.right)
	}
}

func execInstructions(instructions string, nodesMap map[Node]NodeDirections, currentNode Node, stepsToFinish int) int {
	finish := false
	for _, instruction := range instructions {
		nodeDirections := nodesMap[currentNode]
		currentNode = walk(nodeDirections, instruction)

		stepsToFinish++

		if currentNode == finishNode {
			finish = true
			break
		}
	}

	if !finish {
		return execInstructions(instructions, nodesMap, currentNode, stepsToFinish)
	}

	return stepsToFinish
}
