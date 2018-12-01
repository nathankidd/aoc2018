package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ilist []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		ilist = append(ilist, i)
	}
	freqs := make(map[int]int)
	freq := 0
outer:
	for iter := 0; ; iter++ {
		for _, i := range ilist {
			if _, exist := freqs[freq]; exist {
				break outer
			}
			freqs[freq] = 1
			freq += i
		}
	}

	fmt.Println(freq)
}
