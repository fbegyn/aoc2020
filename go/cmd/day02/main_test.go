package main

import (
	"bufio"
	"regexp"
	"testing"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day02/input.txt")
	defer input.Close()

	validSledPassword := 0

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		for scanner.Scan() {
			line := scanner.Text()

			rule, password := PasswordRuleFromLine(line)

			if rule.ValidSled(password) {
				validSledPassword += 1
			}
		}
	}
}

func BenchmarkPart1Regex(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day02/input.txt")
	defer input.Close()

	validSledPassword := 0

	regex := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)$")

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		for scanner.Scan() {
			line := scanner.Text()

			rule, password := PasswordRuleFromLineRegex(line, regex)

			if rule.ValidSled(password) {
				validSledPassword += 1
			}
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day02/input.txt")
	defer input.Close()

	validTobogganPassword := 0

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		for scanner.Scan() {
			line := scanner.Text()

			rule, password := PasswordRuleFromLine(line)

			if rule.ValidToboggan(password) {
				validTobogganPassword += 1
			}
		}
	}
}

func BenchmarkPart2Regex(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day02/input.txt")
	defer input.Close()

	validTobogganPassword := 0

	regex := regexp.MustCompile("(\\d+)-(\\d+) ([a-z]): ([a-z]+)$")

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		for scanner.Scan() {
			line := scanner.Text()

			rule, password := PasswordRuleFromLineRegex(line, regex)

			if rule.ValidToboggan(password) {
				validTobogganPassword += 1
			}
		}
	}
}
