package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Range struct {
	start  int
	end    int
	length int
}

type MappingRange struct {
	sourceRange      Range
	destinationRange Range
}

func (thisRange *Range) NotOverlapping(anotherRange *Range) bool {
	return thisRange.start > anotherRange.end || thisRange.end < anotherRange.start
}

func (thisRange *Range) MapRange(mappingRange *MappingRange) (*Range, []Range) {
	unmappedRanges := []Range{}
	if thisRange.NotOverlapping(&mappingRange.sourceRange) {
		unmappedRanges = append(unmappedRanges, *thisRange)
		return nil, unmappedRanges
	}
	elementsBeforeOverlap := 0
	elementsAfterOverlap := 0
	startOffset := thisRange.start - mappingRange.sourceRange.start
	if startOffset < 0 {
		elementsBeforeOverlap = -startOffset
		startOffset = 0
	}
	remainingLength := thisRange.length - elementsBeforeOverlap
	overlapStart := mappingRange.destinationRange.start + startOffset
	overlapEnd := overlapStart + remainingLength - 1
	if overlapEnd > mappingRange.destinationRange.end {
		elementsAfterOverlap = overlapEnd - mappingRange.destinationRange.end
		overlapEnd = mappingRange.destinationRange.end
	}
	overlappingRange := Range{overlapStart, overlapEnd, overlapEnd - overlapStart + 1}
	if elementsBeforeOverlap > 0 {
		rangeBefore := Range{mappingRange.sourceRange.start - elementsBeforeOverlap, mappingRange.sourceRange.start - 1, elementsBeforeOverlap}
		unmappedRanges = append(unmappedRanges, rangeBefore)
	}
	if elementsAfterOverlap > 0 {
		rangeAfter := Range{mappingRange.sourceRange.end + 1, mappingRange.sourceRange.end + elementsAfterOverlap, elementsAfterOverlap}
		unmappedRanges = append(unmappedRanges, rangeAfter)
	}
	return &overlappingRange, unmappedRanges
}

// Answer: 136096660
func main() {
	start := time.Now()
	var err error
	file, err := os.ReadFile("./src/day5/p2/input.txt")
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

	totalSize := 0
	for i := 0; i < len(seeds); i += 2 {
		totalSize += seeds[i+1]
	}
	seedRanges := make([]Range, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedRanges[i/2] = Range{seeds[i], seeds[i] + seeds[i+1] - 1, seeds[i+1]}
	}

	currentMappingRanges := []MappingRange{}
	allDestinationMaps := [][]MappingRange{}
	for lineNumber, line := range inputStrings {
		if line == "" {
			continue
		}
		if !unicode.IsDigit(rune(line[0])) || lineNumber == len(inputStrings)-1 {
			allDestinationMaps = append(allDestinationMaps, currentMappingRanges)
			currentMappingRanges = []MappingRange{}
		} else {
			nums := strings.Fields(line)
			destinationStart, sourceStart, numRange := getNumber(nums[0]), getNumber(nums[1]), getNumber(nums[2])
			currentMappingRanges = append(currentMappingRanges, MappingRange{
				Range{sourceStart, sourceStart + numRange - 1, numRange},
				Range{destinationStart, destinationStart + numRange - 1, numRange},
			})
		}
	}

	for _, mp := range allDestinationMaps {
		tmpSeedRanges := []Range{}
		for len(seedRanges) > 0 {
			mappedSeedRanges := []Range{}
			unmappedSeedRanges := []Range{seedRanges[0]}
			for _, mappingRange := range mp {
				for _, unmappedRange := range unmappedSeedRanges {
					mappedRange, newUnmappedRanges := unmappedRange.MapRange(&mappingRange)
					if mappedRange != nil {
						mappedSeedRanges = append(mappedSeedRanges, *mappedRange)
					}
					if newUnmappedRanges != nil {
						unmappedSeedRanges = newUnmappedRanges
					}
				}
			}
			combinedRanges := append(mappedSeedRanges, unmappedSeedRanges...)
			tmpSeedRanges = append(tmpSeedRanges, combinedRanges...)
			seedRanges = seedRanges[1:]
		}
		seedRanges = tmpSeedRanges
	}

	minLocation := math.MaxInt
	for seedId := range seedRanges {
		if seedRanges[seedId].start < minLocation {
			minLocation = seedRanges[seedId].start
		}
	}

	fmt.Println("Answer:", minLocation)
	duration := time.Since(start)
	fmt.Println("Done in: ", duration.Microseconds(), "μS")
}

func getNumber(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
