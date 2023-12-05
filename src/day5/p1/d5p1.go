package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) AddNode(value int) {
	newNode := &Node{
		value: value,
		prev:  nil,
		next:  nil,
	}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

type DestinationMapping struct {
	sourceStart      int
	sourceEnd        int
	destinationStart int
	destinationEnd   int
}

func (destinationMap *DestinationMapping) FindDestination(input int, destinationVar *int) {
	if input < destinationMap.sourceStart || input > destinationMap.sourceEnd {
		return
	}
	offset := input - destinationMap.sourceStart
	*destinationVar = destinationMap.destinationStart + offset
}

// Answer: 825516882
func main() {
	var err error
	file, err := os.ReadFile("./src/day5/p1/input.txt")
	if err != nil {
		panic(err)
	}

	inputStrings := strings.Split(string(file), "\n")
	seedsAsStrings := strings.Fields(strings.Split(inputStrings[0], ":")[1])
	inputStrings = inputStrings[3:]
	seeds := make([]int, len(seedsAsStrings))
	for i := range seedsAsStrings {
		seeds[i] = getNumber(seedsAsStrings[i])
	}

	seedsMap := make(map[int]*DoublyLinkedList)
	for _, seed := range seeds {
		seedDLL := DoublyLinkedList{}
		seedDLL.AddNode(seed)
		seedsMap[seed] = &seedDLL
	}

	destinationMap := []DestinationMapping{}
	for lineNumber, line := range inputStrings {
		if line == "" {
			continue
		}
		if !unicode.IsDigit(rune(line[0])) || lineNumber == len(inputStrings)-1 {
			for i := range seedsMap {
				seedDestination := seedsMap[i].tail.value
				for j := range destinationMap {
					destinationMap[j].FindDestination(seedsMap[i].tail.value, &seedDestination)
				}
				seedsMap[i].AddNode(seedDestination)
			}
			destinationMap = []DestinationMapping{}
		} else {
			nums := strings.Fields(line)
			destinationStart, sourceStart, numRange := getNumber(nums[0]), getNumber(nums[1]), getNumber(nums[2])
			destinationMap = append(destinationMap, DestinationMapping{
				sourceStart,
				sourceStart + numRange - 1,
				destinationStart,
				destinationStart + numRange - 1,
			})
		}
	}

	minLocation := math.MaxInt
	for i := range seedsMap {
		if seedsMap[i].tail.value < minLocation {
			minLocation = seedsMap[i].tail.value
		}
	}

	fmt.Println("Answer:", minLocation)
}

func getNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
