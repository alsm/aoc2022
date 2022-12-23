package main

import (
	"fmt"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/grid"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	checks := [][]aoc.Point{
		{{-1, -1}, {0, -1}, {1, -1}},
		{{-1, 1}, {0, 1}, {1, 1}},
		{{-1, -1}, {-1, 0}, {-1, 1}},
		{{1, -1}, {1, 0}, {1, 1}},
	}

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

	fmt.Println(do1(grove.Clone(), slices.Clone(checks)))
	fmt.Println(do2(grove.Clone(), slices.Clone(checks)))
}

func rotate(s [][]aoc.Point) [][]aoc.Point {
	return append(s[1:], s[0])
}

func filterOut(m *map[aoc.Point]aoc.Point, v aoc.Point) {
	for f, t := range *m {
		if t == v {
			(*m)[f] = f
		}
	}
}

func do1(grove *grid.IGrid[string], checks [][]aoc.Point) int {
	for i := 0; i < 10; i++ {
		grove.Iterate(func(m map[aoc.Point]string) map[aoc.Point]string {
			ret := make(map[aoc.Point]string)
			newlocs := make(map[aoc.Point]aoc.Point)
			notMove, move := Partition(maps.Keys(m), func(p aoc.Point) bool {
				return len(grove.Neighbours(p)) == 0
			})

			for _, p := range move {
				for _, c := range checks {
					if len(grove.NeighboursAt(p, c...)) == 0 {
						newlocs[p] = p.Add(c[1])
						break
					}
					newlocs[p] = p
				}
			}

			for p, c := range Tally(maps.Values(newlocs)) {
				if c > 1 {
					filterOut(&newlocs, p)
				}
			}

			for _, p := range append(maps.Values(newlocs), notMove...) {
				ret[p] = "#"
			}

			return ret
		})
		checks = rotate(checks)
	}

	xmin, xmax := MinMax(Map(maps.Keys(grove.States()), func(p aoc.Point) int64 {
		return p.X
	}))
	ymin, ymax := MinMax(Map(maps.Keys(grove.States()), func(p aoc.Point) int64 {
		return p.Y
	}))

	return int(ymax+1-ymin)*int(xmax+1-xmin) - int(len(maps.Keys(grove.States())))
}

func do2(grove *grid.IGrid[string], checks [][]aoc.Point) int {
	for i := 1; ; i++ {
		grove.Iterate(func(m map[aoc.Point]string) map[aoc.Point]string {
			ret := make(map[aoc.Point]string)
			newlocs := make(map[aoc.Point]aoc.Point)
			notMove, move := Partition(maps.Keys(m), func(p aoc.Point) bool {
				return len(grove.Neighbours(p)) == 0
			})

			if len(notMove) == len(grove.States()) {
				return nil
			}

			for _, p := range move {
				for _, c := range checks {
					if len(grove.NeighboursAt(p, c...)) == 0 {
						newlocs[p] = p.Add(c[1])
						break
					}
					newlocs[p] = p
				}
			}

			for p, c := range Tally(maps.Values(newlocs)) {
				if c > 1 {
					filterOut(&newlocs, p)
				}
			}

			for _, p := range append(maps.Values(newlocs), notMove...) {
				ret[p] = "#"
			}

			return ret
		})
		checks = rotate(checks)
		if len(maps.Keys(grove.States())) == 0 {
			return i
		}
	}
}
