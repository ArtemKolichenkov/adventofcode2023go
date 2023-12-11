package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type HashSymbol struct {
	row int
	col int
}

func main() {
	file, err := os.ReadFile("./src/day11/p1/input.txt")
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
	printMatrix(matrix, "Original matrix")

	// Matrix expansion
	toExpandRow, toExpandCol := findRowsAndColumnsToExpand(matrix)
	expandedColCount := getExtraItems(toExpandCol, len(matrix[0]))
	expandedMatrix := expandMatrix(matrix, toExpandRow, toExpandCol, expandedColCount)
	printMatrix(expandedMatrix, "Expanded matrix")

	// Getting hash symbol locations
	hashSymbols := getHashSymbols(expandedMatrix)
	answer := getHashDistances(hashSymbols)
	fmt.Println("Answer:", answer) // 9974721
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

func getHashSymbols(matrix [][]string) []HashSymbol {
	hashSymbols := []HashSymbol{}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == "#" {
				hashSymbol := HashSymbol{row, col}
				hashSymbols = append(hashSymbols, hashSymbol)
			}
		}
	}
	return hashSymbols
}

func findRowsAndColumnsToExpand(matrix [][]string) ([]bool, []bool) {
	toExpandRow := make([]bool, len(matrix))
	toExpandCol := make([]bool, len(matrix[0]))
	for i := range toExpandCol {
		toExpandCol[i] = true
	}
	for row := range matrix {
		toExpandRow[row] = allDots(matrix[row])
		for col := range matrix[row] {
			if matrix[row][col] == "#" {
				toExpandCol[col] = false
			}
		}
	}
	return toExpandRow, toExpandCol
}

func expandMatrix(matrix [][]string, toExpandRow []bool, toExpandCol []bool, expandedColCount int) [][]string {
	expandedMatrix := [][]string{}
	for row := range matrix {
		if toExpandRow[row] {
			emptyRow := getEmptyRow(expandedColCount)
			expandedMatrix = append(expandedMatrix, emptyRow, emptyRow)
			continue
		}
		expandedMatrix = append(expandedMatrix, matrix[row])
		for col := range matrix[row] {
			if toExpandCol[col] {
				offset := len(expandedMatrix[len(expandedMatrix)-1]) - len(matrix[row])
				expandedMatrix[len(expandedMatrix)-1] = slices.Insert(expandedMatrix[len(expandedMatrix)-1], col+offset, "Y")
			}
		}
	}
	return expandedMatrix
}

func printMatrix(matrix [][]string, title string) {
	fmt.Println(title, ":", len(matrix), "rows", len(matrix[0]), "columns")
	for row := range matrix {
		fmt.Println(matrix[row])
	}
}

func getExtraItems(binaryMap []bool, originalLength int) int {
	extraItems := 0
	for _, item := range binaryMap {
		if item {
			extraItems++
		}
	}
	return originalLength + extraItems
}

func getEmptyRow(size int) []string {
	emptyRow := make([]string, size)
	for i := range emptyRow {
		emptyRow[i] = "X"
	}
	return emptyRow
}

func allDots(slce []string) bool {
	for _, item := range slce {
		if item != "." {
			return false
		}
	}
	return true
}
