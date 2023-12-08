package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name          string
	tuple         []string
	reachedZsteps int
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

// Answer:
func main() {
	file, err := os.ReadFile("./src/day8/p2/input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(file), "\n")

	commandStrings := strings.Split(inputStrings[0], "")
	commands := make([]int, len(commandStrings))
	for i := range commandStrings {
		if commandStrings[i] == "R" {
			commands[i] = 1
		} else {
			commands[i] = 0
		}
	}
	commandsLength := len(commands)

	nodes := inputStrings[2:]

	preprocessedNodes := make(map[string]Node, len(nodes))

	for i := range nodes {
		nodeLine := nodes[i]
		nodeName := strings.Split(nodeLine, " ")[0]
		nodeTupleString := strings.Split(nodeLine, " = ")[1]
		nodeTuple := strings.Split(nodeTupleString[1:len(nodeTupleString)-1], ", ")
		preprocessedNodes[nodeName] = Node{nodeName, nodeTuple, 0}
	}

	steps := 0
	activeNodes := []Node{}
	for nodeName := range preprocessedNodes {
		if nodeName[2] == 'A' {
			activeNodes = append(activeNodes, preprocessedNodes[nodeName])
		}
	}
	nodesEndingWithZ := 0

	for nodesEndingWithZ != len(activeNodes) {
		nodesEndingWithZ = 0
		for i := range activeNodes {
			if activeNodes[i].reachedZsteps != 0 {
				nodesEndingWithZ++
				continue
			}
			nextNodeName := activeNodes[i].tuple[commands[steps%commandsLength]]
			nextNode := preprocessedNodes[nextNodeName]
			activeNodes[i].name = nextNodeName
			activeNodes[i].tuple = nextNode.tuple
			if nextNodeName[2] == 'Z' {
				activeNodes[i].reachedZsteps = steps + 1
				nodesEndingWithZ++
			}
		}
		steps++
	}

	nodeSteps := make([]int, len(activeNodes))
	for i := range activeNodes {
		nodeSteps[i] = activeNodes[i].reachedZsteps
	}
	answer := LCM(nodeSteps[0], nodeSteps[1], nodeSteps...)

	fmt.Println("Answer:", answer)
}
