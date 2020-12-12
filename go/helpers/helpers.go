package helpers

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// OpenFile well, it opens a file based on a path :p
func OpenFile(f string) (file *os.File) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("failed to read file into scanner: %v", err)
	}
	return
}

func StreamLines(file string, output chan<- string) {
	input := OpenFile(file)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		output <- scanner.Text()
	}
	close(output)
}

func StreamStrings(file string, output chan<- string) {
	input := OpenFile(file)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		output <- scanner.Text()
	}
	close(output)
}

func StreamRunes(file string, output chan<- rune) {
	input := OpenFile(file)
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		output <- []rune(scanner.Text())[0]
	}
	close(output)
}

// SumOfFloat64Array sums all float64 in the array
func SumOfFloat64Array(test []float64) (result float64) {
	for _, v := range test {
		result += v
	}
	return
}

// SumOfIntArray sums all int in the array
func SumOfIntArray(test []int) (result int) {
	for _, v := range test {
		result += v
	}
	return
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func LinesToInts(file *os.File) (ints []int, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		integer, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, integer)
	}
	return
}

func LinesToFloats(file *os.File) (floats []float64, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		float, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		floats = append(floats, float)
	}
	return
}

// Min finds the min in a string - int map
func Min(m map[string]int) (ind string) {
	min := 320000000
	for k, v := range m {
		if v == min {
			ind = "D"
		}
		if v < min {
			min = v
			ind = k
		}
	}
	return
}

// Min finds the min in a string - int map
func MinInt(m []int) (min int) {
	min = 320000000
	for _, v := range m {
		if v == min {
			continue
		}
		if v < min {
			min = v
		}
	}
	return
}

// Max finds the max in a string - int map
func Max(m map[string]int) (ind string) {
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			ind = k
		}
	}
	return
}

func RunProgram(prog []string, output chan<- int, halt, loop chan<- bool) {
	mem := make([]string, len(prog))
	copy(mem, prog)

	instrFreq := make(map[uint]uint)

	pc := uint(0)
	acc := 0
	looping := false

	for !looping {
		if uint(len(mem)) <= pc {
			halt <- true
			output <- acc
			break
		}

		if fr, _ := instrFreq[pc]; 0 < fr {
			looping = true
			loop <- looping
			output <- acc
		}

		instruction := mem[pc]
		opcode := strings.Split(instruction, " ")
		instrFreq[pc] += 1
		switch opcode[0] {
		case "acc":
			arg, err := strconv.Atoi(opcode[1])
			if err != nil {
				log.Fatalf("failed to parse arg: %v", err)
			}
			acc += arg
			pc += 1
		case "jmp":
			arg, err := strconv.Atoi(opcode[1])
			if err != nil {
				log.Fatalf("failed to parse arg: %v", err)
			}
			pc += uint(arg)
		case "nop":
			pc += 1
		}
	}
}

func ToggleInstruction(prog []string, ind int) []string {
	change := make([]string, len(prog))
	copy(change, prog)
	instr := change[ind]
	switch strings.Split(instr, " ")[0] {
	case "jmp":
		change[ind] = strings.ReplaceAll(change[ind], "jmp", "nop")
	case "nop":
		change[ind] = strings.ReplaceAll(change[ind], "nop", "jmp")
	}
	return change
}

type Point struct {
	x, y int64
}

func NewPoint(x, y int64) *Point {
	return &Point{x, y}
}

func (p *Point) Move(n [2]int64) {
	p.x += n[0]
	p.y += n[1]
}

func (p *Point) MoveDir(dir rune) {
	switch {
	case dir == 'N' || dir == 'U':
		p.Move([2]int64{0, 1})
	case dir == 'S' || dir == 'D':
		p.Move([2]int64{0, -1})
	case dir == 'E' || dir == 'R':
		p.Move([2]int64{1, 0})
	case dir == 'W' || dir == 'L':
		p.Move([2]int64{-1, 0})
	}
}

func (p *Point) MoveDirN(dir rune, steps int64) {
	switch {
	case dir == 'N' || dir == 'U':
		p.Move([2]int64{0, steps})
	case dir == 'S' || dir == 'D':
		p.Move([2]int64{0, -1 * steps})
	case dir == 'E' || dir == 'R':
		p.Move([2]int64{steps, 0})
	case dir == 'W' || dir == 'L':
		p.Move([2]int64{-1 * steps, 0})
	}
}

func (p *Point) MoveRelative(n *Point) {
	p.Move([2]int64{n.x, n.y})
}

func (p *Point) MoveRelativeN(n *Point, times int64) {
	p.Move([2]int64{
		times * (n.x),
		times * (n.y),
	})
}

func (p *Point) Angle(t Point) (angle float64) {
	angle = math.Atan2(float64(t.x-p.x), float64(t.y-p.y)) * 180 / math.Pi
	if angle < 0 {
		angle += 360
	}
	return
}

func (p *Point) Rotate90(cc bool) {
	if cc {
		p.x, p.y = -p.y, p.x
	} else {
		p.x, p.y = p.y, -p.x
	}
}

func (p *Point) ManhattanDist(t Point) int64 {
	return Abs(p.x-t.x) + Abs(p.y-t.y)
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(x, y int64) int64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func RenderGrid(grid map[Point]int64) [][]string {
	minX, minY := int64(9999999), int64(99999999)
	maxX, maxY := int64(-9999999), int64(-9999999)
	for k, _ := range grid {
		if k.y < minY {
			minY = k.y
		}
		if k.x < minX {
			minX = k.x
		}
		if maxY < k.y {
			maxY = k.y
		}
		if maxX < k.x {
			maxX = k.x
		}
	}
	height := maxY - minY + 1
	width := maxX - minX + 1
	image := make([][]string, height)
	for k, v := range grid {
		x := k.x - minX
		y := k.y - minY
		if image[y] == nil {
			image[y] = make([]string, width)
		}
		switch v {
		case 0:
			image[y][x] = " "
		case 1:
			image[y][x] = ""
		case 2:
			image[y][x] = ""
		case 3:
			image[y][x] = "-"
		case 4:
			image[y][x] = "o"
		}
	}
	return image
}

func PrintImage(image [][]string) {
	for _, y := range image {
		for _, x := range y {
			fmt.Printf("%s", x)
		}
		fmt.Printf("\n")
	}
}

func RunRobot(grid map[Point]int64, start Point, input <-chan int64, output chan<- int64) {
	direction := 'U'
	location := start

	for {
		// Generate output
		if _, ok := grid[location]; !ok {
			grid[location] = 0
		}
		output <- grid[location]

		// Read input
		instruction := make([]int64, 2)
		for i := range instruction {
			instruction[i] = <-input
		}
		switch instruction[0] {
		case 0:
			grid[location] = 0
		case 1:
			grid[location] = 1
		}
		switch instruction[1] {
		case 0:
			switch direction {
			case 'U':
				direction = 'L'
			case 'L':
				direction = 'D'
			case 'D':
				direction = 'R'
			case 'R':
				direction = 'U'
			}
		case 1:
			switch direction {
			case 'U':
				direction = 'R'
			case 'L':
				direction = 'U'
			case 'D':
				direction = 'L'
			case 'R':
				direction = 'D'
			}
		}

		// Evaluate movement
		switch direction {
		case 'U':
			location.y++
		case 'D':
			location.y--
		case 'L':
			location.x--
		case 'R':
			location.x++
		}
	}
}
