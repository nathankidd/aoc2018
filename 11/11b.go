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
	rackID := x + 10
	v := rackID * y
	v += g.serial
	v *= rackID
	v = (v - (v / 1000 * 1000)) / 100
	v -= 5
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

type Coord struct {
	x int
	y int
}

// Eliminate consistently negative rows/columns
func (g *Grid) FindPosEdges() (int, int, int, int) {

	starty := 0
loop_starty:
	for y := 0; y < GridSize-1; y += 2 {
		for x := 0; x < GridSize; x++ {
			if grid[x][y]+grid[x][y+1] > 0 {
				continue loop_starty
			}
		}
		starty++
	}
	endy := 0
loop_endy:
	for y := GridSize - 1; y >= 0; y -= 2 {
		for x := 0; x < GridSize; x++ {
			if grid[x][y]+grid[x][y-1] > 0 {
				continue loop_endy
			}
		}
		endy--
	}
	return starty, endy, startx, endx
}

func (g *Grid) FindHighestValueSquare() (Coord, int) {
	highc := Coord{-1, -1}
	highv := math.MinInt64
	highsz := 0
	starty, endy, startx, endx := g.FindPosEdges()

	// Find hightest-value square
	for size := GridSize; size > 0; size-- {
		if size > endy-starty || size > endx-startx {
			continue
		}
		fmt.Println(size)
		var offsets []Coord
		for szy := 0; szy < size; szy++ {
			for szx := 0; szx < size; szx++ {
				offsets = append(offsets, Coord{szx, szy})
			}
		}
		for y := starty; y < endy-size; y++ {
			for x := startx; x < endx-size; x++ {
				v := 0
				for _, c := range offsets {
					v += g.grid[x+c.x][y+c.y]
				}
				if v > highv {
					highv = v
					highc = Coord{x, y}
					highsz = size
				}
			}
		}
		fmt.Printf("size %d: high = %d\n", size, highv)
	}
	fmt.Println(highv)
	return highc, highsz
}

func main() {
	g := NewGrid(8772)
	fmt.Println(g.FindHighestValueSquare())
}
