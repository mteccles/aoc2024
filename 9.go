package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("9.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %s", err)
	}

	fs := BuildFileSystem(line)

	// Part 1
	for i := 0; i < fs.Length(); i++ {
		if fs.GetData(i) < 0 {
			fs.RemoveTrailingSpace()
			fs.SetData(i, fs.GetData(fs.Length()-1))
			fs.RemoveData(fs.Length() - 1)
		}
	}

	part1 := 0
	for i := 0; i < fs.Length(); i++ {
		part1 += i * fs.GetData(i)
	}
	fmt.Println("Part 1:", part1)

	// Part 2
	fs = BuildFileSystem(line)
	var currentFId int
	var currentFLen int
	h := fs.Length() - 1
	for i := fs.Length() - 1; i >= 0; i-- {
		fId := fs.GetData(i)
		if i == h {
			currentFId = fId
			currentFLen = 1
			continue
		}
		if i == 0 {
			break
		}
		if fId == -1 {
			currentFLen = 0
			currentFId = fId
			continue
		} else {
			currentFId = fId
		}
		if fId == fs.GetData(i-1) {
			currentFLen++
			continue
		}
		currentFLen++
		freeSpace := 0
		startIndex := -1
		for j := 0; j < fs.Length(); j++ {
			if j >= i {
				break
			}
			if fs.GetData(j) == -1 {
				if startIndex == -1 {
					startIndex = j
					freeSpace = 0
				}
				freeSpace++

				if freeSpace == currentFLen {
					for k := 0; k < currentFLen; k++ {
						fs.SetData(startIndex+k, currentFId)
						fs.SetData(i+k, -1)
					}
					break
				}
			} else {
				startIndex = -1
				freeSpace = 0
			}
		}
		currentFLen = 0

	}
	part2 := 0
	for i := 0; i < fs.Length(); i++ {
		if fs.GetData(i) != -1 {
			part2 += i * fs.GetData(i)
		}
	}
	fmt.Println("Part 2:", part2)

}

func BuildFileSystem(data string) FileSystem {
	pos := 0
	fs := FileSystem{}
	fs.Initialize()
	for i, char := range data {
		blockSize, _ := strconv.Atoi(string(char))
		if i%2 == 0 {
			// We have a file block.
			for k := 0; k < blockSize; k++ {
				fs.SetData(pos+k, i/2)
			}
			pos += blockSize
		} else {
			// We have empty space.
			for k := 0; k < blockSize; k++ {
				fs.SetData(pos+k, -1)
			}
			pos += blockSize
		}
	}
	return fs
}

type FileSystem struct {
	data map[int]int
}

func (fs *FileSystem) Initialize() {
	fs.data = make(map[int]int)
}

func (fs *FileSystem) SetData(addr int, data int) {
	fs.data[addr] = data
}

func (fs *FileSystem) GetData(addr int) int {
	return fs.data[addr]
}

func (fs *FileSystem) RemoveData(addr int) {
	delete(fs.data, addr)
}

func (fs *FileSystem) Length() int {
	return len(fs.data)
}

func (fs *FileSystem) RemoveTrailingSpace() {
	for fs.GetData(fs.Length()-1) == -1 {
		fs.RemoveData(fs.Length() - 1)
	}
}

func (fs *FileSystem) Print() {
	for i := 0; i < len(fs.data); i++ {
		if fs.data[i] == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(fs.data[i])
		}
	}
	fmt.Println()
}
