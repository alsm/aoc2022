package routing

import (
	"github.com/alsm/aoc2017/aoc"
)

type Graph interface {
	Neighbours(p aoc.Point) []aoc.Point
}