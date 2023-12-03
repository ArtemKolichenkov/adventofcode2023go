package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isDigitOrDot(char *string) bool {
	return slices.Contains([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "."}, *char)
}

func charIsNumber(char *string) bool {
	_, err := strconv.Atoi(*char)
	return err == nil
}

type NumberInMatrix interface {
	finalize() int
	rowBefore() int
	rowAfter() int
	colBefore() int
	colAfter() int
	hasAdjecentSymbols() bool
}

type EnginePart struct {
	row           int
	colStart      *int
	colEnd        *int
	value         string
	valueAsNumber int
}

func (enginePart *EnginePart) finalize(_colEnd *int, partsList *[]EnginePart) {
	enginePart.colEnd = _colEnd
	err := error(nil)
	enginePart.valueAsNumber, err = strconv.Atoi(enginePart.value)
	if err != nil {
		panic(err)
	}
	*partsList = append(*partsList, *enginePart)
	enginePart.row, enginePart.valueAsNumber = 0, 0
	enginePart.colStart, enginePart.colEnd = nil, nil
	enginePart.value = ""
}

func (enginePart EnginePart) rowBefore() int {
	return enginePart.row - 1
}

func (enginePart EnginePart) rowAfter() int {
	return enginePart.row + 1
}

func (enginePart EnginePart) colBefore() int {
	return *enginePart.colStart - 1
}

func (enginePart EnginePart) colAfter() int {
	return *enginePart.colEnd + 1
}

func (enginePart EnginePart) hasAdjecentSymbols(matrix [][]string) bool {
	for row := enginePart.rowBefore(); row <= enginePart.rowAfter(); row++ {
		for col := enginePart.colBefore(); col <= enginePart.colAfter(); col++ {
			if row == enginePart.row && col >= *enginePart.colStart && col <= *enginePart.colEnd {
				continue
			}
			if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
				continue
			}
			if !isDigitOrDot(&matrix[row][col]) {
				return true
			}
		}
	}
	return false
}

func main() {
	file, err := os.ReadFile("./src/day3/p2/input.txt")
	if err != nil {
		panic(err)
	}

	matrix := [][]string{}
	for _, line := range strings.Split(string(file), "\n") {
		matrix = append(matrix, strings.Split(line, ""))
	}

	tempEnginePart := EnginePart{}
	engineParts := []EnginePart{}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			char := matrix[row][col]

			if charIsNumber(&char) {
				currentCol := col
				if tempEnginePart.colStart == nil {
					tempEnginePart.colStart = &currentCol
					tempEnginePart.row = row
					tempEnginePart.value = char
				} else {
					tempEnginePart.value = tempEnginePart.value + char
				}
				if col == len(matrix[row])-1 {
					tempEnginePart.finalize(&currentCol, &engineParts)
				}
			} else {
				if tempEnginePart.colStart != nil {
					prevCol := col - 1
					tempEnginePart.finalize(&prevCol, &engineParts)
				}
			}
		}
	}

	sum := 0
	for _, num := range engineParts {
		if num.hasAdjecentSymbols(matrix) {
			sum += num.valueAsNumber
		}
	}

	fmt.Println("Sum:", sum) // Answer: 521515
}
