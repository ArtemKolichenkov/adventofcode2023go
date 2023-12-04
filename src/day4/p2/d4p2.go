package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`\d+`)

func main() {
	answer := Solve("./src/day4/p2/input.txt")
	fmt.Println("Answer:", answer) // Answer: 5667240
}

type Card struct {
	cardNumber     int
	winningNumbers []string
	myNumbers      []string
	winCount       int
}

func Solve(filePath string) int {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	rawStrings := strings.Split(string(file), "\n")
	originalCards := make([]Card, len(rawStrings))
	for i, cardString := range rawStrings {
		originalCards[i] = *scratchCard(parseCardString(cardString))
	}
	// You can also use copy() here, but it's a little bit slower than append
	cardsQueue := make([]Card, 0, len(rawStrings))
	cardsQueue = append(cardsQueue, originalCards...)
	totalCards := 0

	for len(cardsQueue) > 0 {
		totalCards += 1
		for i := 0; i < cardsQueue[0].winCount; i++ {
			if cardsQueue[0].cardNumber+i <= len(originalCards) {
				cardsQueue = append(cardsQueue, originalCards[cardsQueue[0].cardNumber+i])
			}
		}
		cardsQueue = cardsQueue[1:]

	}
	return totalCards
}

func scratchCard(card *Card) *Card {
	winCount := 0
	for _, num := range card.myNumbers {
		if slices.Contains(card.winningNumbers, num) {
			winCount++
		}
	}
	card.winCount = winCount
	return card
}

func parseCardString(cardString string) *Card {
	cardNumber, err := strconv.Atoi(re.FindString(strings.Split(cardString, ":")[0]))
	if err != nil {
		panic(err)
	}
	winningNumbersRaw := strings.Split(strings.TrimSpace(strings.Split(strings.Split(cardString, ":")[1], "|")[0]), " ")
	myNumbersRaw := strings.Split(strings.TrimSpace(strings.Split(strings.Split(cardString, ":")[1], "|")[1]), " ")
	winningNumbers := filterNonNumbers(winningNumbersRaw)
	myNumbers := filterNonNumbers(myNumbersRaw)
	return &Card{cardNumber, winningNumbers, myNumbers, 0}
}

func filterNonNumbers(s []string) []string {
	var r []string
	for _, str := range s {
		if _, err := strconv.Atoi(str); err == nil {
			r = append(r, str)
		}
	}
	return r
}
