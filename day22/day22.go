package main

import (
	"github.com/alsm/aoc2022/aoc"
	"github.com/alsm/aoc2022/aoc/grid"
	"github.com/gammazero/deque"
)

func main() {
	checks := deque.New[[]aoc.Point]()
	checks.PushBack([]aoc.Point{{-1, -1}, {0, -1}, {1, -1}})
	checks.PushBack([]aoc.Point{{-1, 1}, {0, 1}, {1, 1}})
	checks.PushBack([]aoc.Point{{-1, -1}, {-1, 0}, {-1, 1}})
	checks.PushBack([]aoc.Point{{1, -1}, {1, 0}, {1, 1}})

	grove := grid.NewIGrid[string](grid.Directions8)

	lines := aoc.SliceFromFile("day23.txt", func(s string) string { return s })
	for y, l := range lines {
		for x, c := range l {
			if c == '.' {
				continue
			}
			grove.SetState(aoc.Point{int64(x), int64(y)}, "#")
		}
	}
}
