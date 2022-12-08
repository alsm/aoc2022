package main

import (
	"fmt"
	"os"
	"strings"

	. "github.com/alsm/aoc2022/aoc/collections"
)

type Entry struct {
	Name     string
	FileSize int
	Parent   *Entry
	Subs     []*Entry
}

func (e *Entry) Size() int {
	return e.FileSize + Sum(Map(e.Subs, func(e *Entry) int {
		return e.Size()
	}))
}

func (e *Entry) SizeSubDirs() []int {
	var ret []int
	for _, s := range e.Subs {
		ret = append(ret, s.Size())
		ret = append(ret, s.SizeSubDirs()...)
	}

	return ret
}

func main() {
	data, _ := os.ReadFile("day7.txt")
	cmds := Map(strings.Split(string(data), "$ "), func(s string) []string {
		return strings.Split(s, "\n")
	})
	root := &Entry{
		Name: "root",
	}
	e := root
	for _, c := range cmds[2:] {
		switch c[0] {
		case "cd ..":
			e = e.Parent
		case "ls":
			for _, l := range c[1:] {
				var size int
				fmt.Sscanf(l, "%d ", &size)
				e.FileSize += size
			}
		default:
			var dir string
			fmt.Sscanf(c[0], "cd %s", &dir)
			ne := &Entry{Name: dir, Parent: e}
			e.Subs = append(e.Subs, ne)
			e = ne
		}
	}

	fmt.Println(do1(root))
	fmt.Println(do2(root))
}

func do1(root *Entry) int {
	return Sum(Select(root.SizeSubDirs(), func(i int) bool {
		return i <= 100000
	}))
}

func do2(root *Entry) int {
	return Min(Select(root.SizeSubDirs(), func(i int) bool {
		return i >= 30000000-(70000000-root.Size())
	}))
}