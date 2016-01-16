package task12

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strconv"
)

func Solve() {
	file, err := os.Open("res/task12.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	numbers, _ := regexp.Compile(`-?\d+`)
	matches := numbers.FindAllString(scanner.Text(), -1)

	total := 0
	for i := len(matches) - 1; i >= 0; i-- {
		number, _ := strconv.Atoi(matches[i])
		total += number
	}

	fmt.Println("Numbers total", total)
}