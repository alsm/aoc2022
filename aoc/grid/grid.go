package grid

import (
	"fmt"
	"log"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	"golang.org/x/exp/slices"
)

var Directions4 = []aoc.Point{
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
	{X: -1, Y: 0},
}

var Directions8 = []aoc.Point{
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 1, Y: 0},
	{X: 1, Y: -1},
	{X: 0, Y: -1},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
}

type Grid[T any] struct {
	xLen      int64
	yLen      int64
	state     [][]T
	movements []aoc.Point
}

func New[T any](xLen, yLen int64, movements []aoc.Point) *Grid[T] {
	state := make([][]T, yLen)
	for y := int64(0); y < yLen; y++ {
		state[y] = make([]T, xLen)
	}
	return &Grid[T]{
		xLen:      xLen,
		yLen:      yLen,
		state:     state,
		movements: movements,
	}
}

func (g *Grid[T]) isValidPoint(x, y int64) bool {
	switch {
	case x < 0, x >= g.xLen, y < 0, y >= g.yLen:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) Neighbours(p aoc.Point) []aoc.Point {
	log.Println("beighbours")
	var ret []aoc.Point

	for _, m := range g.movements {
		np := p.Add(m)
		if g.isValidPoint(np.X, np.Y) {
			ret = append(ret, np)
		}
	}

	return ret
}

func (g *Grid[T]) SetState(x, y int64, state T) {
	g.state[y][x] = state
}

func (g *Grid[T]) GetState(y, x int64) T {
	return g.state[y][x]
}

func (g *Grid[T]) StateString() string {
	var ret strings.Builder

	for _, y := range g.state {
		for _, x := range y {
			ret.WriteString(fmt.Sprintf("%v", x))
		}
		ret.WriteRune('\n')
	}

	return ret.String()
}

func (g *Grid[T]) Clone() *Grid[T] {
	ng := Grid[T]{
		xLen:      g.xLen,
		yLen:      g.yLen,
		movements: slices.Clone(g.movements),
		state:     make([][]T, g.yLen),
	}

	for yi := range g.state {
		ng.state[yi] = slices.Clone(g.state[yi])
	}

	return &ng
}
