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
	time := uint64(0)

	for ind, busN := range inService[1:] {
		var times uint64
		for {
			times++
			timestamp := time + busID * times
			isDeparture := (timestamp + offset[ind+1]) % busN
			if isDeparture != 0 {
				continue
			}
			time = timestamp
			busID = busID * busN
			break
		}
	}

	fmt.Printf("solution for part 2: %d\n", time)
}
