package main

import (
	"testing"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	file := "../../../inputs/day07/input.txt"
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	bags := []*Bag{}

	for inp := range input {
		bag := NewBagFromLine(inp)
		bags = append(bags, bag)
	}

	for i := 0; i < b.N; i++ {
		for _, bag := range bags {
			for ind := range bag.content {
				for _, ba := range bags[1:] {
					if bag.content[ind].Compare(*ba) {
						bag.content[ind] = ba
						continue
					}
				}
			}
		}

		search := Bag{
			"shiny",
			"gold",
			[]int{0},
			nil,
		}

		goldBags := 0
		for _, bag := range bags {
			if bag.Contains(search) {
				goldBags += 1
			}
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	file := "../../../inputs/day07/input.txt"
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	bags := []*Bag{}

	for inp := range input {
		bag := NewBagFromLine(inp)
		bags = append(bags, bag)
	}

	for i := 0; i < b.N; i++ {
		for _, bag := range bags {
			for ind := range bag.content {
				for _, ba := range bags[1:] {
					if bag.content[ind].Compare(*ba) {
						bag.content[ind] = ba
						continue
					}
				}
			}
		}

		search := Bag{
			"shiny",
			"gold",
			[]int{0},
			nil,
		}

		goldBagIndex := 0
		for ind, bag := range bags {
			if bag.Compare(search) {
				goldBagIndex = ind
			}
		}
		bags[goldBagIndex].CountBags()
	}
}

func BenchmarkPart1ReadInput(b *testing.B) {
	file := "../../../inputs/day07/input.txt"
	for i := 0; i < b.N; i++ {
		input := make(chan string, 5)
		go helpers.StreamLines(file, input)

		bags := []*Bag{}

		for inp := range input {
			bag := NewBagFromLine(inp)
			bags = append(bags, bag)
		}

		for _, bag := range bags {
			for ind := range bag.content {
				for _, ba := range bags[1:] {
					if bag.content[ind].Compare(*ba) {
						bag.content[ind] = ba
						continue
					}
				}
			}
		}

		search := Bag{
			"shiny",
			"gold",
			[]int{0},
			nil,
		}

		goldBags := 0
		for _, bag := range bags {
			if bag.Contains(search) {
				goldBags += 1
			}
		}
	}
}

func BenchmarkPart2ReadInput(b *testing.B) {
	file := "../../../inputs/day07/input.txt"
	for i := 0; i < b.N; i++ {
		input := make(chan string, 5)
		go helpers.StreamLines(file, input)

		bags := []*Bag{}

		for inp := range input {
			bag := NewBagFromLine(inp)
			bags = append(bags, bag)
		}

		for _, bag := range bags {
			for ind := range bag.content {
				for _, ba := range bags[1:] {
					if bag.content[ind].Compare(*ba) {
						bag.content[ind] = ba
						continue
					}
				}
			}
		}

		search := Bag{
			"shiny",
			"gold",
			[]int{0},
			nil,
		}

		goldBagIndex := 0
		for ind, bag := range bags {
			if bag.Compare(search) {
				goldBagIndex = ind
			}
		}
		bags[goldBagIndex].CountBags()
	}
}
