package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	isHighCard = iota
	isOnePair
	isTwoPair
	isThreeOfAKind
	isFullHouse
	isFourOfAKind
	isFiveOfAKind
)

type Hand struct {
	hand            string
	handType        int
	bid             int
	winning         int
	binaryFootprint string
}

func (thisHand *Hand) isStrongerThan(otherHand *Hand) bool {
	for i := 0; i < 64*5; i++ {
		if thisHand.binaryFootprint[i] == otherHand.binaryFootprint[i] {
			continue
		}
		return string(thisHand.binaryFootprint[i]) == "1"
	}
	return false
}

// P1 answer 253205868
func main() {
	file, err := os.ReadFile("./src/day7/p1/input.txt")
	if err != nil {
		panic(err)
	}

	answer := 0

	inputStrings := strings.Split(string(file), "\n")
	hands := make([]Hand, len(inputStrings))

	for lineNumber, line := range inputStrings {
		hand := strings.Split(line, " ")[0]
		bid := getNumber(strings.Split(line, " ")[1])

		charMap := make(map[string]int)
		maxOfAKind := 0
		pairCount := 0
		for i := 0; i < len(hand); i++ {
			charMap[string(hand[i])]++
			if charMap[string(hand[i])] > maxOfAKind {
				maxOfAKind = charMap[string(hand[i])]
			}
			if charMap[string(hand[i])] == 2 {
				pairCount++
			}
			if charMap[string(hand[i])] == 3 {
				pairCount--
			}
		}

		handType := isHighCard
		if maxOfAKind == 5 {
			handType = isFiveOfAKind
		}
		if maxOfAKind == 4 {
			handType = isFourOfAKind
		}
		if maxOfAKind == 3 {
			for _, count := range charMap {
				if count == 2 {
					handType = isFullHouse
				}
				if count == 1 {
					handType = isThreeOfAKind
				}
			}
		}
		if pairCount == 2 {
			handType = isTwoPair
		}
		if maxOfAKind == 2 && pairCount == 1 {
			handType = isOnePair
		}

		binaryFootprint := ""
		for i := 0; i < len(hand); i++ {
			cardValue := getCardValue(string(hand[i]))
			cardAsBinary := strconv.FormatInt(int64(cardValue), 2)
			binaryFootprint += fmt.Sprintf("%0*s", 64, cardAsBinary)
		}
		hands[lineNumber] = Hand{hand, handType, bid, 0, binaryFootprint}
	}

	categories := map[int][]Hand{}
	categoryKeys := []int{isHighCard, isOnePair, isTwoPair, isThreeOfAKind, isFullHouse, isFourOfAKind, isFiveOfAKind}
	for _, hand := range hands {
		categories[hand.handType] = append(categories[hand.handType], hand)
	}

	currentRank := 1
	for _, cat := range categoryKeys {
		if len(categories[cat]) == 1 {
			answer += categories[cat][0].bid * currentRank
			currentRank++
			continue
		}
		for _, hand := range categories[cat] {
			beatsHands := 0
			for _, otherHand := range categories[cat] {
				if hand.hand == otherHand.hand {
					continue
				}
				if hand.isStrongerThan(&otherHand) {
					beatsHands++
				}
			}
			answer += hand.bid * (currentRank + beatsHands)
		}
		currentRank += len(categories[cat])
	}

	fmt.Println("Answer:", answer)
}

func getCardValue(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		return getNumber(card)
	}
}

func getNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
