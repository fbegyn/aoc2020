package main

import (
	"fmt"

	"bufio"
	"strings"

	"log"
	"regexp"
	"strconv"

	"github.com/fbegyn/aoc2020/go/helpers"
)

type Passport struct {
	byr, iyr, eyr int
	hgt           int
	hgtUnit       string
	hcl           string
	ecl           string
	pid           string
	cid           string
}

func PassportFromData(data []string) Passport {
	passport := Passport{}
	for _, field := range data {
		trimmed := strings.TrimSpace(field)
		kv := strings.Split(trimmed, ":")
		switch kv[0] {
		case "byr":
			year, err := strconv.Atoi(kv[1])
			if err != nil {
				log.Fatalln(err)
			}
			passport.byr = year
		case "iyr":
			year, err := strconv.Atoi(kv[1])
			if err != nil {
				log.Fatalln(err)
			}
			passport.iyr = year
		case "eyr":
			year, err := strconv.Atoi(kv[1])
			if err != nil {
				log.Fatalln(err)
			}
			passport.eyr = year
		case "hgt":
			index := strings.IndexRune(kv[1], 'c')
			if index != -1 {
				heightString := kv[1][:index]
				height, err := strconv.Atoi(heightString)
				if err != nil {
					log.Fatalln(err)
				}
				passport.hgt = height
				passport.hgtUnit = "cm"
				continue
			}
			index = strings.IndexRune(kv[1], 'i')
			if index != -1 {
				heightString := kv[1][:index]
				height, err := strconv.Atoi(heightString)
				if err != nil {
					log.Fatalln(err)
				}
				passport.hgt = height
				passport.hgtUnit = "in"
				continue
			}
			passport.hgtUnit = "unset"
		case "hcl":
			passport.hcl = kv[1]
		case "ecl":
			passport.ecl = kv[1]
		case "pid":
			passport.pid = kv[1]
		case "cid":
			passport.cid = kv[1]
		}
	}
	return passport
}

func (p *Passport) IsValid() bool {
	if p.byr == 0 {
		return false
	}
	if p.iyr == 0 {
		return false
	}
	if p.eyr == 0 {
		return false
	}
	if p.hgtUnit == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.pid == "" {
		return false
	}
	return true
}

func (p *Passport) IsValidStrict() bool {
	if p.byr < 1920 || 2002 < p.byr {
		return false
	}
	if p.iyr < 2010 || 2020 < p.iyr {
		return false
	}
	if p.eyr < 2020 || 2030 < p.eyr {
		return false
	}
	switch p.hgtUnit {
	case "cm":
		if p.hgt < 150 || 193 < p.hgt {
			return false
		}
	case "in":
		if p.hgt < 59 || 76 < p.hgt {
			return false
		}
	case "unset":
		return false
	default:
		return false
	}
	hclRegex := regexp.MustCompile("#[0-9a-f]{6}")
	if !hclRegex.MatchString(p.hcl) {
		return false
	}
	eclRegex := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
	if !eclRegex.MatchString(p.ecl) {
		return false
	}
	pidRegex := regexp.MustCompile("^[\\d]{9}$")
	if !pidRegex.MatchString(p.pid) {
		return false
	}
	return true
}

func main() {
	input := helpers.OpenFile("../inputs/day04/input.txt")
	defer input.Close()

	passports := []Passport{}

	scanner := bufio.NewScanner(input)
	passportData := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		passportData = append(passportData, fields...)
		if line == "" {
			passport := PassportFromData(passportData)
			passports = append(passports, passport)
			passportData = []string{}
			continue
		}
	}
	passport := PassportFromData(passportData)
	passports = append(passports, passport)

	validPassports := 0
	validStrictPassports := 0

	for _, pass := range passports {
		if pass.IsValid() {
			validPassports += 1
		}
		if pass.IsValidStrict() {
			validStrictPassports += 1
		}
	}

	fmt.Printf("solution for part 1: %d\n", validPassports)
	fmt.Printf("solution for part 2: %d\n", validStrictPassports)
}
