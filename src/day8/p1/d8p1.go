package main

import (
	"fmt"
	"os"
	"strings"
)

// Answer: 14429
func main() {
	file, err := os.ReadFile("./src/day8/p1/input.txt")
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

	preprocessedNodes := make(map[string][]string, len(nodes))

	for i := range nodes {
		nodeLine := nodes[i]
		nodeName := strings.Split(nodeLine, " ")[0]
		nodeTupleString := strings.Split(nodeLine, " = ")[1]
		nodeTuple := strings.Split(nodeTupleString[1:len(nodeTupleString)-1], ", ")
		preprocessedNodes[nodeName] = nodeTuple
	}

	steps := 0
	currentNodeName := "AAA"
	currentNode := preprocessedNodes[currentNodeName]

	for currentNodeName != "ZZZ" {
		currentNodeName = currentNode[commands[steps%commandsLength]]
		currentNode = preprocessedNodes[currentNodeName]
		steps++
	}

	fmt.Println("Answer:", steps)
}
