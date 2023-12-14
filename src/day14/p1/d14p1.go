package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./src/day14/p1/input.txt")
	if err != nil {
		panic(err)
	}
	strings := strings.Split(string(file), "\n")

	rowCount := len(strings)
	colCount := len(strings[0])

	matrix := make([][]rune, rowCount)
	for row := 0; row < rowCount; row++ {
		matrix[row] = make([]rune, colCount)
		for col := 0; col < colCount; col++ {
			matrix[row][col] = rune(strings[row][col])
		}
	}

	for col := 0; col < colCount; col++ {
		for r := 0; r < rowCount; r++ {
			for row := 0; row < rowCount; row++ {
				if matrix[row][col] == 'O' && row > 0 && matrix[row-1][col] == '.' {
					matrix[row][col] = '.'
					matrix[row-1][col] = 'O'
				}
			}
		}
	}

	answer := 0

	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			if matrix[row][col] == 'O' {
				answer += rowCount - row
			}
		}
	}

	fmt.Println("Answer:", answer) // 108840
}
