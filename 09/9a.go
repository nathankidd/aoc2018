package main

import (
	"fmt"
)

const Log = false

type Player struct {
	id    int
	score int
}

type Game struct {
	numplayers int
	nummarbles int
	cur        int   // Current index in the circle
	circle     []int // Positive index increases in clockwise direction
	players    []Player
}

func NewGame(numplayers, nummarbles int) *Game {
	g := new(Game)
	g.nummarbles = nummarbles
	g.players = make([]Player, numplayers)
	for i := 0; i < len(g.players); i++ {
		g.players[i].id = i
	}
	return g
}

// Positive offset is clockwise, negative is counter-clockwise
func (g *Game) RelativeIndex(offset int) int {
	if len(g.circle) == 0 {
		return 0
	}
	offset = g.cur + (offset % len(g.circle))
	if offset < 0 {
		return len(g.circle) + offset
	} else {
		return offset % len(g.circle)
	}
}

// Inserts before index AKA counter-clockwise from index
// Returns index of insert
func (g *Game) Insert(i, v int) int {
	g.circle = append(g.circle, 0) // Grow by one
	copy(g.circle[i+1:], g.circle[i:])
	g.circle[i] = v
	return i // New current position
}

// Returns value of deleted
func (g *Game) Delete(i int) int {
	v := g.circle[i]
	if len(g.circle) < 0 {
		return 0
	} else if i+1 == len(g.circle) {
		g.circle = g.circle[:i]
	} else {
		g.circle = append(g.circle[:i], g.circle[i+1:]...)
	}
	return v
}

func (g *Game) Play() {
	for marble := 0; marble <= g.nummarbles; marble++ {
		if marble == 0 {
			g.circle = append(g.circle, 0)
		} else if marble%23 == 0 {
			curplayer := marble % len(g.players)
			g.players[curplayer].score += marble
			g.cur = g.RelativeIndex(-7)
			g.players[curplayer].score += g.circle[g.cur]
			g.Delete(g.cur)

		} else {
			g.cur = g.Insert(g.RelativeIndex(2), marble)
		}
	}
}

func (g *Game) GetWinningScore() int {
	score := 0
	for _, p := range g.players {
		if p.score > score {
			score = p.score
		}
	}

	return score
}

func main() {
	g := NewGame(411, 71058)
	g.Play()
	fmt.Println(g.GetWinningScore())
}
