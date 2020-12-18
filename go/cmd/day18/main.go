package main

import (
	"fmt"
	"os"

	_ "strconv"
	_ "strings"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func eval(s string, ind int) (int, int) {
	result := int(0)
	op := '+'
	for ind < len(s) {
		switch ch := s[ind]; ch {
		case ' ':
			ind++
			continue
		case '+':
			ind++
			op = '+'
		case '*':
			ind++
			op = '*'
		case '(':
			ind++
			r, j := eval(s, ind)
			ind = j
			if op == '+' {
				result += r
			} else {
				result *= r
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			ind++
			number := int(ch - '0')
			if op == '+' {
				result += number
			} else {
				result *= number
			}
		default:
			ind++
			break
		}
	}
	return result, ind
}

func main() {
	file := os.Args[1]
	input := make(chan string)
	go helpers.StreamLines(file, input)

	results := []int{}
	for inp := range input {
		res, _ := eval(inp, 0)
		results = append(results, res)
	}

	fmt.Println(results)

	fmt.Printf("solution to part 1: %d\n", helpers.SumOfIntArray(results))
}
