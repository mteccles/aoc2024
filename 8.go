package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("8.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var chars [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []rune
		text := scanner.Text()
		for _, char := range text {
			line = append(line, char)
		}
		chars = append(chars, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Part 1:", len(part1(&chars)))
	fmt.Println("Part 2:", len(part2(&chars)))
}

type Node struct {
	x, y int
}

func part1(chars *[][]rune) map[Node]struct{} {
	antinodes := make(map[Node]struct{})
	antennaMap := antennaMap(chars)

	for _, nodes := range antennaMap {
		if len(nodes) > 1 {
			for i, node := range nodes {
				for j, otherNode := range nodes {
					if i != j {
						dx := otherNode.x - node.x
						dy := otherNode.y - node.y
						if (node.x+(dx*2) >= 0 && node.x+(dx*2) < len((*chars)[0])) && (node.y+(dy*2) >= 0 && node.y+(dy*2) < len(*chars)) {
							antinodes[Node{node.x + dx*2, node.y + dy*2}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return antinodes
}

func part2(chars *[][]rune) map[Node]struct{} {
	antinodes := make(map[Node]struct{})
	antennaMap := antennaMap(chars)

	for _, nodes := range antennaMap {
		if len(nodes) > 1 {
			for i, node := range nodes {
				for j, otherNode := range nodes {
					if i != j {
						dx := otherNode.x - node.x
						dy := otherNode.y - node.y
						antinodes[node] = struct{}{}
						antinodes[otherNode] = struct{}{}
						for k := 2; ; k++ {
							nx := node.x + dx*k
							ny := node.y + dy*k
							if nx < 0 || nx >= len((*chars)[0]) || ny < 0 || ny >= len(*chars) {
								break
							}
							antinodes[Node{nx, ny}] = struct{}{}
						}
						for k := 1; ; k-- {
							nx := node.x + dx*k
							ny := node.y + dy*k
							if nx < 0 || nx >= len((*chars)[0]) || ny < 0 || ny >= len(*chars) {
								break
							}
							antinodes[Node{nx, ny}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return antinodes
}

func antennaMap(grid *[][]rune) map[rune][]Node {
	antennas := make(map[rune][]Node)
	for y, line := range *grid {
		for x, _ := range line {
			if (*grid)[y][x] != '.' {
				antennas[(*grid)[y][x]] = append(antennas[(*grid)[y][x]], Node{x, y})
			}
		}
	}
	return antennas
}
