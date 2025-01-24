package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("3.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	regexPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regexPattern.FindAllStringSubmatch(string(data), -1)
	result := 0
	for _, match := range matches {
		result += sum(match)
	}
	fmt.Printf("Part 1: %v\n", result)
	part2(string(data))
}

func sum(match []string) int {
	a, err := strconv.Atoi(match[1])
	b, err := strconv.Atoi(match[2])
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	return a * b
}

func part2(data string) {
	enabled := true
	regexPattern := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	matches := regexPattern.FindAllStringSubmatch(string(data), -1)
	result := 0
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			result += sum(match)
		}
	}
	fmt.Printf("Part 2: %v\n", result)
}
