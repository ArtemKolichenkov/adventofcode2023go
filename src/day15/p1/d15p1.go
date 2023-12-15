package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./src/day15/p1/input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for _, str := range strings.Split(string(file), ",") {
		total += hash(str)
	}
	fmt.Println("Answer:", total) // 503154
}

func hash(str string) (hash int) {
	for i := 0; i < len(str); i++ {
		hash = ((hash + int(str[i])) * 17) % 256
	}
	return hash
}
