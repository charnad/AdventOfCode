package task5

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isNice(s string) bool {
	vowels := 0
	previous := ""
	hasTwice := false

	for _, s := range strings.Split(s, "") {
		switch s {
		case "a", "e", "i", "o", "u":
			vowels += 1

		case "b":
			if previous == "a" {
				return false
			}

		case "d":
			if previous == "c" {
				return false
			}

		case "q":
			if previous == "p" {
				return false
			}

		case "y":
			if previous == "x" {
				return false
			}
		}
		if previous == s {
			hasTwice = true
		}
		previous = s
	}

	return vowels > 2 && hasTwice
}

func isNice2(s string) bool {
	hasRepeated := false
	hasTwoPairs := false
	for i := 0; i < len(s); i++ {
		if i > 1 && s[i] == s[i-2] {
			hasRepeated = true
		}

		for j := 0; j < i-2; j++ {
			if s[j:j+2] == s[i-1:i+1] {
				hasTwoPairs = true
			}
		}
	}

	return hasRepeated && hasTwoPairs
}

func Solve() {
	file, err := os.Open("res/task5.txt")
	if err != nil {
		panic(err)
	}

	nice, nice2 := 0, 0

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if isNice(scanner.Text()) {
			nice++
		}
		if isNice2(scanner.Text()) {
			nice2++
		}
	}

	fmt.Println("Nice: ", nice, "  Nice 2", nice2)
}
