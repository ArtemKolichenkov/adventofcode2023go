package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var MAX_CUBES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Answer: 2105
func main() {
	file, err := os.Open("./src/day2/p1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumOfPossibleGameIds := 0
	for scanner.Scan() {
		gameMetaAndSets := strings.Split(scanner.Text(), ":")
		gameId, err := strconv.Atoi(strings.Split(gameMetaAndSets[0], " ")[1])
		if err != nil {
			panic(err)
		}
		setsOfCubes := strings.Split(gameMetaAndSets[1], ";")
		isGamePossible := true
		for _, setOfCubes := range setsOfCubes {
			for _, cube := range strings.Split(setOfCubes, ",") {
				cubeCountAndColor := strings.Split(strings.Trim(cube, " "), " ")
				cubeCount, err := strconv.Atoi(string(cubeCountAndColor[0]))
				if err != nil {
					panic(err)
				}
				color := string(cubeCountAndColor[1])
				if cubeCount > MAX_CUBES[color] {
					isGamePossible = false
				}
			}
		}
		if isGamePossible {
			sumOfPossibleGameIds += gameId
		}
	}

	fmt.Println("Sum:", sumOfPossibleGameIds)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
