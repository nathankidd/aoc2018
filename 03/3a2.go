package main

import "fmt"
import "os"
import "bufio"

type Rect struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var fab Rect
	var claims []Rect
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r Rect
		// #1271 @ 527,837: 27x17
		_, err = fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d\n", &r.id, &r.x, &r.y, &r.width, &r.height)
		if err != nil {
			panic("scanf error")
		}
		if r.x+r.width > fab.width {
			fab.width = r.x + r.width
		}
		if r.y+r.width > fab.height {
			fab.height = r.y + r.height
		}
		claims = append(claims, r)
	}
	var f [1000][1000]byte

	for _, r := range claims {
		for y := 0; y < r.height; y++ {
			for x := 0; x < r.width; x++ {
				f[r.x+x][r.y+y]++
			}
		}
	}
	overlaps := 0
	for _, r := range claims {
		for y := 0; y < r.height; y++ {
			for x := 0; x < r.width; x++ {
				v := f[r.x+x][r.y+y]
				if v > 1 {
					overlaps++
					f[r.x+x][r.y+y] = 0 // don't count again
				}

			}
		}
	}

	fmt.Println(overlaps)
}
