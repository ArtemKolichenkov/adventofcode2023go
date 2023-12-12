package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var hashMap = map[string]int{}

func main() {
	file, err := os.ReadFile("./src/day12/p1/input.txt")
	if err != nil {
		panic(err)
	}
	inputStrings := strings.Split(string(file), "\n")

	totalArrangements := 0
	for _, line := range inputStrings {
		spaceSplit := strings.Split(line, " ")
		springs := strings.Split(spaceSplit[0], "")
		blocksStr := strings.Split(spaceSplit[1], ",")
		blocks := make([]int, len(blocksStr))
		for i, block := range blocksStr {
			blocks[i], err = strconv.Atoi(block)
			if err != nil {
				panic(err)
			}
		}

		hashMap = map[string]int{}
		totalArrangements += solve(springs, blocks, 0, 0, 0)
	}

	fmt.Println("Answer:", totalArrangements) // 8270
}

func solve(springs []string, blocks []int, springIndex int, blockIndex int, current int) int {
	hashKey := strconv.Itoa(springIndex) + strconv.Itoa(blockIndex) + strconv.Itoa(current)
	if value, exists := hashMap[hashKey]; exists {
		return value
	}
	if springIndex == len(springs) {
		if blockIndex == len(blocks) && current == 0 {
			return 1
		} else if blockIndex == len(blocks)-1 && blocks[blockIndex] == current {
			return 1
		} else {
			return 0
		}
	}
	answer := 0
	for _, char := range []string{".", "#"} {
		if springs[springIndex] == char || springs[springIndex] == "?" {
			if char == "." {
				if current == 0 {

					answer += solve(springs, blocks, springIndex+1, blockIndex, 0)
				}
				if current > 0 && blockIndex < len(blocks) && blocks[blockIndex] == current {

					answer += solve(springs, blocks, springIndex+1, blockIndex+1, 0)
				}
			} else if char == "#" {
				answer += solve(springs, blocks, springIndex+1, blockIndex, current+1)
			}
		}
	}
	hashMap[hashKey] = answer
	return answer
}
