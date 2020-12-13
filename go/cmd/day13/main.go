package main

import (
	"os"

	"strconv"

	"fmt"
	"strings"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	earliest, err := strconv.ParseUint(<-input, 10, 64)
	if err != nil {
		panic(err)
	}
	inServiceStr := strings.Split(<-input, ",")
	inService := []uint64{}
	offset := []uint64{}
	for ind, bus := range inServiceStr {
		if bus == "x" {
			continue
		}
		busID, err := strconv.ParseUint(bus, 10, 64)
		if err != nil {
			panic(err)
		}
		offset = append(offset, uint64(ind))
		inService = append(inService, busID)
	}

	waitingTime := uint64(100000000)
	waitingBus := uint64(0)
	for _, busID := range inService {
		times := earliest / busID
		nextDeparture := (times + 1) * busID
		if t := nextDeparture - earliest; t < waitingTime {
			waitingTime = t
			waitingBus = busID
		}
	}

	fmt.Printf("solution for part 1: %d\n", waitingTime*waitingBus)


	busID := inService[0]
	times := uint64(0)
	for {
		timestamp := busID * times
		matching := 0
		for ind, busN := range inService[1:] {
			t := timestamp + offset[ind+1]
			if r := t % busN; r != 0 {
				continue
			}
			busID = busN
			matching++
		}
		if matching == len(inService[1:]) {
			fmt.Printf("YES timestamp: %d\n", timestamp)
			break
		}
		times++
	}
}
