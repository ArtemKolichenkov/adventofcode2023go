package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./src/day9/p1/input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(file), "\n")

	answer := 0
	for _, line := range inputStrings {
		byWord := strings.Fields(line)
		numbers := []int{}
		for _, word := range byWord {
			if num, err := strconv.Atoi(word); err == nil {
				numbers = append(numbers, num)
			}
		}
		rex := recursiveExpand(numbers)
		answer += rex
	}

	fmt.Println("Answer:", answer)
}

func recursiveExpand(numbers []int) int {
	if allZeros(numbers) {
		return 0
	}
	nextArr := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		nextArr[i] = numbers[i+1] - numbers[i]
	}
	lastEl := numbers[len(numbers)-1]
	return lastEl + recursiveExpand(nextArr)
}

func allZeros(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != 0 {
			return false
		}
	}
	return true
}
