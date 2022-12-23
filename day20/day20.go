package main

import (
	"fmt"
	"strconv"

	"github.com/alsm/aoc2022/aoc"
	"github.com/alsm/aoc2022/aoc/ring"
	"github.com/gammazero/deque"
	"golang.org/x/exp/slices"
)

func main() {
	input := aoc.SliceFromFile("day20.txt", func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	})
	mix := deque.New[int]()
	for _, v := range input {
		mix.PushBack(v)
	}

	// fmt.Println(pring(mix))
	// fmt.Println(pring(mix))

	fmt.Println(do1(mix, input))
}

func pring(r *ring.Ring[int]) string {
	var s string
	r.Do(func(x int) {
		s += fmt.Sprintf("%d,", x)
	})
	return s
}

func do1(mix deque.Deque[int], codes []int) int {
	for _, v := range codes {
		j := slices.Index(mix, v)
		slices.Delete(mix, j, j+1)
		nl := (j + v)
		if nl <= 0 {
			nl = (nl + len(mix)*10 - 1) % len(mix)
		}
		if nl >= len(mix) {
			nl = (nl + len(mix)*10 + 1) % len(mix)
		}
		slices.Insert(mix, nl, v)
		fmt.Println(v, nl)
		// fmt.Println("End: moved", v, pring(mix))
	}
	fmt.Println(mix)
	zero := slices.Index(mix, 0)
	// zero.Do(func(r int) {
	// 	fmt.Println(r)
	// })
	first := mix[(zero+1000)%len(mix)]
	second := mix[(zero+2000)%len(mix)]
	third := mix[(zero+3000)%len(mix)]
	fmt.Println(first, second, third)

	return first + second + third
}
