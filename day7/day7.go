package main

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/alsm/aoc2022/aoc"
)

func main() {
	lines := aoc.SliceFromFile("day7.txt", func(i string) string {
		return i
	})
	fmt.Println("DAY 7")
	fmt.Println("Part 1:", PartOneSolution(lines))
	fmt.Println("Part 2:", PartTwoSolution(lines))
}

func parseDirectories(lines []string) map[string]int {

	directorySizes := make(map[string]int)

	cwd := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			dir := strings.Trim(line[5:], " ")
			cwd = path.Join(cwd, dir)
		} else if strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir") {
			continue
		} else {
			file := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(file[0])
			directorySizes[cwd] += fileSize
		}
	}

	return directorySizes
}

func PartOneSolution(lines []string) uint64 {
	directorySizes := parseDirectories(lines)

	for n, v := range directorySizes {
		fmt.Println(n, v)
	}

	totalMap := make(map[string]int)

	for key, value := range directorySizes {
		total := value
		for innerKey, innerValue := range directorySizes {
			if key != innerKey && strings.HasPrefix(innerKey, key) {
				total += innerValue
			}
		}
		totalMap[key] = total
	}

	var finalTotal uint64
	for _, v := range totalMap {
		if v < 100_000 {
			finalTotal += uint64(v)
		}
	}

	return finalTotal
}

func PartTwoSolution(lines []string) int {
	return -1
}
