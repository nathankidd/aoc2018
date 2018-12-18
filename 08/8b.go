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
	children []Node
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

func (l *License) AddNode() Node {
	var n Node
	n.numchild = l.GetNextInt()
	n.nummeta = l.GetNextInt()
	for i := 0; i < n.numchild; i++ {
		child := l.AddNode()
		n.children = append(n.children, child)
	}
	for i := 0; i < n.nummeta; i++ {
		n.meta = append(n.meta, l.GetNextInt())
	}
	l.nodes = append(l.nodes, n)
	return n
}

func (l *License) SumNode(n Node) int {
	v := 0
	if n.numchild == 0 {
		for i := 0; i < n.nummeta; i++ {
			v += n.meta[i]
		}
	} else {
		for i := 0; i < n.nummeta; i++ {
			ci := n.meta[i] - 1
			if ci >= 0 && ci < n.numchild {
				v += l.SumNode(n.children[ci])
			}
		}
	}
	return v
}

func (l *License) RootNodeValue() int {
	root := l.AddNode()
	return l.SumNode(root)
}

func main() {
	l := NewLicense("input")
	fmt.Println(l.RootNodeValue())
}
