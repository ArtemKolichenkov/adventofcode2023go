package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type BoxItem struct {
	label string
	value string
}

func main() {
	file, err := os.ReadFile("./src/day15/p2/input.txt")
	if err != nil {
		panic(err)
	}

	boxes := make([][]BoxItem, 256)
	for _, str := range strings.Split(string(file), ",") {
		if str[len(str)-1] == '-' {
			label := str[:len(str)-1]
			boxId := hash(label)
			boxes[boxId] = slices.DeleteFunc(boxes[boxId], func(boxItems BoxItem) bool {
				return boxItems.label == label
			})
			continue
		}
		// "=" case, no need for redundant check
		label := str[:len(str)-2]
		boxId := hash(label)
		focalLength := string(str[len(str)-1])
		if slices.ContainsFunc(boxes[boxId], func(boxItem BoxItem) bool {
			return boxItem.label == label
		}) {
			for i, boxItem := range boxes[boxId] {
				if boxItem.label == label {
					boxes[boxId][i].value = focalLength
				}
			}
		} else {
			boxes[boxId] = append(boxes[boxId], BoxItem{label, focalLength})
		}
	}

	answer := 0
	for i, box := range boxes {
		for j, boxItem := range box {
			inVal, err := strconv.Atoi(boxItem.value)
			if err != nil {
				panic(err)
			}
			answer += (i + 1) * (j + 1) * inVal
		}
	}
	fmt.Println("Answer:", answer) // 251353
}

func hash(str string) (hash int) {
	for i := 0; i < len(str); i++ {
		hash = ((hash + int(str[i])) * 17) % 256
	}
	return hash
}
