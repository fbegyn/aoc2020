package main

import (
	"fmt"
	"os"

	"strconv"
	"strings"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func scanTicket(ticket string) (t []int) {
	split := strings.Split(ticket, ",")
	for _, number := range split {
		n, _ := strconv.Atoi(number)
		t = append(t, n)
	}
	return
}

func CheckIntRange(check int, r [2]int) bool {
	return r[0] <= check && check <= r[1]
}

func CheckIntRanges(check int, rs map[string][2][2]int) bool {
	for _, r := range rs {
		lower := CheckIntRange(check, r[0])
		upper := CheckIntRange(check, r[1])
		if lower || upper {
			return true
		}
	}
	return false
}

func ValidateTicket(ticket []int, rs map[string][2][2]int) (int, bool) {
	for _, number := range ticket {
		if !CheckIntRanges(number, rs) {
			return number, false
		}
	}
	return -1, true
}

func main() {
	file := os.Args[1]
	input := make(chan string, 5)
	go helpers.StreamLines(file, input)

	ruleScan := true
	ticketScan := false
	nearbyTicketScan := false

	validBounds := make(map[string][2][2]int)
	ownTicket := []int{}
	nearbyTicket := [][]int{}
	invalidNumbers := []int{}

	for inp := range input {
		switch {
		case ruleScan:
			if inp == "your ticket:" {
				ruleScan = false
				ticketScan = true
				continue
			}
			if inp == "" {
				continue
			}
			split := strings.Split(inp, ": ")
			ranges := strings.Split(split[1], " or ")
			for ind, r := range ranges {
				bounds := strings.Split(r, "-")
				lower, _ := strconv.Atoi(bounds[0])
				upper, _ := strconv.Atoi(bounds[1])
				b := validBounds[split[0]]
				b[ind] = [2]int{lower, upper}
				validBounds[split[0]] = b
			}
		case ticketScan:
			if inp == "nearby tickets:" {
				ticketScan = false
				nearbyTicketScan = true
				continue
			}
			if inp == "" {
				continue
			}
			ownTicket = scanTicket(inp)
		case nearbyTicketScan:
			t := scanTicket(inp)
			number, valid := ValidateTicket(t, validBounds)
			if valid {
				nearbyTicket = append(nearbyTicket, t)
			} else {
				invalidNumbers = append(invalidNumbers, number)
			}
		}
	}

	candidates := make(map[string][]int)
	for rule, bounds := range validBounds {
		var count int
		for col := range nearbyTicket[0] {
			count = 0
			for _, row := range nearbyTicket {
				lower := CheckIntRange(row[col], bounds[0])
				upper := CheckIntRange(row[col], bounds[1])
				if lower != upper {
					count++
				}
			}
			if count == len(nearbyTicket) {
				candidates[rule] = append(candidates[rule], col)
			}
		}
	}

	// filter through the candidates
	// work means we aren't done yet
	work := true
	for work {
		// go through each rule and it's possible candidates
		for rule, candids := range candidates {
			// if there is more than 1 candidate, we need to filter more before
			// processing this ruleset. There also needs to be a candidate for
			// every rule
			if len(candids) > 1 {
				continue
			} else if len(candids) == 0 {
				panic("oops, no candidates found")
			}

			// if there is only 1 candidate, select it for processing
			correct := candids[0]
			// iterate through the other ruleset to filter out the current
			for k, v := range candidates {
				// skip rulesets with 1 candidate and of course, the
				// current ruleset
				if len(v) == 1 || k == rule {
					continue
				}
				// filter out the current col from the other candidates
				for i, c := range v {
					if c == correct {
						candidates[k] = append(v[:i],v[i+1:]...)
						break
					}
				}
			}
		}
		// if we are here, we assume no work is done
		work = false
		// if there is a ruleset with more than 1 candidate we need to
		// start work again untill all rules are filtered
		for _, c := range candidates {
			if len(c) > 1 {
				work = true
				break
			}
		}
	}

	depart := []int{}
	for k, v := range candidates {
		if strings.HasPrefix(k, "departure") {
			depart = append(depart, v[0])
		}
	}
	
	mult := ownTicket[depart[0]]
        for _, n := range depart[1:] {
                mult *= ownTicket[n]
        }

	fmt.Printf("solution for part 1: %d\n", helpers.SumOfIntArray(invalidNumbers))
	fmt.Printf("solution for part 2: %d\n", mult)
}
