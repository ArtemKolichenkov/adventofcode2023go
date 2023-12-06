package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Answer: 1155175
func main() {
	var err error
	file, err := os.ReadFile("./src/day6/p1/input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(file), "\n")
	durations := strings.Fields(strings.Split(inputStrings[0], ":")[1])
	records := strings.Fields(strings.Split(inputStrings[1], ":")[1])

	waysToBeatRecords := []int{}

	for i := 0; i < len(durations); i++ {
		dur := getNumber(durations[i])
		rec := getNumber(records[i])

		waysToBeatThisRecord := 0
		for holdSeconds := 0; holdSeconds < dur; holdSeconds++ {
			if holdSeconds*(dur-holdSeconds) > rec {
				waysToBeatThisRecord++
			}
		}
		waysToBeatRecords = append(waysToBeatRecords, waysToBeatThisRecord)
	}

	answer := waysToBeatRecords[0]

	for _, ways := range waysToBeatRecords[1:] {
		answer *= ways
	}

	fmt.Println("Answer:", answer)
}

func getNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
