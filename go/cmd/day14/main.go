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

	memory := make(map[string]int64)
	memory2 := make(map[string]int64)
	var and, or int64
	var mask string

	for inp := range input {
		split := strings.Split(inp, " = ")
		switch split[0][:3] {
		case "mas":
			and, _ = strconv.ParseInt(strings.ReplaceAll(split[1], "X", "1"), 2, 0)
			or, _ = strconv.ParseInt(strings.ReplaceAll(split[1], "X", "0"), 2, 0)
			mask = split[1]
		case "mem":
			addressStr := strings.TrimRight(split[0][4:], "]")
			address, _ := strconv.ParseInt(addressStr, 10, 64)
			val, _ := strconv.Atoi(split[1])
			value := int64(val)

			for _, genAddr := range generateAddresses("", mask, fmt.Sprintf("%036b", address)) {
				memory2[genAddr] = value
			}
			
			value &= and
			value |= or
			memory[addressStr] = value
		}
	}

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	sum2 := int64(0)
	for _, v := range memory2 {
		sum2 += v
	}

	fmt.Printf("solution to part 1: %d\n", sum)
	fmt.Printf("solution to part 2: %d\n", sum2)
}

func generateAddresses(mask, remain, addr string) []string {
	if len(remain) == 0 {
		return []string{mask}
	}

	switch remain[0] {
	case '0':
		return generateAddresses(mask+string(addr[len(mask)]), remain[1:], addr)
	case '1':
		return generateAddresses(mask+"1", remain[1:], addr)
	case 'X':
		return append(
			generateAddresses(mask+"0", remain[1:], addr),
			generateAddresses(mask+"1", remain[1:], addr)...
		)
	default:
		panic(fmt.Errorf("unkown mask %s", string(remain[0])))
	}
}
