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

	freq := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		freq += i
	}

	fmt.Println(freq)
}
