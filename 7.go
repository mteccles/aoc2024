package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("7.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1 := 0
	part2 := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		testValue := parts[0]
		testValueInt, _ := strconv.Atoi(testValue)
		operands := strings.Split(parts[1], " ")
		operandsInts := make([]int, len(operands))
		for i, operand := range operands {
			operandsInts[i], _ = strconv.Atoi(operand)
		}
		if CanCalibrate(testValueInt, operandsInts, []rune{'+', '*'}) {
			part1 += testValueInt
		}
		if CanCalibrate(testValueInt, operandsInts, []rune{'+', '*', '|'}) {
			part2 += testValueInt
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

}

func CanCalibrate(testValue int, operands []int, operators []rune) bool {
	if len(operands) == 1 {
		return testValue == operands[0]
	}
	if operands[0] > testValue {
		return false
	}
	for _, op := range operators {
		var next int
		if op == '+' {
			next = operands[0] + operands[1]
		} else if op == '*' {
			next = operands[0] * operands[1]
		} else if op == '|' {
			next, _ = strconv.Atoi(strconv.Itoa(operands[0]) + strconv.Itoa(operands[1]))
		}
		if CanCalibrate(testValue, append([]int{next}, operands[2:]...), operators) {
			return true
		}
	}
	return false
}
