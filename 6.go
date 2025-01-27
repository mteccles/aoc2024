package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("6.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var data [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		data = append(data, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// x is the column, y is the row
	var startingX, startingY int
	// Point up by default.
	direction := []int{0, -1}
	for i, row := range data {
		for j, char := range row {
			if char == '^' {
				startingY = i
				startingX = j
			}
		}
	}
	visited := NewGenericSet[[2]int]()
	x := startingX
	y := startingY
	for {
		visited.Add([2]int{x, y})
		// Check if new position is out of bounds
		if x == 0 || x == len(data[0])-1 || y == 0 || y == len(data)-1 {
			break
		}

		newX := x + direction[0]
		newY := y + direction[1]

		// Check if new position is an obstacle
		up := []int{0, -1}
		down := []int{0, 1}
		left := []int{-1, 0}
		right := []int{1, 0}
		if data[newY][newX] == '#' {
			// Change direction clockwise
			if direction[0] == up[0] && direction[1] == up[1] { // up
				direction = right
			} else if direction[0] == right[0] && direction[1] == right[1] { // right
				direction = down
			} else if direction[0] == down[0] && direction[1] == down[1] { //down
				direction = left
			} else if direction[0] == left[0] && direction[1] == left[1] { // left
				direction = up
			}
		} else {
			// Move to the new position
			x = newX
			y = newY
		}
	}
	fmt.Println("Number of visited positions:", visited.Size())

	part2 := 0
	for pos := range visited.elements {
		if pos[0] == startingX && pos[1] == startingY {
			continue
		}
		// Make a copy of the data
		dataCopy := make([][]rune, len(data))
		for i := range data {
			dataCopy[i] = make([]rune, len(data[i]))
			copy(dataCopy[i], data[i])
		}
		dataCopy[pos[1]][pos[0]] = '#'
		if _Circular(dataCopy, startingX, startingY) {
			part2++
		}
	}
	fmt.Println("Number of circular paths:", part2)

}

func _Circular(data [][]rune, x, y int) bool {
	// x, y, dx, dy
	visitedWithDirection := NewGenericSet[[4]int]()
	direction := []int{0, -1}
	for {
		if visitedWithDirection.Contains([4]int{x, y, direction[0], direction[1]}) {
			// We have visited this position with this direction before. This is the definition of a circular path.
			return true
		}
		visitedWithDirection.Add([4]int{x, y, direction[0], direction[1]})
		// Check if new position is out of bounds
		if x == 0 || x == len(data[0])-1 || y == 0 || y == len(data)-1 {
			break
		}

		// Below code is copied from part 1. Should refactor it to a function.
		newX := x + direction[0]
		newY := y + direction[1]

		// Check if new position is an obstacle
		up := []int{0, -1}
		down := []int{0, 1}
		left := []int{-1, 0}
		right := []int{1, 0}
		if data[newY][newX] == '#' {
			// Change direction clockwise
			if direction[0] == up[0] && direction[1] == up[1] { // up
				direction = right
			} else if direction[0] == right[0] && direction[1] == right[1] { // right
				direction = down
			} else if direction[0] == down[0] && direction[1] == down[1] { //down
				direction = left
			} else if direction[0] == left[0] && direction[1] == left[1] { // left
				direction = up
			}
		} else {
			// Move to the new position
			x = newX
			y = newY
		}
	}
	return false
}

// GenericSet is a generic set implementation
type GenericSet[T comparable] struct {
	elements map[T]struct{}
}

func NewGenericSet[T comparable]() GenericSet[T] {
	return GenericSet[T]{elements: make(map[T]struct{})}
}

func (s GenericSet[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s GenericSet[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s GenericSet[T]) Size() int {
	return len(s.elements)
}
