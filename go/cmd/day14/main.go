package main

import (
	"fmt"
	"os"

	"github.com/fbegyn/aoc2020/go/helpers"
	"strings"
	"strconv"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	memory := make(map[int]int)
	mask := make(map[int]int, 36)

	memory2 := make(map[int]int)
	mask2 := make(map[int]int, 36)
	
	for inp := range input {
		split := strings.Split(inp, " = ")
		switch split[0][:3] {
		case "mas":
			for ind, ch := range split[1] {
				switch ch {
				case '1':
					mask[36-ind-1] = int(ch-'0')
					mask2[36-ind-1] = int(ch-'0')
				case '0':
					mask[36-ind-1] = int(ch-'0')
					mask2[36-ind-1] = int(ch-'0')
				case 'X':
					delete(mask, 36-ind-1)
					mask2[36-ind-1] = -1
				}
			}
		case "mem":
			addressStr := strings.TrimRight(split[0][4:], "]")
			address, err := strconv.Atoi(addressStr)
			if err != nil {
				panic(err)
			}
			address2 := address
			value, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			for k, v := range mask {
				if v == 1 {
					value = SetBit(value, k)
				}
				if v == 0 {
					value = ClearBit(value, k)
				}
			}
			for k, v := range mask2 {
				if v == 1 || v == 0 {
				    address2 |= v << k
				}
			}
			memory[address] = value
			memory2[address2] = value
		}
	}

	fmt.Println(memory2)

	sum := 0
	for _, v := range memory {
		sum += v
	}

	fmt.Printf("solution to part 1: %d\n", sum)
}

func SetBit(n, pos int) int {
	n |= (1 << pos)
	return n
}

func ClearBit(n, pos int) int {
	n &= ^(1 << pos)
	return n
}
