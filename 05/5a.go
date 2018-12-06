package main

import "fmt"
import "os"
import "bufio"

func React(in []byte) int {
	reactions := 0

	var previ int
	for previ := 0; previ < len(in) && in[previ] == 0; previ++ {
	}
	for i := previ + 1; i < len(in); i++ {
		if in[i] == 0 {
			continue
		}
		if in[previ] != in[i] && in[previ]|32 == in[i]|32 {
			reactions++
			in[i] = 0
			in[previ] = 0
		}
		previ = i

	}
	return reactions
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var in []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in = []byte(scanner.Text())
	}
	var reacts int
	for {
		reacts = React(in)
		if reacts == 0 {
			break
		}
	}

	num := 0
	for _, i := range in {
		if i != 0 {
			num++
		}
	}
	fmt.Println(num)
}
