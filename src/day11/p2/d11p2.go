package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type HashSymbol struct {
	row int
	col int
}

func main() {
	file, err := os.ReadFile("./src/day11/p2/input.txt")
	if err != nil {
		panic(err)
	}

	// Text to matrix
	inputStrings := strings.Split(string(file), "\n")
	matrix := [][]string{}
	for _, line := range inputStrings {
		textRow := strings.Split(line, "")
		matrix = append(matrix, textRow)
	}

	hashSymbols := getFastHashes(matrix, 1000000-1)
	answer := getHashDistances(hashSymbols)
	fmt.Println("Answer:", answer) // 702770569197
}

func getHashDistances(hashSymbols []HashSymbol) (totalDistance int) {
	for i, hashSymbol := range hashSymbols {
		for _, otherHashSymbol := range hashSymbols[i+1:] {
			distance := int(math.Abs(float64(otherHashSymbol.col-hashSymbol.col)) + math.Abs(float64(otherHashSymbol.row-hashSymbol.row)))
			totalDistance += distance
		}
	}
	return totalDistance
}

func getFastHashes(matrix [][]string, expansionFactor int) []HashSymbol {
	hashSymbols := []HashSymbol{}
	columnOffsets := []int{0}

	for col := range matrix[0] {
		isColumnEmpty := true
		for row := range matrix {
			if matrix[row][col] == "#" {
				isColumnEmpty = false
				break
			}
		}
		lastOffset := columnOffsets[len(columnOffsets)-1]
		if isColumnEmpty {
			columnOffsets = append(columnOffsets, lastOffset+expansionFactor)
		} else {
			columnOffsets = append(columnOffsets, lastOffset)
		}
	}

	columnOffsets = columnOffsets[1:]

	rowOffset := 0
	for row := range matrix {
		isEmptyRow := true
		for col := range matrix[row] {
			if matrix[row][col] == "#" {
				hashSymbols = append(hashSymbols, HashSymbol{row + rowOffset, col + columnOffsets[col]})
				isEmptyRow = false
			}
		}
		if isEmptyRow {
			rowOffset += expansionFactor
		}
	}
	return hashSymbols
}
