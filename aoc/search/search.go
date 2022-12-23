package search

import (
	// "log"

	"github.com/alsm/aoc2022/aoc/queue"
	"github.com/alsm/aoc2022/aoc/routing"
)

func BFS[T comparable](g routing.Graph[T], start T, goal T) []T {
	var frontier queue.Queue[T]
	frontier.Put(start)

	cameFrom := make(map[T]*T)
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

	ret := []T{goal}
	for n := cameFrom[goal]; n != nil; n = cameFrom[*n] {
		ret = append(ret, *n)
	}

	return ret
}

// func AllPaths[T comparable](adjList map[T][]T, start T, end func(T) bool) [][]T {
// 	var paths [][]T
// 	dfs(adjList, start, end, []T{}, &paths)

// 	return paths
// }

// func dfs[T comparable](adjList map[T][]T, start T, end func(T) bool, path []T, paths *[][]T) {
// 	if end(start) || slices.Contains(path, start) {
// 		*paths = append(*paths, path)
// 		return
// 	}

// 	for _, n := range adjList[start] {
// 		path = append(path, n)
// 		dfs(adjList, n, end, path, paths)
// 		path = path[:len(path)-1]
// 	}
// }

// func DoFS[T comparable](adjList map[T][]T, start T, end func(T) bool) [][]T {

// 	dfs := func(path []T, used map[T]bool, result *[][]int) {
// 		if len(perm) == len(nums) {
// 			*result = append(*result, append([]int{}, perm...))
// 			return
// 		}

// 		for _, n := range nums {
// 			if used[n] {
// 				continue
// 			}

// 			perm = append(perm, n)
// 			used[n] = true
// 			dfs(perm, used, result) // recursive dfs call

// 			// backtrack/go back
// 			perm = perm[:len(perm)-1] // remove last number
// 			used[n] = false           // mark as not used
// 		}

// 	}

// 	result := [][]int{}
// 	dfs([]int{}, map[int]bool{}, &result)
// 	return result
// }

// func defs[T comparable](node T, end func(T) bool, adjList map[T][]T, paths *[][]T, visited *map[T]int) bool {
// 	neighbours, ok := adjList[node]
// 	if !ok {
// 		return true
// 	}

// 	if (*visited)[node] == -1 {
// 		return false
// 	}

// 	if (*visited)[node] == 1 {
// 		return true
// 	}

// 	(*visited)[node] = -1

// 	for _, neighbour := range neighbours {
// 		if !dfs(neighbour, adjList, paths, visited) {
// 			return false
// 		}
// 	}
// 	(*visited)[node] = 1
// 	return true
// }
