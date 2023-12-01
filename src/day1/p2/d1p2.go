package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var digitRegex = regexp.MustCompile(`\d|oneight|one|twone|two|threeight|three|four|fiveight|five|six|sevenine|seven|eightwo|eighthree|eight|nine`)

var wordToNum = map[string]string{
	"oneight":   "18",
	"one":       "1",
	"twone":     "21",
	"two":       "2",
	"threeight": "38",
	"three":     "3",
	"four":      "4",
	"fiveight":  "58",
	"five":      "5",
	"six":       "6",
	"sevenine":  "79",
	"seven":     "7",
	"eightwo":   "82",
	"eighthree": "83",
	"eight":     "8",
	"nine":      "9",
}

// Answer: 53312
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
		firstDigit := convertDigit(digits[0], true)
		lastDigit := convertDigit(digits[len(digits)-1], false)
		twoDigitNumber, err := strconv.Atoi(firstDigit + lastDigit)
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

func convertDigit(s string, first bool) string {
	if len(s) > 1 {
		tmpDigit := wordToNum[s]
		if len(tmpDigit) == 1 {
			return tmpDigit
		}
		if first {
			return string(tmpDigit[0])
		}
		return string(tmpDigit[1])
	}
	return s
}
