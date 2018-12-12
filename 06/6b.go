package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const DistLimitSum = 10000

func Parse(filename string) []Coord {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []Coord
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c Coord
		_, err := fmt.Sscanf(scanner.Text(), "%d, %d\n", &c.x, &c.y)
		if err != nil {
			panic("bad input")
		}
		input = append(input, c)
	}
	return input
}

type Coord struct {
	x int
	y int
}

func Abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}

func ManhattanDistance(a, b Coord) int {
	return Abs(b.x-a.x) + Abs(b.y-a.y)
}

func SumDistanceFromPointToAllCoords(input []Coord, x, y int) int {
	sum := 0
	for _, c := range input {
		sum += ManhattanDistance(c, Coord{x, y})
	}
	return sum
}

func Answer(input []Coord) int {
	top := math.MaxInt64
	left := math.MaxInt64
	bottom := math.MinInt64
	right := math.MinInt64

	for _, c := range input {
		if c.x < left {
			left = c.x
		}
		if c.y < top {
			top = c.y
		}
		if c.x > right {
			right = c.x
		}
		if c.y > bottom {
			bottom = c.y
		}
	}
	// TODO do we need to expand border of search coordinates?

	count := 0
	for y := 0; y < bottom; y++ {
		for x := 0; x < right; x++ {
			if SumDistanceFromPointToAllCoords(input, x, y) < DistLimitSum {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(Answer(Parse("input")))
}
