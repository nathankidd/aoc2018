package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

const NumWorkers = 5

type Step struct {
	name     byte
	deps     []byte
	children []byte
}

type Worker struct {
	cur     byte
	donesec int
}

type Stepper struct {
	input     map[byte]Step
	completed map[byte]bool
	answer    []byte
	list      []byte
	workers   []Worker
	sec       int // Current second of simulation
}

func NewStepper(filename string) *Stepper {
	s := new(Stepper)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s.input = make(map[byte]Step)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name byte
		var dep byte
		// "Step G must be finished before step W can begin."
		_, err := fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.\n", &dep, &name)
		if err != nil {
			panic("bad input")
		}
		step := s.input[name]
		step.name = name
		step.deps = append(step.deps, dep)
		s.input[name] = step
	}
	s.completed = make(map[byte]bool)
	return s

}

func (s *Stepper) GetNextStep() (byte, error) {
	// Iterate through current list till reaching end
list:
	for i, c := range s.list {
		// Pick first letter with completed dependencies
		for _, dep := range s.input[c].deps {
			if !s.completed[dep] {
				continue list
			}
		}
		// Remove letter from list
		if i+1 > len(s.list) {
			s.list = s.list[:i]
		} else {
			s.list = append(s.list[:i], s.list[i+1:]...)
		}
		return c, nil
	}
	return 0, errors.New("No step available")
}

func In(c byte, list []byte) bool {
	for _, cc := range list {
		if c == cc {
			return true
		}
	}
	return false
}

func (s *Stepper) Complete(c byte) {
	s.completed[c] = true
	s.answer = append(s.answer, c)
	for _, cc := range s.input[c].children {
		if !In(cc, s.list) {
			s.list = append(s.list, cc)
		}
	}
}

func (s *Stepper) Process() {
	// Populate children
	// TODO consider going in reverse, then reversing the final string
	for k, v := range s.input {
		for _, d := range v.deps {
			step := s.input[d]
			step.children = append(step.children, k)
			s.input[d] = step
		}
	}

	// Find starting step, which has no dependencies
	for k, v := range s.input {
		if len(v.deps) == 0 {
			s.list = append(s.list, k)
		}
	}

	s.workers = make([]Worker, NumWorkers)
	s.sec = 0
	for len(s.list) > 0 {
		sort.Slice(s.list, func(i, j int) bool { return s.list[i] < s.list[j] })
		for _, w := range s.workers {
			if w.cur == 0 {
				var err error
				w.cur, err = s.GetNextStep()
				if err == nil {
					w.donesec = s.sec + int(w.cur-'A') + 61
				}
			}
			if w.donesec == s.sec {
				s.Complete(w.cur)
				w.cur = 0
			}
		}
		s.sec++
	}
}

func main() {
	s := NewStepper("input2")
	s.Process()
	fmt.Println(s.answer)
	fmt.Println(s.sec)
}
