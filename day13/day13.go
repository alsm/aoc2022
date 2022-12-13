package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	. "github.com/alsm/aoc2022/aoc/collections"
	"golang.org/x/exp/slices"
)

func main() {
	data, _ := os.ReadFile("day13.txt")
	pairs := strings.Split(string(data), "\n\n")
	input := Map(pairs, func(i string) [2][]any {
		a, b, _ := strings.Cut(i, "\n")
		var x, y []any
		json.Unmarshal([]byte(a), &x)
		json.Unmarshal([]byte(b), &y)
		return [2][]any{x, y}
	})

	fmt.Println(do1(input))
	fmt.Println(do2(input))
}

func ordered(a, b any) int {
	switch {
	case reflect.TypeOf(a).Kind() == reflect.Slice && reflect.TypeOf(b).Kind() == reflect.Slice:
		for j, x := range a.([]any) {
			if j < len(b.([]any)) {
				switch ordered(x, b.([]any)[j]) {
				case -1:
					return -1
				case 1:
					return 1
				}
			}
		}
		return ordered(float64(len(a.([]any))), float64(len(b.([]any))))
	case reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64:
		switch {
		case a.(float64) < b.(float64):
			return 1
		case a.(float64) > b.(float64):
			return -1
		default:
			return 0
		}
	case reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Slice:
		return ordered([]any{a}, b.([]any))
	case reflect.TypeOf(a).Kind() == reflect.Slice && reflect.TypeOf(b).Kind() == reflect.Float64:
		return ordered(a.([]any), []any{b})
	}

	return 1
}

func do1(in [][2][]any) int {
	return Sum(MapWithIndex(in, func(x int, i [2][]any) int {
		if ordered(i[0], i[1]) == 1 {
			return x + 1
		}
		return 0
	}))
}

func do2(in [][2][]any) int {
	var input [][]any
	x := []any{[]any{6.0}}
	y := []any{[]any{2.0}}
	for _, v := range in {
		input = append(input, v[0], v[1])
	}
	input = append(input, x, y)

	slices.SortFunc(input, func(a, b []any) bool {
		return ordered(a, b) == 1
	})

	xi, _ := slices.BinarySearchFunc(input, x, func(a, b []any) int {
		return -ordered(a, b)
	})
	yi, _ := slices.BinarySearchFunc(input, y, func(a, b []any) int {
		return -ordered(a, b)
	})

	return (xi + 1) * (yi + 1)
}
