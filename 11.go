package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("11.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := string(data)
	splitContent := strings.Split(content, " ")

	cache = map[CacheKey]int{} // stone, # blinks -> # output stones
	part1, part2 := 0, 0
	for _, stone := range splitContent {
		stoneInt, _ := strconv.Atoi(stone)
		part1 += numStones(stoneInt, 25)
	}
	fmt.Println(part1)

	for _, stone := range splitContent {
		stoneInt, _ := strconv.Atoi(stone)
		part2 += numStones(stoneInt, 75)
	}
	fmt.Println(part2)
}

var multiplier = 2024

var Key struct {
	stone  int
	blinks int
}

type CacheKey struct {
	stone, blinks int
}

var cache map[CacheKey]int

func numStones(stone, blinks int) int {
	if blinks == 0 {
		return 1
	}
	key := CacheKey{stone, blinks}
	if val, found := cache[key]; found {
		return val
	}
	if stone == 0 {
		n := numStones(1, blinks-1)
		cache[CacheKey{stone, blinks}] = n
		return n
	} else if stoneStr := strconv.Itoa(stone); len(stoneStr)%2 == 0 {
		a := removeLeadingZeros(stoneStr[0 : len(stoneStr)/2])
		b := removeLeadingZeros(stoneStr[len(stoneStr)/2:])
		aInt, _ := strconv.Atoi(a)
		bInt, _ := strconv.Atoi(b)
		n := numStones(aInt, blinks-1) + numStones(bInt, blinks-1)
		cache[CacheKey{stone, blinks}] = n
		return n
	} else {
		n := numStones(stone*multiplier, blinks-1)
		cache[CacheKey{stone, blinks}] = n
		return n
	}

}

func removeLeadingZeros(stone string) string {
	for strings.HasPrefix(stone, "00") {
		stone = stone[1:]
	}
	return stone
}
