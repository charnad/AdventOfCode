package task18

import (
	"bufio"
	"fmt"
	"os"
)

type field struct {
	cells [][]bool
}

func (f *field) step() {
	f2 := field{}

	for i, row := range f.cells {
		newRow := []bool{}
		for j, on := range row {
			neighbours := f.countNeighbours(i, j)
			if neighbours < 2 || neighbours > 3 {
				newRow = append(newRow, false)
			}

			if neighbours == 3 {
				newRow = append(newRow, true)
			}

			if neighbours == 2 {
				newRow = append(newRow, on)
			}
		}
		f2.cells = append(f2.cells, newRow)
	}

	f.cells = f2.cells
}

func (f field) countNeighbours(x, y int) int {
	result := 0

	length := len(f.cells) - 1

	if x < 0 || y < 0 || x > length || y > length {
		return 0
	}

	if y > 0 {
		if x > 0 && f.cells[x-1][y-1] {
			result++
		}

		if f.cells[x][y-1] {
			result++
		}

		if x < length && f.cells[x+1][y-1] {
			result++
		}
	}

	if x > 0 && f.cells[x-1][y] {
		result++
	}

	if x < length && f.cells[x+1][y] {
		result ++
	}

	if y < length {
		if x > 0 && f.cells[x-1][y+1] {
			result++
		}

		if f.cells[x][y+1] {
			result++
		}

		if x < length && f.cells[x+1][y+1] {
			result++
		}
	}

	return result
}

func (f field) countOn() int {
	result := 0

	for _, row := range f.cells {
		for _, on := range row {
			if on {
				result++
			}
		}
	}

	return result
}

func (f field) print() {
	for _, row := range f.cells {
		for _, on := range row {
			if on {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func Solve() {
	file, err := os.Open("res/task18.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	field := field{}

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []bool{}
		for _, v := range scanner.Text() {
			line = append(line, string(v) == "#")
		}
		field.cells = append(field.cells, line)
	}

	// Comment out part you don't need
	// Part 1
	for i := 0; i < 100; i++ {
		field.step()
	}

	fmt.Println(field.countOn())

	// Part 2
	size := len(field.cells)-1
	for i := 0; i < 100; i++ {
		field.step()
		field.cells[0][0] = true
		field.cells[0][size] = true
		field.cells[size][0] = true
		field.cells[size][size] = true
	}

	fmt.Println(field.countOn())
}
