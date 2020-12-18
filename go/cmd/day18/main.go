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
		case ')':
			ind++
			break
		default:
			ind++
			if '0' <= ch || ch <= '9' {
				number := int(ch - '0')
				if op == '+' {
					result += number
				} else {
					result *= number
				}
			} else {
				break
			}
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
