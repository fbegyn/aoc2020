package main

import (
	"bufio"
	"testing"

	"sort"
	"strconv"
	"strings"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day05/input.txt")
	defer input.Close()

	ids := []int{}

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.ReplaceAll(line, "F", "0")
			line = strings.ReplaceAll(line, "B", "1")
			line = strings.ReplaceAll(line, "L", "0")
			line = strings.ReplaceAll(line, "R", "1")
			ID, err := strconv.ParseInt(line, 2, 32)
			if err != nil {
				b.Fatalf("failed to parse ID: %v", err)
			}
			ids = append(ids, int(ID))
		}

		// part 1
		sort.Ints(ids)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day05/input.txt")
	defer input.Close()

	ids := []int{}

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.ReplaceAll(line, "F", "0")
			line = strings.ReplaceAll(line, "B", "1")
			line = strings.ReplaceAll(line, "L", "0")
			line = strings.ReplaceAll(line, "R", "1")
			ID, err := strconv.ParseInt(line, 2, 32)
			if err != nil {
				b.Fatalf("failed to parse ID: %v", err)
			}
			ids = append(ids, int(ID))
		}

		// part 1
		sort.Ints(ids)

		// part 2
		var myseat int
		for ind := range ids {
			if myseat = ids[ind] + 1; myseat != ids[ind+1] {
				break
			}
		}
	}
}
