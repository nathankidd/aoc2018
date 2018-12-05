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

func HasOverlappingClaims(claims []Rect, c Rect) bool {

	for _, r := range claims {
		if r.id == c.id {
			continue
		}
		if (c.x < r.x+r.width && c.x+c.width >= r.x) && (c.y < r.y+r.height && c.y+c.height >= r.y) {
			return true
		}
	}
	return false
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
	for _, r := range claims {
		if !HasOverlappingClaims(claims, r) {
			fmt.Println(r.id)
			break
		}

	}

}
