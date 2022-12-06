package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/alsm/aoc2022/aoc/collections"
)

func main() {
	data, _ := os.ReadFile("day6.txt")
	input := strings.TrimSpace(string(data))

	fmt.Println(do(input, 4))
	fmt.Println(do(input, 14))
}

func do(in string, size int) int {
	for i := 0; i < len(in)-size; i++ {
		if len(Tally([]rune(in[i:i+size]))) == size {
			return i + size
		}
	}

	return 0
}
