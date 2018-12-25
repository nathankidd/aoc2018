package main

import (
	"fmt"
	"math"
)

const Log = false
const GridSize = 300

type Grid struct {
	grid   [GridSize][GridSize]int
	serial int
}

func (g *Grid) PowerLevel(x, y int) int {
	//fmt.Printf("(%d, %d)\n", x, y)
	rackID := x + 10
	//fmt.Printf(" rackID : %d\n", rackID)
	v := rackID * y
	//fmt.Printf(" v := rackID * y : %d\n", v)
	v += g.serial
	//fmt.Printf(" v += serial : %d\n", v)
	v *= rackID
	//fmt.Printf(" v *= rackID : %d\n", v)
	v = (v - (v / 1000 * 1000)) / 100
	//fmt.Printf(" v 100s : %d\n", v)
	v -= 5
	//fmt.Printf(" v -=5 : %d\n", v)
	return v
}

func NewGrid(serial int) *Grid {
	g := new(Grid)
	g.serial = serial
	for y := 0; y < GridSize; y++ {
		for x := 0; x < GridSize; x++ {
			g.grid[x][y] = g.PowerLevel(x, y)
		}
	}
	return g
}

const SquareSize = 3

type Coord struct {
	x int
	y int
}

func (g *Grid) FindHighestValueSquare() Coord {
	offsets := []Coord{
		{0, 0}, {1, 0}, {2, 0},
		{0, 1}, {1, 1}, {2, 1},
		{0, 2}, {1, 2}, {2, 2},
	}
	highc := Coord{-1, -1}
	highv := math.MinInt64
	for y := 0; y < GridSize-SquareSize; y++ {
		for x := 0; x < GridSize-SquareSize; x++ {
			v := 0
			for _, c := range offsets {
				v += g.grid[x+c.x][y+c.y]
			}
			if v > highv {
				highv = v
				highc = Coord{x, y}
			}
		}
	}
	fmt.Println(highv)
	return highc
}

func main() {
	g := NewGrid(8772)
	fmt.Println(g.FindHighestValueSquare())
}
