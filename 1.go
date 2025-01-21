package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("1.txt")
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	defer data.Close()

	lhs, rhs := []int{}, []int{}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(string(line), "   ")
		l, err := strconv.Atoi(s[0])
		r, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatalf("failed reading file: %s", err)
		}
		lhs = append(lhs, l)
		rhs = append(rhs, r)
	}
	sort.Ints(lhs)
	sort.Ints(rhs)

	// Part 1
	result := 0
	for i := 0; i < len(lhs); i++ {
		result += int(math.Abs(float64(lhs[i] - rhs[i])))
	}
	fmt.Printf("Part 1: %v\n", result)

	// Part 2
	result = 0
	for i := 0; i < len(lhs); i++ {
		for j := 0; j < len(rhs); j++ {
			if lhs[i] == rhs[j] {
				result += lhs[i]
			} else if lhs[i] > rhs[j] {
				// It's sorted, so we can stop iterating rhs.
				continue
			}

		}
	}
	fmt.Printf("Part 2: %v\n", result)
}
