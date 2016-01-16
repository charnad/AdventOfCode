package task2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	width  int
	height int
	depth  int
}

func (b Box) smallestSides() (int, int) {
	min, max := 0, 0

	if b.width > b.height {
		min, max = b.height, b.width
	} else {
		min, max = b.width, b.height
	}

	if b.depth > max {
		return b.height, b.width
	}

	if b.depth < min {
		return b.depth, min
	}

	return b.depth, min
}

func (b Box) wrappingPaper() int {
	sideA, sideB := b.smallestSides()
	return 2*b.width*b.height + 2*b.width*b.depth + 2*b.height*b.depth + sideA*sideB
}

func (b Box) ribbon() int {
	sideA, sideB := b.smallestSides()
	return sideA*2 + sideB*2 + b.height*b.width*b.depth
}

func fromString(str string) Box {
	b := Box{}

	var dimensions []int
	for _, i := range strings.Split(str, "x") {
		dim, _ := strconv.Atoi(i)
		dimensions = append(dimensions, dim)
	}

	b.height = dimensions[0]
	b.width = dimensions[1]
	b.depth = dimensions[2]

	return b
}

func Solve() {
	totalPaper := 0
	totalRibbon := 0

	file, err := os.Open("res/task2.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		box := fromString(scanner.Text())
		totalPaper += box.wrappingPaper()
		totalRibbon += box.ribbon()
	}

	fmt.Println("Total paper needed: ", totalPaper, " Total ribbon: ", totalRibbon)
}
