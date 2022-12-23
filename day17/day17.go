package main

import (
	"fmt"
	"os"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/grid"
	"github.com/alsm/aoc2022/aoc/ring"
)

func main() {
	data, _ := os.ReadFile("day17.txt")
	input := ring.NewFromSlice(Map(data, func(b byte) int64 {
		return map[byte]int64{'<': -1, '>': 1}[b]
	}))

	fmt.Println(do1(input))
	// input.Do(func(v int64) {
	// 	fmt.Println(v)
	// })
}

type Rock []aoc.Point

func (r Rock) Bottom() []aoc.Point {
	ret := []aoc.Point{r[0]}
	for _, p := range r[1:] {
		switch {
		case p.Y < ret[0].Y:
			ret = []aoc.Point{p}
		case p.Y == ret[0].Y:
			ret = append(ret, p)
		}
	}

	return ret
}

func (r Rock) Top() []aoc.Point {
	ret := []aoc.Point{r[0]}
	for _, p := range r[1:] {
		switch {
		case p.Y > ret[0].Y:
			ret = []aoc.Point{p}
		case p.Y == ret[0].Y:
			ret = append(ret, p)
		}
	}

	return ret
}

func (r Rock) Left() []aoc.Point {
	ret := []aoc.Point{r[0]}
	for _, p := range r[1:] {
		switch {
		case p.X < ret[0].X:
			ret = []aoc.Point{p}
		case p.X == ret[0].X:
			ret = append(ret, p)
		}
	}

	return ret
}

func (r Rock) Right() []aoc.Point {
	ret := []aoc.Point{r[0]}
	for _, p := range r[1:] {
		switch {
		case p.X > ret[0].X:
			ret = []aoc.Point{p}
		case p.X == ret[0].X:
			ret = append(ret, p)
		}
	}

	return ret
}

var rocks = []Rock{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},          //h line
	{{0, 0}, {-1, 1}, {0, 1}, {1, 1}, {0, 2}}, //plus
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},  //L
	{{0, 0}, {0, 1}, {0, 2}, {0, 2}},          //v line
	{{0, 0}, {1, 0}, {0, 1}, {1, 1}},          //cube
}

type Cave struct {
	*grid.Grid[string]
}

func (c *Cave) Draw(p aoc.Point, rock []aoc.Point) {
	for _, rp := range rock {
		c.SetStateP(p.Add(rp), "#")
	}
}

func do1(in *ring.Ring[int64]) int {
	var level int64
	cave := Cave{Grid: grid.NewWithDefault[string](7, 10, grid.Directions4, ".")}
	for i := 0; i < 1; i++ {
		shape := rocks[i%5]
		loc := aoc.Point{2, level + 3}
		for loc.Y >= 0 {
			action := in.Value
			in = in.Next()
			fmt.Println("moving shape", action)
			if All(append(shape.Left(), shape.Right()...), func(p aoc.Point) bool {
				nx := loc.X + p.X + action
				ny := loc.Y + p.Y
				return cave.IsValid(nx, ny) && cave.GetState(nx, ny) == "."
			}) {
				loc = loc.Add(aoc.Point{action, 0})
			}
			if Any(shape.Bottom(), func(p aoc.Point) bool {
				nx := loc.X + p.X
				ny := loc.Y + p.Y - 1
				return cave.IsValid(nx, ny) && cave.GetState(nx, ny) == "#"
			}) {
				if cave.IsValidPoint(loc.Add(aoc.Point{0, -1})) {
					loc = loc.Add(aoc.Point{0, -1})
				}
				break
			}
			if cave.IsValidPoint(loc.Add(aoc.Point{0, -1})) {
				loc = loc.Add(aoc.Point{0, -1})
			}
		}
		fmt.Println("drawing shape")
		cave.Draw(loc, shape)
		level = loc.Add(shape.Top()[0]).Y
		fmt.Println(cave.StateStringInvertY())
	}

	return 0
}
