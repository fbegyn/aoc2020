package main

import (
	"log"
	"os"
	"strings"

	"strconv"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
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
	goldBagIndex := 0
	for ind, bag := range bags {
		if bag.Contains(search) {
			goldBags += 1
		}
		if bag.Compare(search) {
			goldBagIndex = ind
		}
	}

	log.Printf("solution for part 1: %d", goldBags)
	log.Printf("solution for part 2: %d", bags[goldBagIndex].CountBags())
}

type Bag struct {
	typ     string
	color   string
	count   []int
	content []*Bag
}

func (b *Bag) Compare(ba Bag) bool {
	return b.typ == ba.typ && b.color == ba.color
}

func (b *Bag) CountBags() int {
	count := 0
	for ind := range b.content {
		count += b.count[ind] + b.count[ind] * b.content[ind].CountBags()
	}
	return count
}

func (b *Bag) Contains(ba Bag) bool {
	for ind := range b.content {
		if b.content[ind].Compare(ba) {
			return true
		}
		if b.content[ind].Contains(ba) {
			return true
		}
	}
	return false
}

func NewBagFromLine(line string) (b *Bag) {
	bag := Bag{}
	line = strings.TrimRight(line, ".")
	split := strings.SplitAfter(line, "contain")

	bagDesc := strings.Split(split[0], " ")
	bag.typ = bagDesc[0]
	bag.color = bagDesc[1]

	allContents := strings.TrimSpace(split[1])
	contents := strings.Split(allContents, ",")

	for _, cont := range contents {
		cont = strings.TrimSpace(cont)
		b := ParseContents(cont)
		if b != nil {
			bag.content = append(bag.content, b)
			bag.count = append(bag.count, b.count...)
		}
	}

	return &bag
}

func ParseContents(desc string) *Bag {
	content := strings.Split(desc, " ")
	if content[0] == "no" {
		return nil
	}
	count, err := strconv.Atoi(content[0])
	if err != nil {
		log.Fatalln(err)
	}
	return &Bag{
		typ:     content[1],
		color:   content[2],
		count:   []int{count},
		content: nil,
	}
}
