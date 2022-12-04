package main

import (
	"fmt"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
)

func main() {
	input := aoc.SliceFromFile("day4.txt", func(i string) [4]int {
		var a [4]int
		fmt.Sscanf(i, "%d-%d,%d-%d", &a[0], &a[1], &a[2], &a[3])
		return a
	})

	fmt.Println(do1(input))
	fmt.Println(do2(input))
}

func do1(in [][4]int) int {
	return Count(in, func(a [4]int) bool {
		return a[0] >= a[2] && a[1] <= a[3] || a[2] >= a[0] && a[3] <= a[1]
	})
}

func toRange(x,y int) []int {
	var ret []int
	for i:= x; i <= y; i++ {
		ret = append(ret, i)
	}

	return ret
}

func do2(in [][4]int) int {
	return Count(in, func(a [4]int) bool {
		return len(Intersection(toRange(a[0],a[1]), toRange(a[2],a[3]))) > 0
	})
}