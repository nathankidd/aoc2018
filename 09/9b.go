package main

import (
	"fmt"
)

const Log = false

type Marble struct {
	v    int
	next *Marble
	prev *Marble
}

type Player struct {
	id    int
	score int
}

type Game struct {
	numplayers int
	nummarbles int
	head       *Marble // .next moves in clockwise direction, .prev is CCW
	cur        *Marble // Current position in the circle
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

// Positive steps are clockwise, negative are counter-clockwise
func (g *Game) Walk(steps int) *Marble {
	m := g.cur
	if steps < 0 {
		for i := 0; i < -steps; i++ {
			m = m.prev
		}
	} else {
		for i := 0; i < steps; i++ {
			m = m.next
		}
	}
	return m
}

// Returns inserted marble
func (g *Game) InsertAfter(after *Marble, v int) *Marble {
	m := Marble{v, after.next, after}
	after.next.prev = &m
	after.next = &m
	if m.next == g.head {
		g.head.prev = &m
	}
	return &m
}

// Returns value of deleted
func (g *Game) Delete(m *Marble) int {
	if g.head == m {
		g.head = m.next
	}
	m.next.prev = m.prev
	m.prev.next = m.next
	return m.v
}

func (g *Game) Play() {
	for marble := 0; marble <= g.nummarbles; marble++ {
		if marble == 0 {
			g.head = new(Marble)
			g.head.v = marble
			g.head.next = g.head
			g.head.prev = g.head
			g.cur = g.head
		} else if marble%23 == 0 {
			player := &g.players[marble % len(g.players)]
			player.score += marble
			g.cur = g.Walk(-6)
			player.score += g.Delete(g.cur.prev)

		} else {
			g.cur = g.InsertAfter(g.Walk(1), marble)
		}
		if Log {
			fmt.Printf("[%d] ", marble)
			for m := g.head; ; m = m.next {
				if m == g.cur {
					fmt.Printf("(%2d)", m.v)
				} else {
					fmt.Printf(" %2d ", m.v)
				}
				if m.next == g.head {
					break
				}
			}
			fmt.Println()

			fmt.Printf("-%d- ", marble)
			for m := g.head.prev; ; m = m.prev {
				if m == g.cur {
					fmt.Printf("(%2d)", m.v)
				} else {
					fmt.Printf(" %2d ", m.v)
				}
				if m == g.head {
					break
				}
			}
			fmt.Println()
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
	g := NewGame(411, 71058*100)
	g.Play()
	fmt.Println(g.GetWinningScore())
}

