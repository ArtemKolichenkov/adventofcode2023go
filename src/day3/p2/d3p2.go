package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func charIsNumber(char *string) bool {
	_, err := strconv.Atoi(*char)
	return err == nil
}

type Gear struct {
	row         int
	col         int
	engineParts []*EnginePart
}

type NumberInMatrix interface {
	finalize() int
	rowBefore() int
	rowAfter() int
	colBefore() int
	colAfter() int
	getConnectedGear() *Gear
}

type EnginePart struct {
	row           int
	colStart      *int
	colEnd        *int
	value         string
	valueAsNumber int
}

func (enginePart *EnginePart) finalize(_colEnd *int, partsList *[]EnginePart, gearMap map[string]*Gear, matrix [][]string) {
	enginePart.colEnd = _colEnd
	err := error(nil)
	enginePart.valueAsNumber, err = strconv.Atoi(enginePart.value)
	if err != nil {
		panic(err)
	}
	*partsList = append(*partsList, *enginePart)
	connectedGear := enginePart.getConnectedGear(matrix)
	if connectedGear != nil {
		hash := strconv.Itoa(connectedGear.row) + strconv.Itoa(connectedGear.col)
		gearMap[hash] = connectedGear
	}
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

func (enginePart EnginePart) getConnectedGear(matrix [][]string) *Gear {
	for row := enginePart.rowBefore(); row <= enginePart.rowAfter(); row++ {
		for col := enginePart.colBefore(); col <= enginePart.colAfter(); col++ {
			if row == enginePart.row && col >= *enginePart.colStart && col <= *enginePart.colEnd {
				continue
			}
			if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
				continue
			}
			if matrix[row][col] == "*" {
				return &Gear{row, col, []*EnginePart{}}
			}
		}
	}
	return nil
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
	gears := map[string]*Gear{}
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
					tempEnginePart.finalize(&currentCol, &engineParts, gears, matrix)
				}
			} else {
				if tempEnginePart.colStart != nil {
					prevCol := col - 1
					tempEnginePart.finalize(&prevCol, &engineParts, gears, matrix)
				}
			}
		}
	}

	sum := 0
	for gearHash := range gears {
		for i := range engineParts {
			if gearIsInVicinityOfEnginePart(gears[gearHash], engineParts[i]) {
				gears[gearHash].engineParts = append(gears[gearHash].engineParts, &engineParts[i])
			}
		}
		if len(gears[gearHash].engineParts) == 2 {
			sum += gears[gearHash].engineParts[0].valueAsNumber * gears[gearHash].engineParts[1].valueAsNumber
		}
	}

	fmt.Println("Sum:", sum) // Answer: 69527306
}

func gearIsInVicinityOfEnginePart(gear *Gear, enginePart EnginePart) bool {
	return gear.row >= enginePart.rowBefore() && gear.row <= enginePart.rowAfter() && gear.col >= enginePart.colBefore() && gear.col <= enginePart.colAfter()
}
