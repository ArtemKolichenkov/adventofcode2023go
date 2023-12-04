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
	quantity       int
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
	cardsLength := len(rawStrings)
	originalCards := make([]Card, cardsLength)
	for i, cardString := range rawStrings {
		originalCards[i] = *scratchCard(parseCardString(cardString))
	}
	totalCards := 0

	for _, card := range originalCards {
		for i := 0; i < card.winCount; i++ {
			if card.cardNumber+i <= cardsLength {
				originalCards[card.cardNumber+i].quantity += card.quantity
			}
		}
		totalCards += card.quantity
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
	return &Card{cardNumber, 1, winningNumbers, myNumbers, 0}
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
