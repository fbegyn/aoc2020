package main

import (
	"fmt"
	"os"

	"bytes"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	m := [][]byte{}
	for inp := range input {
		in := make([]byte, len(inp))
		for col, ch := range []byte(inp) {
			in[col] = ch
		}
		m = append(m, in)
	}
	l := NewLife(m, 4)
	changed := true
	for changed {
		changed = l.Step()
	}
	fmt.Printf("solution to part 1: %d\n", l.c.CountOccupied())

	l = NewLife(m, 5)
	changed = true
	for changed {
		changed = l.StepPart2()
	}
	fmt.Printf("solution to part 2: %d\n", l.c.CountOccupied())
}

type Field struct {
	m          [][]byte
	rows, cols int
	threshold  int
}

func NewField(m [][]byte, thresh int) *Field {
	rows := len(m)
	cols := len(m[0])

	mm := make([][]byte, rows)
	for r := range m {
		rr := make([]byte, cols)
		copy(rr, m[r])
		mm[r] = rr
	}

	return &Field{
		m:         mm,
		rows:      rows,
		cols:      cols,
		threshold: thresh,
	}
}

func (f *Field) IsSeat(x, y int) bool {
	if f.m[y][x] != '.' {
		return true
	}
	return false
}

func (f *Field) Set(x, y int, state bool) bool {
	if state {
		f.m[y][x] = '#'
		return true
	}
	f.m[y][x] = 'L'
	return false
}

func (f *Field) Occupied(x, y int) bool {
	x += f.cols
	x %= f.cols
	y += f.rows
	y %= f.rows
	state := f.m[y][x]
	switch state {
	case '#':
		return true
	case 'L':
		return false
	case '.':
		return false
	default:
		fmt.Println("I don't know this state")
		return false
	}
}

func (f *Field) Next(x, y int) int {
	if !f.IsSeat(x, y) {
		return 0
	}
	occupied := 0
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if 0 <= ny && ny < f.rows && 0 <= nx && nx < f.cols {
			if f.Occupied(nx, ny) {
				occupied++
			}
		}
	}
	seatTaken := f.Occupied(x, y)
	if occupied == 0 && !seatTaken {
		return 1
	}
	if occupied >= f.threshold && seatTaken {
		return -1
	}
	return 0
}

func (f *Field) NextPart2(x, y int) int {
	if !f.IsSeat(x, y) {
		return 0
	}
	occupied := 0
	directions := [][2]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		steps := 1
		for 0 <= ny && ny < f.rows && 0 <= nx && nx < f.cols {
			if f.Occupied(nx, ny) {
				occupied++
				break
			} else {
				if f.m[ny][nx] == 'L' {
					break
				}
				nx += dir[0]
				ny += dir[1]
				steps++
			}
		}
	}
	seatTaken := f.Occupied(x, y)
	if occupied == 0 && !seatTaken {
		return 1
	}
	if occupied >= f.threshold && seatTaken {
		return -1
	}
	return 0
}

type Life struct {
	c, n       *Field
	rows, cols int
}

func NewLife(m [][]byte, thresh int) *Life {
	a := NewField(m, thresh)
	b := NewField(m, thresh)
	return &Life{
		c:    a,
		n:    b,
		rows: len(m),
		cols: len(m[0]),
	}
}

func (l *Life) Step() bool {
	change := false
	for y := 0; y < l.rows; y++ {
		for x := 0; x < l.cols; x++ {
			switch op := l.c.Next(x, y); op {
			case 1:
				l.n.Set(x, y, true)
				change = true
				continue
			case -1:
				l.n.Set(x, y, false)
				change = true
				continue
			case 0:
				l.n.m[y][x] = l.c.m[y][x]
			default:
				continue
			}
		}
	}
	l.c, l.n = l.n, l.c
	return change
}

func (l *Life) StepPart2() bool {
	change := false
	for y := 0; y < l.rows; y++ {
		for x := 0; x < l.cols; x++ {
			switch op := l.c.NextPart2(x, y); op {
			case 1:
				l.n.Set(x, y, true)
				change = true
				continue
			case -1:
				l.n.Set(x, y, false)
				change = true
				continue
			case 0:
				l.n.m[y][x] = l.c.m[y][x]
			default:
				continue
			}
		}
	}
	l.c, l.n = l.n, l.c
	return change
}

func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.rows; y++ {
		for x := 0; x < l.cols; x++ {
			buf.WriteByte(l.c.m[y][x])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (f *Field) String() string {
	var buf bytes.Buffer
	for y := 0; y < f.rows; y++ {
		for x := 0; x < f.cols; x++ {
			buf.WriteByte(f.m[y][x])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (f *Field) CountOccupied() int {
	count := 0
	for y := 0; y < f.rows; y++ {
		for x := 0; x < f.cols; x++ {
			if f.m[y][x] == '#' {
				count++
			}
		}
	}
	return count
}
