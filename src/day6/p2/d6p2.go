package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Answer: 35961505
func main() {
	var err error
	file, err := os.ReadFile("./src/day6/p2/input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(file), "\n")
	duration := getNumber(strings.ReplaceAll(strings.Split(inputStrings[0], ":")[1], " ", ""))
	record := getNumber(strings.ReplaceAll(strings.Split(inputStrings[1], ":")[1], " ", ""))

	waysToBeatThisRecord := 0
	for holdSeconds := 0; holdSeconds < duration; holdSeconds++ {
		if holdSeconds*(duration-holdSeconds) > record {
			waysToBeatThisRecord++
		}
	}

	fmt.Println("Answer:", waysToBeatThisRecord)
}

func getNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
