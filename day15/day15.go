package main

import (
	"fmt"

	"github.com/alsm/aoc2022/aoc"
)

type Sensor struct {
	Location      aoc.Point
	ClosestBeacon aoc.Point
	Distance      int64
}

func main() {
	input := aoc.SliceFromFile("day15.txt", func(i string) Sensor {
		var s Sensor
		fmt.Sscanf(i, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.Location.X, &s.Location.Y, &s.ClosestBeacon.X, &s.ClosestBeacon.Y)
		s.Distance = s.Location.MDistance(s.ClosestBeacon)
		return s
	})

	fmt.Println(do1(input, 2000000))
	fmt.Println(do2(input, 4000000))
}

func do1(in []Sensor, y int64) int {
	covered := make(map[int64]struct{})
	for _, s := range in {
		dy := aoc.Abs(y - s.Location.Y)
		r := s.Distance - dy
		for x := s.Location.X - r; x <= s.Location.X+r; x++ {
			covered[x] = struct{}{}
		}
	}
	for _, s := range in {
		if s.ClosestBeacon.Y == y {
			delete(covered, s.ClosestBeacon.X)
		}
	}
	return len(covered)
}

func Contained(in []Sensor, x, y int64) *Sensor {
	for _, s := range in {
		if s.Location.MDistanceXY(x, y) <= s.Distance {
			return &s
		}
	}

	return nil
}

func do2(in []Sensor, limit int64) int64 {
	for y := int64(0); y <= limit; y++ {
		for x := int64(0); x <= limit; x++ {
			s := Contained(in, x, y)
			if s == nil {
				return x*4000000 + y
			}
			x += s.Location.X - x + s.Distance - aoc.Abs(y-s.Location.Y)
		}
	}

	return 0
}
