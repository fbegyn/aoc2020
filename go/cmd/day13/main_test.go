package main

import (
	"strconv"
	"testing"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	file := "../../../inputs/day12/input.txt"
	for i := 0; i < b.N; i++ {
		ship := Ship{
			location:  helpers.NewPoint(0, 0),
			direction: 'E',
		}
		input := make(chan string, 5)
		go helpers.StreamLines(file, input)
		for inp := range input {
			dir := rune(inp[0])
			steps, err := strconv.ParseInt(inp[1:], 10, 64)
			if err != nil {
				panic(err)
			}

			switch dir {
			case 'N':
				ship.location.MoveDirN(dir, steps)
			case 'E':
				ship.location.MoveDirN(dir, steps)
			case 'S':
				ship.location.MoveDirN(dir, steps)
			case 'W':
				ship.location.MoveDirN(dir, steps)
			case 'F':
				ship.location.MoveDirN(ship.direction, steps)
			case 'R':
				ship.Turn(dir, steps)
			case 'L':
				ship.Turn(dir, steps)
			}
		}
		start := helpers.NewPoint(0, 0)
		ship.location.ManhattanDist(*start)
	}
}
func BenchmarkPart2(b *testing.B) {
	file := "../../../inputs/day12/input.txt"
	for i := 0; i < b.N; i++ {
		ship1 := Ship{
			location:  helpers.NewPoint(0, 0),
			direction: 'E',
		}
		waypoint := helpers.NewPoint(10, 1)

		input := make(chan string, 5)
		go helpers.StreamLines(file, input)
		for inp := range input {
			dir := rune(inp[0])
			steps, err := strconv.ParseInt(inp[1:], 10, 64)
			if err != nil {
				panic(err)
			}

			switch dir {
			case 'N':
				waypoint.MoveDirN(dir, steps)
			case 'E':
				waypoint.MoveDirN(dir, steps)
			case 'S':
				waypoint.MoveDirN(dir, steps)
			case 'W':
				waypoint.MoveDirN(dir, steps)
			case 'F':
				ship1.location.MoveRelativeN(waypoint, steps)
			case 'R':
				turn := int(steps / 90)
				for i := 0; i < turn; i++ {
					waypoint.Rotate90(false)
				}
			case 'L':
				turn := int(steps / 90)
				for i := 0; i < turn; i++ {
					waypoint.Rotate90(true)
				}
			}
		}
		start := helpers.NewPoint(0, 0)
		ship1.location.ManhattanDist(*start)
	}
}
func BenchmarkFull(b *testing.B) {
	file := "../../../inputs/day12/input.txt"
	for i := 0; i < b.N; i++ {
		ship := Ship{
			location:  helpers.NewPoint(0, 0),
			direction: 'E',
		}
		ship1 := Ship{
			location:  helpers.NewPoint(0, 0),
			direction: 'E',
		}
		waypoint := helpers.NewPoint(10, 1)

		input := make(chan string, 5)
		go helpers.StreamLines(file, input)
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
		ship.location.ManhattanDist(*start)
		ship1.location.ManhattanDist(*start)
	}
}
