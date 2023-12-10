package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/fatih/color"
)

type Location struct {
	row int
	col int
}

type LoopNode struct {
	nodeType string
	location Location
	exits    []Location
}

func (node *LoopNode) ConnectsToNode(anotherNode *LoopNode) bool {
	isConnected := false
	for _, anotherExit := range anotherNode.exits {
		if node.location.row == anotherExit.row {
			isConnected = node.location.col == anotherExit.col
		}
		if node.location.col == anotherExit.col {
			isConnected = node.location.row == anotherExit.row
		}
		if isConnected {
			break
		}
	}
	return isConnected
}

func (node *LoopNode) AddExits() {
	if node.nodeType == "F" {
		node.exits = []Location{
			{node.location.row + 1, node.location.col},
			{node.location.row, node.location.col + 1},
		}
	}
	if node.nodeType == "L" {
		node.exits = []Location{
			{node.location.row - 1, node.location.col},
			{node.location.row, node.location.col + 1},
		}
	}
	if node.nodeType == "|" {
		node.exits = []Location{
			{node.location.row - 1, node.location.col},
			{node.location.row + 1, node.location.col},
		}
	}
	if node.nodeType == "-" {
		node.exits = []Location{
			{node.location.row, node.location.col - 1},
			{node.location.row, node.location.col + 1},
		}
	}
	if node.nodeType == "7" {
		node.exits = []Location{
			{node.location.row + 1, node.location.col},
			{node.location.row, node.location.col - 1},
		}
	}
	if node.nodeType == "J" {
		node.exits = []Location{
			{node.location.row - 1, node.location.col},
			{node.location.row, node.location.col - 1},
		}
	}
}

type Loop struct {
	closed bool
	items  []*LoopNode
	length int
}

func (loop *Loop) AddNode(node *LoopNode) bool {
	if loop.items[loop.length-1].ConnectsToNode(node) {
		loop.items = append(loop.items, node)
		loop.length++
		loop.closed = loop.length > 2 && loop.items[loop.length-1].ConnectsToNode(loop.items[0])
		return true
	}
	return false
}

func main() {
	file, err := os.ReadFile("./src/day10/p2/input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(file), "\n")

	sNode := LoopNode{"S", Location{}, []Location{}}
	matrix := [][]string{}
	for row := range inputStrings {
		elements := strings.Split(inputStrings[row], "")
		matrix = append(matrix, elements)
		for col := range elements {
			if elements[col] == "S" {
				sNode.location = Location{row, col}
			}
		}
	}

	matrixForPrinting := make([][]string, len(matrix))
	for i, row := range matrix {
		replaceJ := strings.ReplaceAll(strings.Join(row, ""), "J", "┘")
		replace7 := strings.ReplaceAll(replaceJ, "7", "┐")
		replaceF := strings.ReplaceAll(replace7, "F", "┌")
		replaceL := strings.ReplaceAll(replaceF, "L", "└")
		allReplaced := strings.Split(replaceL, "")
		matrixForPrinting[i] = append(matrixForPrinting[i], allReplaced...)
	}

	sLoop := Loop{}
	for srow := sNode.location.row - 1; srow <= sNode.location.row+1; srow++ {
		for scol := sNode.location.col - 1; scol <= sNode.location.col+1; scol++ {
			outOfBounds := srow < 0 || scol < 0 || srow >= len(matrix) || scol >= len(matrix[0])
			if outOfBounds {
				continue
			}
			isDot := matrix[srow][scol] == "."
			isDiagonal := math.Abs(float64(sNode.location.row-srow)) == math.Abs(float64(sNode.location.col-scol))
			isSame := srow == sNode.location.row && scol == sNode.location.col
			if isDot || isDiagonal || isSame || outOfBounds {
				continue
			}
			tempNodeCheck := LoopNode{matrix[srow][scol], Location{srow, scol}, []Location{}}
			tempNodeCheck.AddExits()
			for _, exit := range tempNodeCheck.exits {
				if exit.row == sNode.location.row && exit.col == sNode.location.col {
					sNode.exits = append(sNode.exits, Location{srow, scol})
				}
			}
		}
	}
	sLoop.items = []*LoopNode{&sNode}
	sLoop.length = 1
	matrix[sNode.location.row][sNode.location.col] = "."

	for !sLoop.closed {
		for _, exit := range sLoop.items[sLoop.length-1].exits {
			element := matrix[exit.row][exit.col]
			if element != "." {
				exitOneNode := LoopNode{element, Location{exit.row, exit.col}, []Location{}}
				exitOneNode.AddExits()
				if sLoop.AddNode(&exitOneNode) {
					// Resetting so we won't go there again
					matrix[exit.row][exit.col] = "."
					break
				}
			}
		}
	}

	area := getAreaViaShoelace(sLoop)
	// Via Pick's theorem
	answer := area - (float64(sLoop.length) / 2) + 1

	// printMaze(matrixForPrinting, sLoop)

	// Answer: 367
	fmt.Println("Answer:", answer)
}

func getAreaViaShoelace(sLoop Loop) float64 {
	doubleArea := 0
	amountOfVertices := len(sLoop.items)
	previousVertexId := amountOfVertices - 1
	for i := 0; i < amountOfVertices; i++ {
		prev := sLoop.items[previousVertexId]
		current := sLoop.items[i]
		doubleArea += prev.location.row*current.location.col - prev.location.col*current.location.row
		previousVertexId = i
	}
	return math.Abs(float64(doubleArea)) / 2
}

func printMaze(matrix [][]string, sLoop Loop) {
	colorElement := color.New(color.FgRed)
	sColor := color.New(color.FgHiYellow)
	greyedOut := color.New(color.FgBlack)
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == "S" {
				sColor.Print(matrix[row][col])
				continue
			}
			colored := false
			for _, node := range sLoop.items {
				colored = node.location.row == row && node.location.col == col
				if colored {
					break
				}
			}
			if colored {
				colorElement.Print(matrix[row][col])
				continue
			} else {
				greyedOut.Print(matrix[row][col])
				continue
			}
		}
		fmt.Println()
	}
}
