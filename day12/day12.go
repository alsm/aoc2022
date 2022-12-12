package main

import (
	"fmt"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/grid"
	"github.com/alsm/aoc2022/aoc/search"
	"golang.org/x/exp/maps"
)

func main() {
	var start, end aoc.Point
	data := aoc.SliceFromFile("day12.txt", func(l string) []rune {
		return []rune(l)
	})
	input := grid.New[rune](int64(len(data[0])), int64(len(data)), grid.Directions4)
	for y, l := range data {
		for x, s := range l {
			switch s {
			case 'S':
				start = aoc.Point{int64(x), int64(y)}
				s = 'a'
			case 'E':
				end = aoc.Point{int64(x), int64(y)}
				s = 'z'
			}
			input.SetState(int64(x), int64(y), s)
		}
	}

	fmt.Println(do1(HeightMap{input}, start, end))
	fmt.Println(do2(HeightMap{input}, end))
}

type HeightMap struct {
	*grid.Grid[rune]
}

func (h HeightMap) Neighbours(p aoc.Point) []aoc.Point {
	height := h.GetState(p.Y, p.X)
	return Select(h.Grid.Neighbours(p), func(x aoc.Point) bool {
		return h.GetState(x.Y, x.X) <= height+1
	})
}

func do1(in HeightMap, start, end aoc.Point) int {
	return len(search.BFS(in, start, end)) - 1
}

func do2(in HeightMap, end aoc.Point) int {
	return Min(Select(Map(maps.Keys(in.StateMapWhere(func(i rune) bool { return i == 'a' })), func(p aoc.Point) int {
		return len(search.BFS(in, p, end))
	}), func(i int) bool {
		return i > 1
	})) - 1
}
