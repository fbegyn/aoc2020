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
	l := NewLife(m)
	fmt.Println(l.String())
	l.Step()
	fmt.Println(l.String())
	l.Step()
	fmt.Println(l.String())
	l.Step()
	fmt.Println(l.String())
	l.Step()
	fmt.Println(l.String())
	l.Step()
	fmt.Println(l.String())
}

type Field struct {
	m          [][]byte
	s          [][]bool
	rows, cols int
}

func NewField(m [][]byte) *Field {
	rows := len(m)
	cols := len(m[0])

	s := make([][]bool, rows)
	for i := range s {
		s[i] = make([]bool, rows)
	}

	return &Field{
		m:    m,
		s:    s,
		rows: rows,
		cols: cols,
	}
}

func (f *Field) IsSeat(x, y int) bool {
	if f.m[y][x] != '.' {
		return true
	}
	return false
}

func (f *Field) Set(x, y int, state bool) bool {
	if f.IsSeat(x, y) {
		f.s[y][x] = state
		if state {
			f.m[y][x] = '#'
		} else {
			f.m[y][x] = 'L'
		}
		return true
	}
	return false
}

func (f *Field) Occupied(x, y int) bool {
	return f.s[y][x]
}

func (f *Field) Next(x, y int) int {
	occupied := 0
	for i := -1; i <= 1; i++ {
		if x+i < 0 || f.cols <= x+i {
			continue
		}
		for j := -1; j <= 1; j++ {
			if y+j < 0 || f.rows <= y+j {
				continue
			}
			if !f.IsSeat(x+i, y+j) {
				continue
			}
			if (j != 0 || i != 0) && f.Occupied(x+i, y+j) {
				occupied++
			}
		}
	}
	seatTaken := f.Occupied(x, y)
	if occupied == 0 && !seatTaken {
		return 1
	}
	if occupied >= 4 && seatTaken {
		return -1
	}
	return 0
}

type Life struct {
	c, n       *Field
	rows, cols int
}

func NewLife(m [][]byte) *Life {
	a := NewField(m)
	return &Life{
		c:    a,
		n:    NewField(m),
		rows: len(m),
		cols: len(m[0]),
	}
}

func (l *Life) Step() bool {
	change := false
	for y := 0; y < l.rows; y++ {
		for x := 0; x < l.cols; x++ {
			switch l.c.Next(x, y) {
			case 1:
				l.n.Set(x, y, true)
				change = true
			case -1:
				l.n.Set(x, y, false)
				change = true
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
