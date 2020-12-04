package main

import (
	"bufio"
	"testing"

	"strings"
	"sync"

	"github.com/fbegyn/aoc2020/go/helpers"
)

func BenchmarkPart1(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day04/input.txt")
	defer input.Close()

	for i := 0; i < b.N; i++ {
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

		for _, pass := range passports {
			if pass.IsValid() {
				validPassports += 1
			}
		}
	}
}

func BenchmarkPart1Stream(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day04/input.txt")
	defer input.Close()

	passportStream := make(chan Passport)
	validStream := make(chan int)
	validPassports := 0

	var wg sync.WaitGroup

	go StreamValidPart1(passportStream, validStream)
	wg.Add(1)
	go StreamCount(&validPassports, validStream, &wg)

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		passportData := []string{}
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, " ")
			passportData = append(passportData, fields...)
			if line == "" {
				passport := PassportFromData(passportData)
				passportStream <- passport
				passportData = []string{}
				continue
			}
		}
		passport := PassportFromData(passportData)
		passportStream <- passport
	}
	close(passportStream)
	wg.Wait()
}

func BenchmarkPart2(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day04/input.txt")
	defer input.Close()

	for i := 0; i < b.N; i++ {
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

		validStrictPassports := 0

		for _, pass := range passports {
			if pass.IsValidStrict() {
				validStrictPassports += 1
			}
		}
	}
}

func BenchmarkPart2Stream(b *testing.B) {
	input := helpers.OpenFile("../../../inputs/day04/input.txt")
	defer input.Close()

	passportStream := make(chan Passport)
	validStream := make(chan int)
	strictValidStream := make(chan int)
	validPassports := 0
	strictValidPassports := 0

	var wg sync.WaitGroup

	go StreamValid(passportStream, validStream, strictValidStream)
	wg.Add(1)
	go StreamCount(&validPassports, validStream, &wg)
	wg.Add(1)
	go StreamCount(&strictValidPassports, strictValidStream, &wg)

	scanner := bufio.NewScanner(input)
	for i := 0; i < b.N; i++ {
		passportData := []string{}
		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Split(line, " ")
			passportData = append(passportData, fields...)
			if line == "" {
				passport := PassportFromData(passportData)
				passportStream <- passport
				passportData = []string{}
				continue
			}
		}
		passport := PassportFromData(passportData)
		passportStream <- passport
	}
	close(passportStream)
	wg.Wait()
}
