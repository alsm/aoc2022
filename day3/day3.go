package main

import (
	"log"
	"math"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"golang.org/x/exp/slices"
)

var (
	charValues []rune = []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func main() {
	input := aoc.SliceFromFile("day3.txt", func(i string) string {
		return i
	})

	log.Println(do1(input))
	log.Println(do2(input))
	log.Println(do1second(input))
	log.Println(do2second(input))
}

func do1(in []string) int {
	cmps := Map(in, func(in string) [][]rune {
		return Chunk([]rune(in), len(in)/2)
	})
	return Sum(Map(cmps, func(x [][]rune) int {
		return slices.Index(charValues, Intersection(x[0], x[1])[0])
	}))
}

func do2(in []string) int {
	groups := Chunk(in, 3)
	return Sum(Map(groups, func(x []string) int {
		return slices.Index(charValues, Intersection([]rune(x[0]), []rune(x[1]), []rune(x[2]))[0])
	}))
}

func toBits(s string) int64 {
	var ret int64
	for _, c := range s {
		ret |= 1 << slices.Index(charValues, c)
	}

	return ret
}

func do1second(in []string) int {
	var ret int
	for _, l := range in {
		ret += int(math.Log2(float64(toBits(l[:len(l)/2]) & toBits(l[len(l)/2:]))))
	}

	return ret
}

func do2second(in []string) int {
	var ret int
	val := int64(math.MaxInt64)
	for i, l := range in {
		val &= toBits(l)
		if (i+1)%3 == 0 {
			ret += int(math.Log2(float64(val)))
			val = math.MaxInt64
		}
	}

	return ret
}
