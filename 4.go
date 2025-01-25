package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0
	for i, row := range data {
		for j, char := range row {
			if char == 'X' {
				result += _Search(i, j, data)
			}
		}
	}
	fmt.Println(result)

	result = 0
	for i, row := range data {
		for j, char := range row {
			if char == 'A' {
				if _FindMAS(i, j, data) {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}

func _Search(x, y int, data [][]rune) int {
	directions := []struct{ dx, dy int }{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}, // up, down, left, right
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // diagonals
	}

	count := 0
	for _, dir := range directions {
		if _FindWord("MAS", dir.dx, dir.dy, x, y, data) {
			count++
		}
	}
	return count
}

func _FindWord(word string, dx, dy, x, y int, data [][]rune) bool {
	for _, char := range word {
		x += dx
		y += dy
		if x < 0 || y < 0 || x >= len(data) || y >= len(data[0]) {
			return false
		}
		if data[x][y] != char {
			return false
		}
	}
	return true
}

func _FindMAS(x, y int, data [][]rune) bool {
	if x == 0 || y == 0 || x == len(data)-1 || y == len(data[0])-1 {
		return false
	}
	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 1}, {1, 1}, {1, -1}, // diagonals
	}
	chars := []rune{}
	for _, dir := range directions {
		chars = append(chars, data[x+dir.dx][y+dir.dy])
	}
	return string(chars) == "MMSS" || string(chars) == "SSMM" || string(chars) == "MSSM" ||
		string(chars) == "SMMS"
}
