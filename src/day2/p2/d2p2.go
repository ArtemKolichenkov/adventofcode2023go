package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Answer: 72422
func main() {
	file, err := os.Open("./src/day2/p2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPower := 0
	for scanner.Scan() {
		gameMetaAndSets := strings.Split(scanner.Text(), ":")
		setsOfCubes := strings.Split(gameMetaAndSets[1], ";")
		minCubesNeeded := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, setOfCubes := range setsOfCubes {
			for _, cube := range strings.Split(setOfCubes, ",") {
				countAndColor := strings.Split(strings.Trim(cube, " "), " ")
				cubeCount, err := strconv.Atoi(string(countAndColor[0]))
				if err != nil {
					panic(err)
				}
				color := string(countAndColor[1])
				if cubeCount > minCubesNeeded[color] {
					minCubesNeeded[color] = cubeCount
				}
			}
		}
		totalPower += minCubesNeeded["red"] * minCubesNeeded["green"] * minCubesNeeded["blue"]
	}

	fmt.Println("Total power:", totalPower)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
