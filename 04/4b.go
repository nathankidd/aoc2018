package main

import "fmt"
import "os"
import "bufio"
import "sort"

type Guard struct {
	id    int
	total int     // minutes
	freq  [60]int // only care about one-hour period
}

func MostSleptMinute(guard Guard) (int, int) {
	min := -1
	freq := -1
	for m, f := range guard.freq {
		if f > freq {
			freq = f
			min = m
		}
	}
	return min, freq
}

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
	sort.Strings(list)

	guards := make(map[int]Guard)
	curguard := -1
	sleepstart := -1
	for _, s := range list {
		// [1518-05-16 23:57] Guard #1619 begins shift
		// [1518-08-11 00:57] wakes up
		// [1518-06-01 00:51] falls asleep
		var date string
		var hour int
		var min int
		_, err = fmt.Sscanf(s, "[%s %d:%d] Guard #%d begins shift\n", &date, &hour, &min, &curguard)
		if err == nil {
			continue
		}
		_, err = fmt.Sscanf(s, "[%s %d:%d] wakes up\n", &date, &hour, &min)
		if err == nil {
			g := guards[curguard]
			if sleepstart < 0 {
				panic("no sleep start!")
			}
			g.total += min - sleepstart
			for i := sleepstart; i < min; i++ {
				g.freq[i]++
			}
			g.id = curguard
			guards[curguard] = g
			continue
		}
		_, err = fmt.Sscanf(s, "[%s %d:%d] falls asleep\n", &date, &hour, &min)
		if err == nil {
			sleepstart = min
			continue
		}
		if err != nil {
			panic("scanf error")
		}
	}

	mostid := -1
	mostmin := -1
	mostfreq := -1
	for id, guard := range guards {
		min, freq := MostSleptMinute(guard)
		if freq > mostfreq {
			mostid = id
			mostmin = min
			mostfreq = freq
		}

	}

	fmt.Println(mostid * mostmin)
}
