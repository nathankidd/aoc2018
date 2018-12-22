package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		fmt.Sscanf(scanner.Text(), "position=< %d,  %d> velocity=<%d, %d>", &l.pos.x, &l.pos.y, &l.vel.x, &l.vel.y)
		m.light = append(m.light, l)
	}
	return m
}

// Guess they are using the same font, 8 pixels high, check if all lights are within an 8 pixel vertial range
func (m *MsgSim) LooksLikeMsg(height int) bool {
	miny := math.MaxInt64
	maxy := math.MinInt64
	for _, l := range m.light {
		if l.pos.y < miny {
			miny = l.pos.y
		}
		if l.pos.y > maxy {
			maxy = l.pos.y
		}
	}
	fmt.Println("range: ", maxy-miny)
	if maxy-miny <= height {
		return true
	}
	return false
}

func (m *MsgSim) PrintLights() {
	fmt.Println("bingo")
	//for _, l := range m.light {
	//}
}

func (m *MsgSim) Run() {
	var reset []Light
	copy(reset, m.light)
	for r := 20; r > 5; r-- {
		copy(m.light, reset)
		fmt.Println(r)
		for j := 0; j < 20000; j++ {
			for i := 0; i < len(m.light); i++ {
				m.light[i].pos.x += m.light[i].vel.x
				m.light[i].pos.y += m.light[i].vel.y
			}
			if m.LooksLikeMsg(r) {
				m.PrintLights()
				break
			}
		}
	}
}

func main() {

	m := NewMsgSim("input")
	m.Run()
}
