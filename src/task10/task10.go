package task10

import (
	"fmt"
	"strconv"
)

func step(s string) string {
	result := ""
	var prev rune = ' '
	count := 0

	for _, c := range s {
		if c == prev || prev == ' ' {
			prev = c
			count++
		} else {
			result = result + strconv.Itoa(count) + string(prev)
			count = 1
			prev = c
		}
	}

	result = result + strconv.Itoa(count) + string(prev)

	return result
}

func Solve() {
	input := "1113222113"
	for i := 0; i < 40; i++ {
		input = step(input)
	}

	fmt.Println("Length after 40", len(input))

	for i := 0; i < 10; i++ {
		input = step(input)
	}

	fmt.Println("Length after 50", len(input))
}
