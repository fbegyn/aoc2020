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
	round2 := make(chan int, len(starting))
	age2 := make(chan int)
	end2 := make(chan bool)

	go determineAge(len(starting), 2020, round, age, end)
	go determineAge(len(starting), 30000000, round2, age2, end2)

	for _, start := range starting {
		round <- start
		round2 <- start
	}

	go func() {
		turn := 1
		var cur int
		for stop := range end {
			cur = <-age
			turn++
			if turn <= len(starting) {
				continue
			}
			if stop {
				break
			}
			round <- cur
		}
		close(age)
		close(round)
		fmt.Printf("solution for part 1: %d\n", cur)
	}()

	turn2 := 1
	var cur int
	for _ = range end2 {
		cur = <-age2
		turn2++
		if turn2 <= len(starting) {
			continue
		}
		round2 <- cur
	}
	close(age2)
	close(round2)
	fmt.Printf("solution for part 2: %d\n", cur)
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
