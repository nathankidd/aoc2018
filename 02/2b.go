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

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

outer:
	for _, a := range list {
		for _, b := range list {
			diff := 0
			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					diff++
				}
			}
			if diff == 1 {
				result := make([]byte, len(a))
				for i := 0; i < len(a); i++ {
					if a[i] == b[i] {
						result[i] = a[i]
					}
				}
				fmt.Println(string(result))
				break outer
			}
		}
	}

}
