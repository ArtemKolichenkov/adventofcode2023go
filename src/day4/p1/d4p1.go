package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./src/day4/p1/input.txt")
	if err != nil {
		panic(err)
	}

	ans := 0

	for _, line := range strings.Split(string(file), "\n") {
		winningNumbersRaw := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[0]), " ")
		myNumbersRaw := strings.Split(strings.TrimSpace(strings.Split(strings.Split(line, ":")[1], "|")[1]), " ")
		winningNumbers := filterNonNumbers(winningNumbersRaw)
		myNumbers := filterNonNumbers(myNumbersRaw)

		curPoint := 0
		for _, num := range myNumbers {
			if slices.Contains(winningNumbers, num) {
				if curPoint == 0 {
					curPoint = 1
				} else {
					curPoint = curPoint * 2
				}
			}
		}
		ans += curPoint
	}

	fmt.Println("Answer:", ans)
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
