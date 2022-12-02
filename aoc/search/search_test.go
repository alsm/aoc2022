package search

import (
	"fmt"
	"testing"

	"github.com/alsm/aoc2022/aoc"
	"github.com/alsm/aoc2022/aoc/grid"
)

func TestBFSTestBFS(t *testing.T) {
	g := grid.New[int](10, 10, grid.Directions4)

	route := BFS(g, aoc.Point{0, 0}, aoc.Point{5, 5})

	fmt.Println(route)
}
