package grid

import (
	"fmt"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	"golang.org/x/exp/maps"
)

type IGrid[T any] struct {
	states    map[aoc.Point]T
	movements []aoc.Point
}

func NewIGrid[T any](movements []aoc.Point) *IGrid[T] {
	return &IGrid[T]{
		states:    make(map[aoc.Point]T),
		movements: movements,
	}
}

func (g *IGrid[T]) Neighbours(p aoc.Point) []aoc.Point {
	var ret []aoc.Point

	for _, m := range g.movements {
		np := p.Add(m)
		if _, ok := g.states[np]; ok {
			ret = append(ret, np)
		}
	}

	return ret
}

func (g *IGrid[T]) NeighboursAt(p aoc.Point, moves ...aoc.Point) []aoc.Point {
	var ret []aoc.Point

	for _, m := range moves {
		np := p.Add(m)
		if _, ok := g.states[np]; ok {
			ret = append(ret, np)
		}
	}

	return ret
}

func (g *IGrid[T]) SetState(p aoc.Point, state T) {
	g.states[p] = state
}

func (g *IGrid[T]) GetState(p aoc.Point) T {
	return g.states[p]
}

func (g *IGrid[T]) States() map[aoc.Point]T {
	return maps.Clone(g.states)
}

func (g *IGrid[T]) Iterate(f func(map[aoc.Point]T) map[aoc.Point]T) {
	g.states = f(g.states)
}

func (g *IGrid[T]) StateString(xmin, xmax, ymin, ymax int64) string {
	var ret strings.Builder

	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if c, ok := g.states[aoc.Point{x, y}]; ok {
				ret.WriteString(fmt.Sprintf("%v", c))
			} else {
				ret.WriteString(".")
			}
		}
		ret.WriteRune('\n')
	}

	return ret.String()
}

func (g *IGrid[T]) Clone() *IGrid[T] {
	return &IGrid[T]{
		states:    maps.Clone(g.states),
		movements: g.movements,
	}
}
