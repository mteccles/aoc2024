package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("10.txt")
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, char := range line {
			row[i] = int(char - '0')
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1, part2 := solution(&grid)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

type Point struct {
	x, y int
}

func solution(grid *[][]int) (int, int) {
	part1 := 0
	part2 := 0
	for y, row := range *grid {
		for x, val := range row {
			if val == 0 {
				p1, p2 := numTrailHeads(grid, Point{x, y})
				part1 += p1
				part2 += p2
			}
		}
	}
	return part1, part2
}

func numTrailHeads(grid *[][]int, p Point) (int, int) {
	part1 := make(map[Point]struct{})
	part2 := 0
	dfs(grid, p, &part1, &part2)
	return len(part1), part2
}

func dfs(grid *[][]int, p Point, endPoints *map[Point]struct{}, numPaths *int) {
	a := (*grid)[p.y][p.x]
	if a == 9 {
		(*endPoints)[p] = struct{}{}
		*numPaths++
		return
	}
	if p.y-1 >= 0 && (*grid)[p.y-1][p.x] == a+1 { // Up
		dfs(grid, Point{p.x, p.y - 1}, endPoints, numPaths)
	}
	if p.x+1 < len((*grid)[0]) && (*grid)[p.y][p.x+1] == a+1 { // Right
		dfs(grid, Point{p.x + 1, p.y}, endPoints, numPaths)
	}
	if p.y+1 < len(*grid) && (*grid)[p.y+1][p.x] == a+1 { // Down
		dfs(grid, Point{p.x, p.y + 1}, endPoints, numPaths)
	}
	if p.x-1 >= 0 && (*grid)[p.y][p.x-1] == a+1 { // Left
		dfs(grid, Point{p.x - 1, p.y}, endPoints, numPaths)
	}
}
