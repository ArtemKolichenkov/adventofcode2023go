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

// P2 answer 253907829
func main() {
	file, err := os.ReadFile("./src/day7/p2/input.txt")
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
		var mostFrequent string
		for i := 0; i < len(hand); i++ {
			char := string(hand[i])
			charMap[char]++
			if charMap[char] > charMap[mostFrequent] && char != "J" {
				mostFrequent = char
			}
		}

		charMap[mostFrequent] += charMap["J"]
		charMap["J"] = 0

		maxOfAKind := 0
		pairCount := 0
		for _, count := range charMap {
			if count > maxOfAKind {
				maxOfAKind = count
			}
			if count == 2 {
				pairCount++
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
		return 1
	case "T":
		return 10
	default:
		return getNumber(card)
	}
}

func rankToText(rank int) string {
	switch rank {
	case 0:
		return "High Card"
	case 1:
		return "One Pair"
	case 2:
		return "Two Pairs"
	case 3:
		return "Three of a Kind"
	case 4:
		return "Full House (3 + 2)"
	case 5:
		return "Four of a Kind"
	case 6:
		return "Five of a Kind"
	default:
		return "Unknown"
	}
}

func getNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
