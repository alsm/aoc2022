package main

import (
	"log"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
)

func main() {
	input := aoc.SliceFromFile("day2.txt", func(i string) [2]int {
		a := int(strings.Fields(i)[0][0] - 64)
		b := int(strings.Fields(i)[1][0] - 64 - 23)
		return [2]int{a, b}
	})

	log.Println(do1(input))
	log.Println(do2(input))
}

func do1(game [][2]int) int {
	return Sum(Map(game, func(r [2]int) int {
		switch {
		case r[0] == r[1]:
			return r[1] + 3
		case r[1]-r[0] == 1, r[1]-r[0] == -2:
			return r[1] + 6
		default:
			return r[1]
		}
	}))
}

func do2(game [][2]int) int {
	return Sum(Map(game, func(r [2]int) int {
		switch {
		case r[1] == 2:
			return r[0] + 3
		case r[1] == 3:
			if x := (r[0] + 1) % 3; x == 0 {
				return 9
			} else {
				return x + 6
			}
		default:
			if x := (r[0] - 1) % 3; x == 0 {
				return 3
			} else {
				return x
			}
		}
	}))
}
