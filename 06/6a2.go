package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Coord struct {
	x int
	y int
}

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

// Returns index of closest coord, or -1 if collision
func GetClosestCoord(c Coord, input []Coord) int {
	min := math.MaxInt64
	var closest []int
	for i, c2 := range input {
		d := ManhattanDistance(c, c2)
		if d < min {
			min = d
			closest = []int{i}
		} else if d == min {
			closest = append(closest, i)
		}
	}
	if len(closest) > 1 {
		return -1
	}
	return closest[0]
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

	// Translate to 0-based coordinates, for smallest field
	xoff := left
	yoff := top
	for i := 0; i < len(input); i++ {
		input[i].x -= xoff
		input[i].y -= yoff
	}
	left -= xoff
	right -= xoff
	top -= yoff
	bottom -= yoff

	totals := make([]int, len(input))
	// Count locations closest to each coord
	for y := 0; y < bottom; y++ {
		for x := 0; x < right; x++ {
			id := GetClosestCoord(Coord{x, y}, input)
			if id != -1 {
				totals[id]++
			}
		}
	}

	// Invalidate any coordinate that is closer to an edge than any other coord; it will go infinite
	for i, c := range input {
		if i == GetClosestCoord(Coord{c.y, 0}, input) ||
			i == GetClosestCoord(Coord{c.y, right}, input) ||
			i == GetClosestCoord(Coord{0, c.x}, input) ||
			i == GetClosestCoord(Coord{c.x, bottom}, input) {
			totals[i] = 0
		}
	}

	// Find coordinate with biggest area
	var biggest int
	for _, t := range totals {
		if t > biggest {
			biggest = t
		}
	}

	return biggest
}

func main() {
	fmt.Println(Answer(Parse("input")))
}
