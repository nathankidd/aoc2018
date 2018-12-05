package main

import "fmt"
import "os"
import "bufio"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	twos := 0
	threes := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := make(map[rune]int)
		for _, chr := range scanner.Text() {
			m[chr]++
		}
		havetwo := false
		havethree := false
		for _, value := range m {
			if value == 2 {
				havetwo = true
			} else if value == 3 {
				havethree = true
			}
		}
		if havetwo {
			twos++
		}
		if havethree {
			threes++
		}
	}

	fmt.Println(twos * threes)
}
