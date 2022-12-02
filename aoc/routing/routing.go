package routing

import (
	"github.com/alsm/aoc2022/aoc"
)

type Graph interface {
	Neighbours(p aoc.Point) []aoc.Point
}
