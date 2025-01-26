package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("5.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var rules [][]int
	var manuals [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			rules = append(rules, _StrSliceToIntSlice(rule))
		}
		if strings.Contains(line, ",") {
			manual := strings.Split(line, ",")
			manuals = append(manuals, _StrSliceToIntSlice(manual))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0
	for _, manual := range manuals {
		result += _ScoreManual(manual, rules)
	}
	fmt.Println(result)
	result = 0
	for _, manual := range manuals {
		result += _ScoreManualPart2(manual, rules)
	}
	fmt.Println(result)
}

func _StrSliceToIntSlice(strs []string) []int {
	var ints []int
	for _, s := range strs {
		var num int
		fmt.Sscanf(s, "%d", &num)
		ints = append(ints, num)
	}
	return ints
}

func _ScoreManual(manual []int, rules [][]int) int {
	for _, rule := range rules {
		i0 := _IndexOf(rule[0], manual)
		i1 := _IndexOf(rule[1], manual)
		if i0 > -1 && i1 > -1 {
			if i0 > i1 {
				return 0
			}
		}
	}
	return manual[len(manual)/2]
}

func _ScoreManualPart2(manual []int, rules [][]int) int {
	for _, rule := range rules {
		i0 := _IndexOf(rule[0], manual)
		i1 := _IndexOf(rule[1], manual)
		if i0 > -1 && i1 > -1 {
			if i0 > i1 {
				_SortInts(manual, func(a, b int) bool {
					for _, rule := range rules {
						if (a == rule[0] && b == rule[1]) || (a == rule[1] && b == rule[0]) {
							return a == rule[0]
						}
					}
					return a < b
				})
				return manual[len(manual)/2]
			}
		}
	}
	return 0
}

func _IndexOf(num int, nums []int) int {
	for i, n := range nums {
		if num == n {
			return i
		}
	}
	return -1
}

func _SortInts(nums []int, compare func(a, b int) bool) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if compare(nums[j], nums[i]) {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
