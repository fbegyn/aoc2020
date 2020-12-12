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

	for inp := range input {
		dir := rune(inp[0])
		steps, err := strconv.ParseInt(inp[1:], 10, 64)
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'N':
			ship.location.MoveDirN(dir, steps)
			waypoint.MoveDirN(dir, steps)
		case 'E':
			ship.location.MoveDirN(dir, steps)
			waypoint.MoveDirN(dir, steps)
		case 'S':
			ship.location.MoveDirN(dir, steps)
			waypoint.MoveDirN(dir, steps)
		case 'W':
			ship.location.MoveDirN(dir, steps)
			waypoint.MoveDirN(dir, steps)
		case 'F':
			ship.location.MoveDirN(ship.direction, steps)
			ship1.location.MoveRelativeN(waypoint, steps)
		case 'R':
			ship.Turn(dir, steps)
			turn := int(steps / 90)
			for i := 0; i < turn; i++ {
				waypoint.Rotate90(false)
			}
		case 'L':
			ship.Turn(dir, steps)
			turn := int(steps / 90)
			for i := 0; i < turn; i++ {
				waypoint.Rotate90(true)
			}
		}
	}

	start := helpers.NewPoint(0, 0)
	fmt.Printf("solution for part 1: %d\n", ship.location.ManhattanDist(*start))
	fmt.Printf("solution for part 2: %d\n", ship1.location.ManhattanDist(*start))
}
