package main

import (
	"fmt"
	"strconv"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/grid"
	"golang.org/x/exp/slices"
)

func main() {
	data := aoc.SliceFromFile("day8.txt", func(i string) []int {
		return Map([]rune(i), func(c rune) int {
			v, _ := strconv.Atoi(string(c))
			return v
		})
	})
	input := grid.New[int](int64(len(data[0])), int64(len(data)), grid.Directions4)
	for y, l := range data {
		for x, v := range l {
			input.SetState(int64(x), int64(y), v)
		}
	}

	fmt.Println(do1(input))
	fmt.Println(do2(input))
}

func do1(in *grid.Grid[int]) int {
	var trees int
	for x := int64(1); x < in.XLen()-1; x++ {
		for y := int64(1); y < in.YLen()-1; y++ {
			routes := Map(grid.Directions4, func(i aoc.Point) []int {
				return in.GetSliceToEdge(x, y, i)
			})
			if Any(Map(routes, func(i []int) bool {
				return All(i[1:], func(j int) bool {
					return i[0] > j
				})

			}), func(visible bool) bool {
				return visible
			}) {
				trees++
			}
		}
	}
	return trees + int(in.XLen())*2 + int(in.YLen())*2 - 4
}

func do2(in *grid.Grid[int]) int {
	var bestScore int
	for x := int64(0); x < in.XLen(); x++ {
		for y := int64(0); y < in.YLen(); y++ {
			scenes := Map(grid.Directions4, func(i aoc.Point) []int {
				return in.GetSliceToEdge(x, y, i)
			})
			score := Product(Map(scenes, func(sight []int) int {
				index := slices.IndexFunc(sight[1:], func(i int) bool {
					return i >= sight[0]
				})
				if index == -1 {
					if All(sight[1:], func(i int) bool {
						return i < sight[0]
					}) {
						return len(sight[1:])
					}
					return 0
				}
				return len(sight[1 : index+2])
			}))
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}
