package main

import (
	"fmt"

	"github.com/fbegyn/aoc2020/go/helpers"
	"bufio"
	"strings"
	"strconv"
	"log"
	"regexp"
	"os"
)

type PasswordRule struct {
	limits []int
	letter string
}

func PasswordRuleFromLine(line string) (PasswordRule, string) {
	set1 := strings.Split(line, ":")
	set2 := strings.Split(set1[0], " ")
	set3 := strings.Split(set2[0], "-")

	lower, err := strconv.Atoi(set3[0])
	if err != nil {
		log.Fatalln(err)
	}
	upper, err := strconv.Atoi(set3[1])
	if err != nil {
		log.Fatalln(err)
	}

	return PasswordRule{
		limits: []int{lower, upper},
		letter: set2[1],
	}, set1[1]
}

// regex := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)$")
func PasswordRuleFromLineRegex(line string, regex *regexp.Regexp) (PasswordRule, string) {
	matches := regex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return PasswordRule{}, ""
	}

	lower, err := strconv.Atoi(matches[0][1])
	if err != nil {
		log.Fatalln(err)
	}
	upper, err := strconv.Atoi(matches[0][2])
	if err != nil {
		log.Fatalln(err)
	}

	return PasswordRule{
		limits: []int{lower, upper},
		letter: matches[0][3],
	}, matches[0][4]
}

func (r *PasswordRule) ValidSled(p string) bool {
	count := strings.Count(p, r.letter)
	if count < r.limits[0] || r.limits[1] < count {
		return false
	}
	return true
}

func (r *PasswordRule) ValidToboggan(p string) bool {
	count := 0
	for _, v := range r.limits {
		if p[v-1] == r.letter[0] {
			count += 1
		}
	}
	if count != 1 {
		return false
	}
	return true
}

func main() {
	file := os.Args[1]
	input := helpers.OpenFile(file)
	defer input.Close()

	validSledPassword := 0
	validTobogganPassword := 0

	//regex := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)$")

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		rule, password := PasswordRuleFromLine(line)
		//r, p := PasswordRuleFromLineRegex(line, regex)

		if rule.ValidSled(password) {
			validSledPassword += 1
		}

		if rule.ValidToboggan(password) {
			validTobogganPassword += 1
		}
	}

	fmt.Println(validSledPassword)
	fmt.Println(validTobogganPassword)
}

func part1() (error) {
	return nil
}

func part2() (error) {
	return nil
}
