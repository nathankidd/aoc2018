package main

import "fmt"
import "os"
import "bufio"
import "math"

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

func ReactLen(in []byte) int {
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
	return num
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

	shortestlen := math.MaxInt64
	for c := byte('A'); c <= byte('Z'); c++ {
		tmp := make([]byte, len(in))
		copy(tmp, in)
		for i, cc := range tmp {
			if cc|32 == c|32 {
				tmp[i] = 0
			}
		}
		l := ReactLen(tmp)
		if l < shortestlen {
			shortestlen = l
		}
		fmt.Printf("%c : %d\n", c, l)
	}
	fmt.Println(shortestlen)
}
