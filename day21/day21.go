package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
)

func main() {
	numbers, instructions := Partition(aoc.SliceFromFile("day21.txt", func(s string) string { return s }),
		func(s string) bool {
			_, num, _ := strings.Cut(s, ": ")
			_, err := strconv.Atoi(num)
			return err == nil
		})

	monkeyNums := make(map[string]int)
	Each(numbers, func(s string) {
		name, num, _ := strings.Cut(s, ": ")
		val, _ := strconv.Atoi(num)
		monkeyNums[name] = val
	})
	monkeySums := make(map[string]string)
	Each(instructions, func(s string) {
		name, sum, _ := strings.Cut(s, ": ")
		monkeySums[name] = sum
	})

	fmt.Println(do1(monkeyNums, monkeySums))
	fmt.Println(do2(monkeyNums, monkeySums))
}

func monkeyVal(monkey string, n map[string]int, s map[string]string) int {
	if v, ok := n[monkey]; ok {
		return v
	}

	sum := s[monkey]
	var l, op, r string
	var val int
	fmt.Sscanf(sum, "%s %s %s", &l, &op, &r)
	lnum := monkeyVal(l, n, s)
	rnum := monkeyVal(r, n, s)
	switch op {
	case "+":
		val = lnum + rnum
	case "*":
		val = lnum * rnum
	case "-":
		val = lnum - rnum
	case "/":
		val = lnum / rnum
	}
	return val
}

func do1(n map[string]int, s map[string]string) int {
	return monkeyVal("root", n, s)
}

func do2(n map[string]int, s map[string]string) int {
	x, _ := sort.Find(1e18, func(i int) int {
		n["humn"] = i
		return monkeyVal("fflg", n, s) - monkeyVal("qwqj", n, s)
	})

	return x
}
