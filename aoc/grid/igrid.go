package grid

import (
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
		if _, ok := g.states[p]; ok {
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
