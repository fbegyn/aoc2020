package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/fbegyn/aoc2020/go/helpers"
)

type Ship struct {
	location  *helpers.Point
	direction rune
}

func (pos *Ship) Turn(dir rune, degrees int64) {
	for i := int64(0); i < degrees/90; i++ {
		switch dir {
		case 'L':
			switch pos.direction {
			case 'N':
				pos.direction = 'W'
			case 'E':
				pos.direction = 'N'
			case 'S':
				pos.direction = 'E'
			case 'W':
				pos.direction = 'S'
			}
		case 'R':
			switch pos.direction {
			case 'N':
				pos.direction = 'E'
			case 'E':
				pos.direction = 'S'
			case 'S':
				pos.direction = 'W'
			case 'W':
				pos.direction = 'N'
			}
		}
	}
}

func (pos *Ship) Move(dir rune, steps int64) {
	switch dir {
	case 'N':
		mod := [2]int64{0, steps}
		pos.location.Move(mod)
	case 'S':
		mod := [2]int64{0, -1 * steps}
		pos.location.Move(mod)
	case 'E':
		mod := [2]int64{steps, 0}
		pos.location.Move(mod)
	case 'W':
		mod := [2]int64{-1 * steps, 0}
		pos.location.Move(mod)
	case 'L':
		pos.Turn(dir, steps)
	case 'R':
		pos.Turn(dir, steps)
	case 'F':
		pos.Move(pos.direction, steps)
	}
}

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	ship := Ship{
		location:  helpers.NewPoint(0, 0),
		direction: 'E',
	}

	ship1 := Ship{
		location:  helpers.NewPoint(0, 0),
		direction: 'E',
	}

	waypoint := helpers.NewPoint(10,1)
	fmt.Println(waypoint)

	for inp := range input {
		dir := rune(inp[0])
		steps, err := strconv.ParseInt(inp[1:], 10, 64)
		if err != nil {
			panic(err)
		}
		ship.Move(dir, steps)

		if dir == 'F' {
			ship1.Move(dir, steps)
		}
	}

	start := helpers.NewPoint(0, 0)
	fmt.Printf("solution for part 1: %d\n", ship.location.ManhattanDist(*start))
}
