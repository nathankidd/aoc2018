package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	numchild int
	nummeta  int
	meta     []int
}

type License struct {
	input []int
	i     int // Index of next unused data in input
	nodes []Node
}

func NewLicense(filename string) *License {
	l := new(License)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), " ") {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic("can't grok input")
			}
			l.input = append(l.input, i)
		}
	}
	return l

}

func (l *License) GetNextInt() int {
	i := l.input[l.i]
	l.i++
	return i
}

func (l *License) AddNode() {
	var n Node
	n.numchild = l.GetNextInt()
	n.nummeta = l.GetNextInt()
	for i := 0; i < n.numchild; i++ {
		l.AddNode()
	}
	for i := 0; i < n.nummeta; i++ {
		n.meta = append(n.meta, l.GetNextInt())
	}
	l.nodes = append(l.nodes, n)
}

func (l *License) CountMeta() int {
	l.AddNode()

	meta := 0
	for _, n := range l.nodes {
		for _, m := range n.meta {
			meta += m
		}
	}

	return meta
}

func main() {
	l := NewLicense("input")
	fmt.Println(l.CountMeta())
}
