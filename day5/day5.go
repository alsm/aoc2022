package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

func main() {
	var stacks []string
	var commands [][3]int
	data, _ := os.ReadFile("day5.txt")
	start := strings.Split(strings.Split(string(data), "\n\n")[0], "\n")
	for i := 1; i < len(start[0]); i += 4 {
		stacks = append(stacks, parseCol(start, i))
	}
	for _, c := range strings.Split(strings.Split(string(data), "\n\n")[1], "\n") {
		commands = append(commands, parseCommand(c))
	}

	fmt.Println(do1(slices.Clone(stacks), commands))
	fmt.Println(do2(slices.Clone(stacks), commands))
}

func parseCommand(s string) [3]int {
	var c [3]int
	fmt.Sscanf(s, "move %d from %d to %d", &c[0], &c[1], &c[2])

	return c
}

func parseCol(in []string, col int) string {
	var ret strings.Builder

	for i := 0; i < len(in); i++ {
		if l := in[i][col]; unicode.IsLetter(rune(l)) {
			ret.WriteByte(l)
		}
	}

	return ret.String()
}

func do1(s []string, commands [][3]int) string {
	var ret string

	for _, c := range commands {
		for i := 1; i <= c[0]; i++ {
			s[c[2]-1] = string(s[c[1]-1][0]) + s[c[2]-1]
			s[c[1]-1] = s[c[1]-1][1:]
		}
	}

	for _, i := range s {
		ret += string(i[0])
	}

	return ret
}

func do2(s []string, commands [][3]int) string {
	var ret string

	for _, c := range commands {
		tmpStack := s[c[1]-1][:c[0]]
		s[c[1]-1] = s[c[1]-1][c[0]:]
		s[c[2]-1] = tmpStack + s[c[2]-1]
	}

	for _, i := range s {
		ret += string(i[0])
	}

	return ret
}
