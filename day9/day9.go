package main

import (
	"fmt"
	"math"

	"github.com/alsm/aoc2022/aoc"
	"github.com/alsm/aoc2022/aoc/grid"
)

func main() {
	input := aoc.SliceFromFile("day9.txt", func(l string) [2]int {
		var dir rune
		var mag int
		dirMap := map[rune]int{'U': 0, 'R': 1, 'D': 2, 'L': 3}
		fmt.Sscanf(l, "%c %d", &dir, &mag)
		return [2]int{dirMap[dir], mag}
	})

	fmt.Println(do1(input))
	fmt.Println(do2(input))
}

var tailDiagMoves = [4]aoc.Point{{X: 1, Y: 1}, {X: -1, Y: 1}, {X: -1, Y: -1}, {X: 1, Y: -1}}
var tailMoves = [4]aoc.Point{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}}

func getDir(t, h *aoc.Point, moves [4]aoc.Point) aoc.Point {
	dir := math.Atan2(float64(t.Y-h.Y), float64(t.X-h.X))
	move := int((dir*4)/(2*math.Pi)+4) % 4
	return moves[move]
}

func moveHead(p *aoc.Point, d int) {
	*p = p.Add(grid.Directions4[d])
}

func moveTail(t, h *aoc.Point) {
	switch {
	case t.Neighbour(*h):
	case h.X-t.X == 0 || h.Y-t.Y == 0:
		*t = t.Add(getDir(h, t, tailMoves))
	default:
		*t = t.Add(getDir(h, t, tailDiagMoves))
	}
}

func do1(in [][2]int) int {
	var head, tail aoc.Point
	positions := grid.NewIGrid[rune](grid.Directions4)
	positions.SetState(tail, '#')

	for _, m := range in {
		for i := 0; i < m[1]; i++ {
			moveHead(&head, m[0])
			moveTail(&tail, &head)
			positions.SetState(tail, '#')
		}
	}

	return len(positions.States())
}

func do2(in [][2]int) int {
	rope := make([]aoc.Point, 10)
	positions := grid.NewIGrid[rune](grid.Directions4)
	positions.SetState(rope[9], '#')

	for _, m := range in {
		for i := 0; i < m[1]; i++ {
			moveHead(&rope[0], m[0])
			for j := 0; j < 9; j++ {
				moveTail(&rope[j+1], &rope[j])
			}
			positions.SetState(rope[9], '#')
		}
	}

	return len(positions.States())
}
