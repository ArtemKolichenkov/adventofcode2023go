package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var digitRegex = regexp.MustCompile(`\d`)

// Answer: 53386
func main() {
	file, err := os.Open("./src/day1/p1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		digits := digitRegex.FindAllString(scanner.Text(), -1)
		twoDigitNumber, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
		if err != nil {
			panic(err)
		}
		sum += twoDigitNumber
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
