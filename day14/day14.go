package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/grid"
)

var startPoint = aoc.Point{X: 200, Y: 0}

func main() {
	var maxY int64
	rocks := aoc.SliceFromFile("day14.txt", func(i string) []aoc.Point {
		var ret []aoc.Point
		for _, p := range strings.Split(i, " -> ") {
			x, y, _ := strings.Cut(p, ",")
			xv, _ := strconv.ParseInt(x, 10, 64)
			yv, _ := strconv.ParseInt(y, 10, 64)
			if yv > maxY {
				maxY = yv
			}
			ret = append(ret, aoc.Point{X: xv - 300, Y: yv})
		}
		return ret
	})

	fmt.Println(do1(rocks, maxY))
	fmt.Println(do2(rocks, maxY))
}

var sandDirections = []aoc.Point{
	{X: 0, Y: 1},
	{X: -1, Y: 1},
	{X: 1, Y: 1},
}

type Cave struct {
	*grid.Grid[string]
}

func NewCave(maxY int64) *Cave {
	return &Cave{
		Grid: grid.NewWithDefault(300, maxY, sandDirections, "."),
	}
}

func (c *Cave) Neighbours(p aoc.Point) []aoc.Point {
	return Select(c.Grid.Neighbours(p), func(x aoc.Point) bool {
		return c.Grid.GetState(x.X, x.Y) != "#" && c.Grid.GetState(x.X, x.Y) != "o"
	})
}

func (c *Cave) DrawRocks(r []aoc.Point) {
	start := r[0]
	for _, next := range r {
		for _, p := range start.Line(next) {
			c.SetState(p.X, p.Y, "#")
		}
		start = next
	}
}

func (c *Cave) AddSand() aoc.Point {
	p := startPoint
	for m := c.Neighbours(p); len(m) != 0; m = c.Neighbours(p) {
		p = m[0]
	}
	if p.Y == c.Grid.YLen()-1 {
		return aoc.Point{X: -1, Y: -1}
	}
	c.SetState(p.X, p.Y, "o")

	return p
}

func do1(rocks [][]aoc.Point, maxY int64) int {
	c := NewCave(maxY + 2)
	for _, r := range rocks {
		c.DrawRocks(r)
	}
	var count int
	for p := c.AddSand(); p.Y != -1; p = c.AddSand() {
		count++
	}

	return count
}

func do2(rocks [][]aoc.Point, maxY int64) int {
	rocks = append(rocks, []aoc.Point{{X: 0, Y: maxY + 2}, {X: 999, Y: maxY + 2}})
	c := NewCave(maxY + 3)
	for _, r := range rocks {
		c.DrawRocks(r)
	}
	var count int
	for p := c.AddSand(); p.Y != 0; p = c.AddSand() {
		count++
	}

	fmt.Println(c.StateString())

	return count + 1
}
