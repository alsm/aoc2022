package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	"github.com/alsm/aoc2022/aoc/grid"
	"golang.org/x/exp/slices"
)

var (
	chars = map[bool]string{true: "#", false: "."}
)

func main() {
	input := aoc.SliceFromFile("day10.txt", func(i string) [2]int {
		x, _ := strconv.Atoi(strings.TrimSpace(i[4:]))
		return [2]int{1 + aoc.Abs(aoc.Sign(x)), x}
	})

	fmt.Println(do1(slices.Clone(input)))
	fmt.Println(do2(slices.Clone(input)))
}

func do1(in [][2]int) int64 {
	sigStr := int64(0)
	x := int64(1)
	for c := int64(1); c <= 220; c++ {
		if c%40 == 20 {
			sigStr += c * x
		}
		in[0][0]--
		if in[0][0] == 0 {
			x += int64(in[0][1])
			in = in[1:]
		}
	}

	return sigStr
}

func do2(in [][2]int) string {
	g := grid.New[string](40, 6, nil)
	x := int64(1)
	for c := 1; c <= 240; c++ {
		pixel := int64((c - 1) % 40)
		g.SetState(pixel, int64((c-1)/40), chars[pixel >= x-1 && pixel <= x+1])
		in[0][0]--
		if in[0][0] == 0 {
			x += int64(in[0][1])
			in = in[1:]
		}
	}

	return g.StateString()
}
