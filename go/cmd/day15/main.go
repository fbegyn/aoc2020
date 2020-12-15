package main

import (
	"os"

	"fmt"
	"strconv"
	"strings"

	_ "github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	input := os.Args[1]
	startingStr := strings.Split(input, ",")
	starting := []int{}
	for _, number := range startingStr {
		conv, _ := strconv.Atoi(number)
		starting = append(starting, conv)
	}

	round := make(chan int, len(starting))
	age := make(chan int)
	end := make(chan bool)

	//go determineAge(len(starting), 2020, round, age, end)
	go determineAge(len(starting), 30000000, round, age, end)

	for _, start := range starting {
		round <- start
	}

	turn := 1
	var cur int
	for _ = range end {
		cur = <-age
		turn++
		if turn <= len(starting) {
			continue
		}
		round <- cur
	}
	fmt.Println(cur)
}

func determineAge(startingSize, stopping int, inp <-chan int, spoken chan<- int, end chan<- bool) {
	last := make(map[int][2]int)
	round := 1
	for in := range inp {
		if stopping <= round {
			end <- true
			close(end)
		} else {
			end <- false
		}

		if round <= startingSize {
			last[in] = [2]int{round, 0}
			spoken <- in
			round++
			continue
		}

		var answer int
		l, _ := last[in]
		prev := l[0]
		before := l[1]
		if before == 0 {
			answer = 0
		} else {
			answer = prev - before
		}
		prevAns := last[answer][0]
		last[answer] = [2]int{round, prevAns}
		spoken <- answer

		round++
	}
}
