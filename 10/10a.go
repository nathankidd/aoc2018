package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Pair struct {
	x int
	y int
}

type Light struct {
	pos Pair
	vel Pair
}

type MsgSim struct {
	light []Light
}

func NewMsgSim(filename string) *MsgSim {
	m := new(MsgSim)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l Light
		// "position=< 41710,  52012> velocity=<-4, -5>"
		_, err := fmt.Sscanf(scanner.Text(), "position=<%d, %d> velocity=<%d, %d>", &l.pos.x, &l.pos.y, &l.vel.x, &l.vel.y)
		if err != nil {
			panic(err)
		}
		m.light = append(m.light, l)
	}
	return m
}

const AdjacentHeuristic = 9 // How many adjacent
const SpacingHeuristic = 1  // How many pixels away counts as adjacent

func (m *MsgSim) LooksLikeMsg() bool {
	x := make(map[int][]int)
	// Check for vertically-adjacent pixels
	for _, l := range m.light {
		x[l.pos.x] = append(x[l.pos.x], l.pos.y)
	}
	for k, _ := range x {
		l := x[k]
		if len(l) < AdjacentHeuristic {
			continue
		}
		sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
		adjacent := 0
		for i := 1; i < len(l); i++ {
			if l[i]-l[i-1] <= SpacingHeuristic {
				adjacent++
				if adjacent >= AdjacentHeuristic {
					return true
				}
			} else {
				adjacent = 0
			}
		}
	}
	return false
}

func (m *MsgSim) PrintLights() {
	fmt.Printf("\x1b[2J")
	minx := math.MaxInt64
	miny := math.MaxInt64
	for _, l := range m.light {
		if l.pos.x < minx {
			minx = l.pos.x
		}
		if l.pos.y < miny {
			miny = l.pos.y
		}
	}
	for _, l := range m.light {
		l.pos.x -= minx
		l.pos.y -= miny
		if l.pos.y < 300 && l.pos.x < 300 {
			fmt.Printf("\x1b[%d;%dH\x1b[31m*", l.pos.y, l.pos.x)
		}
	}
	fmt.Printf("\x1b[%d;%dH\x1b[31m*", 20, 0)
}

func (m *MsgSim) Run() {
	iter := 0
	for {
		iter++
		if iter%10000 == 0 {
			fmt.Println(iter)
		}
		if m.LooksLikeMsg() {
			m.PrintLights()
			fmt.Printf("iter=%d\n", iter-1)
			break
		}
		for i := 0; i < len(m.light); i++ {
			m.light[i].pos.x += m.light[i].vel.x
			m.light[i].pos.y += m.light[i].vel.y
		}
	}
}

func main() {
	m := NewMsgSim("input")
	m.Run()
}
