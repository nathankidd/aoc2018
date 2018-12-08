package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

const Visual = false

type Mark struct {
	id   int
	dist int
}

type Point struct {
	x int
	y int
}

type Coord struct {
	id       int
	x        int
	y        int
	total    int  // Number of marks for this id
	expanded bool // Cannot expand area around anymore
	invalid  bool // Touches infinite area, thus invalid
}

func Abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -v
	}
}

const Collision = 255

func ClearScreen() {
	if !Visual {
		return
	}
	fmt.Printf("\x1b[2J")
}

func ResetColorPos() {
	if !Visual {
		return
	}
	fmt.Printf("\x1b[40;37m\x1b[0;0H")
}

func ResetColorLine() {
	if !Visual {
		return
	}
	fmt.Printf("\x1b[40;37m\x1b[900D\n")
}

func PrintMark(m Mark, x, y int) {
	if !Visual {
		return
	}
	x++
	y++
	if m.id == Collision {
		fmt.Printf("\x1b[%d;%dH\x1b[%d;40m%x", y, x, 30+m.id%8, m.dist%16)
	} else {
		fmt.Printf("\x1b[%d;%dH\x1b[%d;%dm%x", y, x, 30+m.id%8, 40+m.id%8, m.dist%16)
	}
}

func PrintCoord(id, x, y int) {
	if !Visual {
		return
	}
	x++
	y++
	fmt.Printf("\x1b[%d;%dH\x1b[%dm%c", y, x, 30+id%8, int('A')+id-1)
}

func BlinkAnswer(c Coord) {
	if !Visual {
		return
	}
	fmt.Printf("\x1b[%d;%dH\x1b[31;40;5m*\x1b[0m\x1b[%500;0H", c.y+1, c.x+1)
}

func ManhattanDistance(a, b Point) int {
	return Abs(b.x-a.x) + Abs(b.y-a.y)
}

// Marks a ring of radius dist around point
// Returns number of Marks placed
func MarkAroundCoord(field [][]Mark, c Coord, radius int) int {
	marked := 0

	sidelen := radius * 2
	sidenum := 0 // Which of the four sides are we on
	sidepos := 0 // Current position on current side

	// Start at (top-1)-left, go CCW, finish at top-left
	x := c.x - radius
	y := c.y - radius + 1
	dx := 0
	dy := 1
loop:
	for {
		if x >= 0 && x < len(field) && y >= 0 && y < len(field[0]) {
			d := ManhattanDistance(Point{c.x, c.y}, Point{x, y})
			if field[x][y].id == 0 || d < field[x][y].dist {
				field[x][y] = Mark{c.id, d}
				marked++
				PrintMark(field[x][y], x, y)
			} else if d == field[x][y].dist {
				field[x][y].id = Collision
				PrintMark(field[x][y], x, y)
			}
		}
		if Visual {
			time.Sleep(1 * time.Microsecond)
		}
		sidepos++
		if sidepos == sidelen {
			switch sidenum {
			case 0:
				dx = 1
				dy = 0
			case 1:
				dx = 0
				dy = -1
			case 2:
				dx = -1
				dy = 0
			case 3:
				break loop
			}
			sidenum++
			sidepos = 0
		}
		x += dx
		y += dy
	}
	return marked
}

func Parse(filename string) []Coord {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []Coord
	id := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var c Coord
		_, err := fmt.Sscanf(scanner.Text(), "%d, %d\n", &c.x, &c.y)
		if err != nil {
			panic("bad input")
		}
		c.id = id
		id++
		input = append(input, c)
	}
	return input
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
	//	fmt.Printf("%d,%d -> %d, %d = %d x %d\n", top, left, bottom, right, right-left, bottom-top)

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

	var f = make([][]Mark, right+1)
	for i := range f {
		f[i] = make([]Mark, bottom+1)
	}

	ResetColorPos()
	ClearScreen()

	// Set initial coords (easier than making MarkAroundCood special-case 0-radius)
	for _, c := range input {
		f[c.x][c.y].id = c.id
		PrintCoord(c.id, c.x, c.y)
	}

	// Draw rings around each coordinate, till nothing further can be marked
	someunexpanded := true
	for radius := 1; someunexpanded; radius++ {
		someunexpanded = false
		for _, c := range input {
			if c.expanded {
				continue
			}
			if MarkAroundCoord(f, c, radius) == 0 {
				c.expanded = true
			} else {
				someunexpanded = true
			}
		}
	}
	ResetColorLine()

	// Count Marks per Coord
	for y := 0; y < len(f[0]); y++ {
		for x := 0; x < len(f); x++ {
			if f[x][y].id == Collision {
				continue
			}
			i := f[x][y].id - 1
			// Mark invalid anything that touches an edge
			if x == 0 || y == 0 || x == len(f)-1 || y == len(f[0])-1 {
				input[i].invalid = true
			}
			input[i].total++
		}
	}

	// Find coordinate with biggest area
	var biggest Coord
	for _, c := range input {
		if c.id == Collision || c.invalid {
			continue
		}
		if c.total > biggest.total {
			biggest = c
		}
	}
	BlinkAnswer(biggest)

	return biggest.total
}

func main() {
	in := Parse("input")
	result := Answer(in)
	if !Visual {
		fmt.Println(result)
	}
}
