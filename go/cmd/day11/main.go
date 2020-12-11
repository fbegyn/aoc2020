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

	m := make(map[helpers.Point]rune)
	rows, cols := 0, 0
	for inp := range input {
		for col, ch := range inp {
			p := helpers.NewPoint(int64(col), int64(rows))
			if cols < col {
				cols = col
			}
			m[*p] = ch
		}
		rows++
	}
	l := NewLife(m, cols, rows, 4)
	changed := true
	for changed {
		l, changed = l.Step(false)
	}
	fmt.Printf("solution to part 1: %d\n", l.CountOccupied())

	l = NewLife(m, cols, rows, 5)
	changed = true
	for changed {
		l, changed = l.Step(true)
	}
	fmt.Printf("solution to part 2: %d\n", l.CountOccupied())

	l = NewLife(m, cols, rows, 4)
	fmt.Println(l.Stabilize(false))
}

type Life struct {
	m                     map[helpers.Point]rune
	threshold, rows, cols int
}

func NewLife(m map[helpers.Point]rune, cols, rows, thresh int) *Life {
	mm := make(map[helpers.Point]rune, cols*rows)
	for k, v := range m {
		mm[k] = v
	}
	return &Life{
		m:         mm,
		threshold: thresh,
		rows:      rows,
		cols:      cols,
	}
}

func (l *Life) IsSeat(p helpers.Point) bool {
	if l.m[p] != '.' {
		return true
	}
	return false
}

func (l *Life) Set(p helpers.Point, state bool) bool {
	if state {
		l.m[p] = '#'
		return true
	}
	l.m[p] = 'L'
	return false
}

func (l *Life) Occupied(p helpers.Point) bool {
	state, ok := l.m[p]
	if ok {
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
	fmt.Println("This point does not exist")
	return false
}

func (l *Life) Next(p helpers.Point, view bool) (int, int) {
	if !l.IsSeat(p) {
		return 0, 0
	}
	occupied := 0
	directions := [][2]int64{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, dir := range directions {
		point := p
		point.Move(dir)
		ch, ok := l.m[point]
		for ok {
			if l.Occupied(point) {
				occupied++
				if view {
					break
				}
			}
			if view {
				if ch == 'L' {
					break
				}
				point.Move(dir)
				ch, ok = l.m[point]
			} else {
				break
			}
		}
	}
	seatTaken := l.Occupied(p)
	//fmt.Printf("point: %v -> occupied: %d - taken: %v\n", p, occupied, seatTaken)
	if occupied == 0 && !seatTaken {
		return 1, occupied
	}
	if occupied >= l.threshold && seatTaken {
		return -1, occupied
	}
	return 0, 0
}

func (l *Life) Step(view bool) (*Life, bool) {
	new := &Life{
		m:         make(map[helpers.Point]rune, l.rows*l.cols),
		threshold: l.threshold,
		rows:      l.rows,
		cols:      l.cols,
	}
	change := false
	for point, ch := range l.m {
		switch op, _ := l.Next(point, view); op {
		case 1:
			new.Set(point, true)
			change = true
			continue
		case -1:
			new.Set(point, false)
			change = true
			continue
		case 0:
			new.m[point] = ch
		default:
			continue
		}
	}
	return new, change
}

func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y <= l.rows; y++ {
		for x := 0; x <= l.cols; x++ {
			point := helpers.NewPoint(int64(x), int64(y))
			buf.WriteRune(l.m[*point])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (l *Life) CountOccupied() int {
	count := 0
	for _, r := range l.m {
		if r == '#' {
			count++
		}
	}
	return count
}

func (l *Life) Stabilize(view bool) int {
	new := &Life{
		m:         make(map[helpers.Point]rune, l.rows*l.cols),
		threshold: l.threshold,
		rows:      l.rows,
		cols:      l.cols,
	}
	check := map[helpers.Point]struct{}{}
	for k, v := range l.m {
		if v == '.' {
			new.m[k] = v
			continue
		}
		check[k] = struct{}{}
	}

	fmt.Println(l.m)
	for 0 < len(check) {
		add := make(map[helpers.Point]rune, l.cols*l.rows)
		permOcc, permEmp := 0,0
		for k := range check {
			switch op, occ := l.Next(k, view); op {
			case 1:
				if occ != 0 {
					continue
				}
				add[k] = '#'
				delete(check, k)
			case -1:
				if occ > 0 {
					add[k] = 'L'
					delete(check, k)
				}
			}
		}
		for k, v := range add {
			new.m[k] = v
		}
	}
	return new.CountOccupied()
}
