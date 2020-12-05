package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"strings"
	"sort"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	ids := []int{}
	
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")
		ID, err := strconv.ParseInt(line, 2, 32)
		if err != nil {
			log.Fatalf("failed to parse ID: %v", err)
		}
		ids = append(ids, int(ID))
	}

	// part 1
	sort.Ints(ids)

	// part 2
	var mySeat int
	for ind := range ids {
		if mySeat =ids[ind]+1; mySeat != ids[ind+1] {
			break
		}
	}

	log.Printf("solution for part 1: %d\n", ids[len(ids)-1])
	log.Printf("solution for part 2: %d\n", mySeat)
}

