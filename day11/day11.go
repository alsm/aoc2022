package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alsm/aoc2022/aoc"
	. "github.com/alsm/aoc2022/aoc/collections"
	"github.com/alsm/aoc2022/aoc/list"
)

var lcm int

func main() {
	var factors []int
	data, _ := os.ReadFile("day11.txt")
	input := strings.Split(string(data), "\n\n")
	monkeys := make([]*Monkey, len(input))
	for _, m := range input {
		var items, op, opd string
		var mn, test, throwT, throwF int
		m = strings.Replace(strings.Join(strings.Split(m, "\n"), " "), ", ", ",", -1)
		fmt.Sscanf(m, `Monkey %d:   Starting items: %s  Operation: new = old %s %s   Test: divisible by %d     If true: throw to monkey %d     If false: throw to monkey %d`, &mn, &items, &op, &opd, &test, &throwT, &throwF)
		monkeys[mn] = NewMonkey(monkeys, Map(strings.Split(items, ","), func(i string) int {
			v, _ := strconv.Atoi(i)
			return v
		}))
		monkeys[mn].Test = func(i int) int {
			if i%test == 0 {
				return throwT
			}
			return throwF
		}
		monkeys[mn].Operation = func(a int) int {
			b, _ := strconv.Atoi(opd)
			if b == 0 {
				b = a
			}
			if op == "*" {
				return a * b
			}
			return a + b
		}
		factors = append(factors, test)
	}

	lcm = aoc.LCM(factors...)

	fmt.Println(do1(monkeys))
	fmt.Println(do2(monkeys))
}

type Monkey struct {
	Inspections  int
	initialItems []int
	Items        *list.List[int]
	Operation    func(int) int
	Test         func(int) int
	Others       []*Monkey
}

func (m *Monkey) Catch(i int) {
	m.Items.PushBack(i)
}

func (m *Monkey) InspectAndThrow(worryDiv int) {
	for m.Items.Len() != 0 {
		i := m.Items.PopFront()
		i = m.Operation(i) / worryDiv
		if i > int(lcm) {
			i = i % int(lcm)
		}
		m.Others[m.Test(i)].Catch(i)
		m.Inspections++
	}
}

func NewMonkey(others []*Monkey, items []int) *Monkey {
	return &Monkey{
		initialItems: items,
		Items:        list.New[int]().InitFromSlice(items),
		Others:       others,
	}
}

func (m *Monkey) Reset() {
	m.Items = list.New[int]().InitFromSlice(m.initialItems)
	m.Inspections = 0
}

func do1(in []*Monkey) int {
	for i := 0; i < 20; i++ {
		for _, m := range in {
			m.InspectAndThrow(3)
		}
	}

	return Product(MaxN(Map(in, func(m *Monkey) int {
		return m.Inspections
	}), 2))
}

func do2(in []*Monkey) int {
	for _, m := range in {
		m.Reset()
	}
	for i := 0; i < 10000; i++ {
		for _, m := range in {
			m.InspectAndThrow(1)
		}
	}

	return Product(MaxN(Map(in, func(m *Monkey) int {
		return m.Inspections
	}), 2))
}
