package main

import (
	"bufio"

	"log"
	"os"

	"github.com/fbegyn/aoc2020/go/helpers"
)

type Point struct {
	x, y int
	tree bool
}

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	routeMap := [][]rune{}
	rowCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for ind, ch := range line {
			row[ind] = ch
		}
		routeMap = append(routeMap, row)
		rowCount += 1
	}

	position := NewLocation(routeMap)
	sol1, _ := part1(position)
	log.Printf("solution for part 1: %d", sol1)
	sol2, _ := part2(position)
	log.Printf("solution for part 2: %d", sol2*sol1)
}

type Location struct {
	x, y       int
	tree       bool
	routeMap   [][]rune
	rows, cols int
	end        bool
}

func NewLocation(routeMap [][]rune) Location {
	return Location{
		x:        0,
		y:        0,
		tree:     false,
		routeMap: routeMap,
		rows:     len(routeMap),
		cols:     len(routeMap[0]),
		end:      false,
	}
}

func (l *Location) Move(x, y int) {
	l.x += x
	l.y -= y
	if l.y < l.rows {
		l.tree = l.IsTree(l.routeMap[l.y%l.rows][l.x%l.cols])
	} else {
		l.end = true
	}
}

func (l *Location) IsTree(r rune) bool {
	if r == '.' {
		return false
	}
	return true
}

func (l *Location) CountTrees(x, y int) int {
	trees := 0
	for !l.end {
		if l.tree {
			trees += 1
		}
		l.Move(x, y)
	}
	return trees
}

func (l *Location) Reset() {
	l.x = 0
	l.y = 0
	l.end = false
	l.tree = false
}

func part1(pos Location) (int, error) {
	return pos.CountTrees(3, -1), nil
}

func part2(pos Location) (int, error) {
	trees1 := pos.CountTrees(1, -1)
	pos.Reset()
	trees3 := pos.CountTrees(5, -1)
	pos.Reset()
	trees4 := pos.CountTrees(7, -1)
	pos.Reset()
	trees5 := pos.CountTrees(1, -2)
	pos.Reset()
	return trees1 * trees3 * trees4 * trees5, nil
}
