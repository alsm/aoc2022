package search

import (
	// "log"

	"github.com/alsm/aoc2022/aoc"
	"github.com/alsm/aoc2022/aoc/queue"
	"github.com/alsm/aoc2022/aoc/routing"
)

func BFS(g routing.Graph, start aoc.Point, goal aoc.Point) []aoc.Point {
	var frontier queue.Queue[aoc.Point]
	frontier.Put(start)

	cameFrom := make(map[aoc.Point]*aoc.Point)
	cameFrom[start] = nil

	for !frontier.Empty() {
		current := frontier.Get()
		if current == goal {
			break
		}

		// log.Println(g.Neighbours(current))

		for _, n := range g.Neighbours(current) {
			if _, ok := cameFrom[n]; !ok {
				frontier.Put(n)
				cameFrom[n] = &current
			}
		}
	}

	ret := []aoc.Point{goal}
	for n := cameFrom[goal]; n != nil; n = cameFrom[*n] {
		ret = append(ret, *n)
	}

	return ret
}
