package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/alsm/aoc2022/aoc/collections"
)

func main() {
	d, _ := os.ReadFile("day1.txt")
	input := Map(strings.Split(string(d), "\n\n"), func(s string) []int {
		return Map(strings.Split(s, "\n"), func(n string) int {
			v, _ := strconv.Atoi(n)
			return v
		})
	})

	fmt.Println(do1(input))
	fmt.Println(do2(input))

	objectsToDelete := make([]int, 5432)

	for i := 0; i <= len(objectsToDelete)/1000; i++ {
		fmt.Println(i)
		chunk := 1000
		if remaining := len(objectsToDelete) - (i * 1000); remaining < 1000 {
			chunk = remaining
		}
		fmt.Println(len(objectsToDelete[i*1000 : i*1000+chunk]))
		fmt.Println(i*1000, i*1000+chunk)
	}

}

func do1(i [][]int) int {
	return Max(Map(i, Sum[int]))
}

func do2(i [][]int) int {
	return Sum(MaxN(Map(i, Sum[int]), 3))
}
