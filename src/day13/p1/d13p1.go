package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./src/day13/p1/input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(file), "\n\n")

	answer := 0
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		rowCount := len(lines)
		colCount := len(lines[0])
		fmt.Println(rowCount, colCount)

		// Vertical
		for col := 0; col < colCount-1; col++ {
			isSymmetric := true
			for anotherCol := 0; anotherCol < colCount; anotherCol++ {
				colLeft := col - anotherCol
				colRight := col + 1 + anotherCol
				if 0 <= colLeft && colLeft < colRight && colRight < colCount {
					for row := 0; row < rowCount; row++ {
						if lines[row][colLeft] != lines[row][colRight] {
							isSymmetric = false
						}
					}
				}
			}
			if isSymmetric {
				answer += col + 1
			}
		}
		for row := 0; row < rowCount-1; row++ {
			isSymmetric := true
			for anotherRow := 0; anotherRow < rowCount; anotherRow++ {
				above := row - anotherRow
				down := row + 1 + anotherRow
				if 0 <= above && above < down && down < rowCount {
					for col := 0; col < colCount; col++ {
						if lines[above][col] != lines[down][col] {
							isSymmetric = false
						}
					}
				}
			}
			if isSymmetric {
				answer += 100 * (row + 1)
			}
		}
	}

	fmt.Println("Answer:", answer)
}
