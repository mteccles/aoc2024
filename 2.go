package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("2.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	report := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(string(line), " ")
		nums := []int{}
		for _, v := range s {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("failed reading file: %s", err)
			}
			nums = append(nums, n)
		}
		report = append(report, nums)
	}
	result := 0
	for i := 0; i < len(report); i++ {
		result += Safe(report[i])
	}
	log.Printf("Part 1: %v\n", result)

	result = 0
	for i := 0; i < len(report); i++ {
		result += SafePart2(report[i])
	}
	log.Printf("Part 2: %v\n", result)
}

func _RemoveElement(report []int, i int) []int {
	s := make([]int, len(report))
	copy(s, report)
	return append(s[:i], s[i+1:]...)
}

func SafePart2(report []int) int {
	if Safe(report) == 1 {
		return 1
	}
	for i := 0; i < len(report); i++ {
		if Safe(_RemoveElement(report, i)) == 1 {
			return 1
		}
	}
	return 0
}

func Safe(report []int) int {
	fmt.Printf("report: %v\n", report)
	asc, desc := false, false
	if report[1] > report[0] {
		asc = true
	} else if report[1] < report[0] {
		desc = true
	} else {
		return 0
	}
	for i := 1; i < len(report); i++ {
		if asc {
			if report[i] > report[i-1] && report[i]-report[i-1] <= 3 {
				continue
			} else {
				return 0
			}
		} else if desc {
			if report[i] < report[i-1] && report[i-1]-report[i] <= 3 {
				continue
			} else {
				return 0
			}
		}
	}

	return 1
}
