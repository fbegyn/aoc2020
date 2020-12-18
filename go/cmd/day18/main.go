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
L:
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
			break L
		}
	}
	return result, ind
}

func eval2(s string, ind int, mult bool) (int, int) {
	result := 0
	op := '+'
L:
	for ind < len(s) {
		switch ch := s[ind]; ch {
		case ' ':
			ind++
		case '+':
			op = '+'
			ind++
		case '*':
			op = '*'
			ind++
			r, j := eval2(s, ind, true)
			ind = j
			result *= r
		case '(':
			ind++
			r, j := eval2(s, ind, false)
			ind = j
			if op == '+' {
				result += r
			} else {
				result *= r
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number := int(ch - '0')
			if op == '+' {
				result += number
			} else {

				result *= number
			}
			ind++
		default:
			if mult {
				break L
			}
			ind++
			break L
		}
	}
	return result, ind
}

func main() {
	file := os.Args[1]
	input := make(chan string)
	go helpers.StreamLines(file, input)

	results := []int{}
	results2 := []int{}
	for inp := range input {
		res, _ := eval(inp, 0)
		res2, _ := eval2(inp, 0, true)
		results = append(results, res)
		results2 = append(results2, res2)
	}

	fmt.Printf("solution to part 1: %d\n", helpers.SumOfIntArray(results))
	fmt.Printf("solution to part 2: %d\n", helpers.SumOfIntArray(results2))
}
