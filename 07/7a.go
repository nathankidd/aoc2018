package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Step struct {
	name     byte
	deps     []byte
	children []byte
}

func Parse(filename string) map[byte]Step {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := make(map[byte]Step)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name byte
		var dep byte

		// "Step G must be finished before step W can begin."
		_, err := fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.\n", &dep, &name)
		if err != nil {
			panic("bad input")
		}
		step := input[name]
		step.name = name
		step.deps = append(step.deps, dep)
		input[name] = step
	}
	return input
}

func In(c byte, list []byte) bool {
	for _, cc := range list {
		if c == cc {
			return true
		}
	}
	return false
}

func Answer(input map[byte]Step) string {
	var answer []byte

	// Populate children
	// TODO consider going in reverse, then reversing the final string
	for k, v := range input {
		for _, d := range v.deps {
			s := input[d]
			s.children = append(s.children, k)
			input[d] = s
		}
	}

	// Find starting step, which has no dependencies
	var list []byte
	for k, v := range input {
		if len(v.deps) == 0 {
			list = append(list, k)
		}
	}

	completed := make(map[byte]bool)
	// Iterate through current list till reaching end
loop:
	for len(list) > 0 {
		sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })
	list:
		for i, c := range list {
			// Pick first letter with completed dependencies
			for _, dep := range input[c].deps {
				if !completed[dep] {
					continue list
				}
			}
			answer = append(answer, c)
			completed[c] = true
			// replace letter with its children, that aren't already in list
			if i+1 > len(list) {
				list = list[:i]
			} else {
				list = append(list[:i], list[i+1:]...)
			}
			for _, cc := range input[c].children {
				if !In(cc, list) {
					list = append(list, cc)
				}
			}
			continue loop
		}
	}

	return string(answer)
}

func main() {
	fmt.Println(Answer(Parse("input")))
}
