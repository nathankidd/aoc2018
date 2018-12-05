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

func CountClaims(claims []Rect, x, y int) int {
	c := 0
	for _, r := range claims {
		if x >= r.x && x < r.x+r.width && y >= r.y && y < r.y+r.height {
			c++
		}
	}
	return c
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
	overlaps := 0
	for y := 0; y < fab.height; y++ {
		for x := 0; x < fab.width; x++ {
			if CountClaims(claims, x, y) >= 2 {
				overlaps++
			}

		}
	}

	fmt.Println(overlaps)
}
