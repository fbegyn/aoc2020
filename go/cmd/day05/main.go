package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2020/go/helpers"
	"bufio"
	"sort"
)

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	seats := []int{}
	
	scanner := bufio.NewScanner(input)
		rowStream := make(chan int, 2)
		colStream := make(chan int, 2)

	for scanner.Scan() {
		line := scanner.Text()
		rows := line[:7]
		cols := line[7:]

		go helpers.BinaryParse(rows, 0, 127, rowStream)
		go helpers.BinaryParse(cols, 0, 7, colStream)

		seatID := <-rowStream * 8 + <-colStream
		seats = append(seats, seatID)
	}

	close(rowStream)
	close(colStream)

	sort.Ints(seats)
	mySeat := 0
	for ind, seat := range seats {
		if seat+1 != seats[ind+1] {
			mySeat = seat + 1
			break
		}
	}

	fmt.Printf("solution for part 1 is: %d\n", seats[len(seats)-1])
	fmt.Printf("solution for part 2 is: %d\n", mySeat)
}

