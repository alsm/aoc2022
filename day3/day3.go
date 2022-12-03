package main

import (
	"log"

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
	return Sum(Map(groups, func (x []string) int {
		return slices.Index(charValues, Intersection([]rune(x[0]), []rune(x[1]), []rune(x[2]))[0])
	}))
}