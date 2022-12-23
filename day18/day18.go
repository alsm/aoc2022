package main

import (
	"fmt"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/grid"
	"golang.org/x/exp/maps"
)

func main() {
	lava := grid.NewIGrid3d[struct{}](grid.Directions6)
	_ = aoc.SliceFromFile("day18.txt", func(i string) grid.Point3 {
		var p grid.Point3
		fmt.Sscanf(i, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		lava.SetState(p, struct{}{})
		return p
	})

	fmt.Println(do1(lava))
}

func do1(in *grid.IGrid3d[struct{}]) int {
	return Sum(Map(maps.Keys(in.States()), func(p grid.Point3) int {
		return 6 - len(in.Neighbours(p))
	}))
}

// func do2(in *grid.IGrid3d[struct{}]) int {
// 	b, t := MinMax(Map(maps.Keys(in.States()), func(p grid.Point3) int64 {
// 		return p.X
// 	}))
// 	for x := b; x <= t; x++ {

// 	}
// }

// func PointsWith(g *grid.IGrid3d[struct{}]) []Point3
