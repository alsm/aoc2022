package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/alsm/aoc2022/aoc/collections"
	"golang.org/x/exp/constraints"
)

func SliceFromFile[T any](file string, conv func(string) T) []T {
	var ret []T

	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", file, err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, conv(scanner.Text()))
	}

	return ret
}

func ReadInts(in string) int64 {
	n, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func MustReadInput(path string) []string {
	var ret []string
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", path, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		ret = append(ret, s.Text())
	}

	return ret
}

func InputToInts(input []string) []int {
	return collections.Map(input, func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("could not convert '%s' to int", s)
		}

		return i
	})
}

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Sign[T constraints.Integer | constraints.Float](x T) T {
	switch {
	case x < 0:
		return -T(1)
	case x > 0:
		return 1
	default:
		return 0
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(ints ...int) int {
	result := ints[0] * ints[1] / GCD(ints[0], ints[1])
	ints = ints[2:]

	for i := 0; i < len(ints); i++ {
		result = LCM(result, ints[i])
	}

	return result
}

type Point struct {
	X int64
	Y int64
}

func (p *Point) MDistance(b Point) int64 {
	return Abs(p.X-b.X) + Abs(p.Y-b.Y)
}

func (p *Point) Add(b Point) Point {
	return Point{
		X: p.X + b.X,
		Y: p.Y + b.Y,
	}
}

func (p *Point) Neighbour(b Point) bool {
	dx, dy := Abs(b.X-p.X), Abs(b.Y-p.Y)
	return dx <= 1 && dy <= 1
}

func IPow(base, exp int64) int64 {
	ret := int64(1)

	for {
		if exp&1 > 0 {
			ret *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return ret
}
