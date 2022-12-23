package main

import (
	"fmt"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/search"
	"github.com/alsm/aoc2022/aoc/set"
	"golang.org/x/exp/maps"
)

func main() {
	valves := make(map[string]*Valve)
	Each(aoc.SliceFromFile("day16.txt", func(i string) string { return i }), func(i string) {
		v := Valve{Others: make(map[string]int)}
		var others string
		i = strings.ReplaceAll(i, "s", "")
		i = strings.ReplaceAll(i, ", ", ",")
		fmt.Sscanf(i, "Valve %s ha flow rate=%d; tunnel lead to valve %s", &v.Name, &v.Flow, &others)
		v.Links = strings.Split(others, ",")
		valves[v.Name] = &v
	})
	vMap := ValveMap{valves: valves}
	EachMap(valves, func(n string, v *Valve) {
		EachMap(valves, func(nn string, nv *Valve) {
			if n == nn {
				return
			}
			distance := len(search.BFS[string](vMap, n, nn)) - 1
			fmt.Printf("From %s to %s is %d steps\n", n, nn, distance)
			v.Others[nn] = distance
		})
	})

	fmt.Println(do1(valves))
}

type ValveMap struct {
	valves map[string]*Valve
}

func (v ValveMap) Neighbours(n string) []string {
	fmt.Println("Getting neighbours for", n)
	if n == "" {
		return nil
	}
	fmt.Println(v.valves[n].Links)
	return v.valves[n].Links
}

type Valve struct {
	Name   string
	Flow   int
	Links  []string
	Others map[string]int
}

func do1(valves map[string]*Valve) int {
	useful := Map(Select(maps.Values(valves), func(v *Valve) bool {
		return v.Flow > 0
	}), func(v *Valve) string {
		return v.Name
	})

	// var best int

	// for _, r := range Permutations(useful) {
	// 	var count int
	// 	next := valves["AA"]
	// 	for x := 30, x > 0; {

	// 	}
	// }

	return 0
}

func do1o(valves map[string]*Valve) int {
	var count int
	seen := make(set.Set[string])
	next := valves["AA"]
	for x := 30; x > 0; {
		seen.Add(next.Name)
		fmt.Println(next, x)
		count += x * next.Flow
		nv, _ := MaxMap(MapMap(SelectMap(valves, func(n string, v *Valve) bool {
			return !seen.Contains(n)
		}), func(n string, v *Valve) (string, int) {
			fmt.Println("  ", n, next.Others[n], v.Flow*(x-next.Others[n]-1))
			return n, v.Flow * (x - next.Others[n] - 1)
		}))
		fmt.Println("Going to", nv)
		if nv == "" {
			break
		}
		x -= next.Others[nv] + 1
		next = valves[nv]
	}

	return count
}
