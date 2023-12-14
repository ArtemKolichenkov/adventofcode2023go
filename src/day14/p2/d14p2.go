package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./src/day14/p2/input.txt")
	if err != nil {
		panic(err)
	}
	strings := strings.Split(string(file), "\n")

	matrix := make([][]rune, len(strings))
	for row := 0; row < len(strings); row++ {
		matrix[row] = make([]rune, len(strings[0]))
		for col := 0; col < len(strings[0]); col++ {
			matrix[row][col] = rune(strings[row][col])
		}
	}

	loadMap := make(map[string]int)
	skip := false
	totalIterations := 1000000000
	iteration := 0
	for iteration < totalIterations {
		matrix = cycle(matrix)
		iteration++
		hash := hashMatrix(matrix)
		if !skip {
			if observedIteration, ok := loadMap[hash]; ok {
				period := iteration - observedIteration
				iterationsLeft := totalIterations - iteration
				periodsLeft := iterationsLeft / period
				iteration += periodsLeft * period
				skip = true
			}
			loadMap[hash] = iteration
		}
	}

	fmt.Println("Answer:", getRocksLoad(matrix)) // 103445
}

func hashMatrix(matrix [][]rune) (hash string) {
	for row := 0; row < len(matrix); row++ {
		hash += string(matrix[row])
	}
	return hash
}

func getRocksLoad(matrix [][]rune) (load int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 'O' {
				load += len(matrix) - row
			}
		}
	}
	return load
}

func cycle(matrix [][]rune) [][]rune {
	for i := 0; i < 4; i++ {
		matrix = tiltNorth(matrix)
		matrix = rotateMatrix(matrix)
	}
	return matrix
}

func tiltNorth(matrix [][]rune) [][]rune {
	for col := 0; col < len(matrix[0]); col++ {
		for r := 0; r < len(matrix); r++ {
			for row := 0; row < len(matrix); row++ {
				if matrix[row][col] == 'O' && row > 0 && matrix[row-1][col] == '.' {
					matrix[row][col] = '.'
					matrix[row-1][col] = 'O'
				}
			}
		}
	}
	return matrix
}

// (matrix) north -> (matrix) west -> (matrix) south -> (matrix) east
func rotateMatrix(matrix [][]rune) [][]rune {
	rowCount := len(matrix)
	colCount := len(matrix[0])
	newMatrix := make([][]rune, colCount)
	for row := 0; row < rowCount; row++ {
		newMatrix[row] = make([]rune, rowCount)
		for col := 0; col < colCount; col++ {
			newMatrix[row][col] = '?'
		}
	}
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			newMatrix[col][rowCount-1-row] = matrix[row][col]
		}
	}
	return newMatrix
}
