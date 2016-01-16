package task1

import (
	"fmt"
	"io/ioutil"
)

func Solve() {
	data, err := ioutil.ReadFile("res/task1.txt")
	if err != nil {
		panic(err)
	}

	floor := 0
	position := 0

	for key, symbol := range data {
		switch string(symbol) {
		case ")":
			floor -= 1
		case "(":
			floor += 1
		}

		if floor < 0 && position == 0 {
			position = key + 1
		}
	}

	fmt.Println("Floor:", floor, ". Went underground at", position)
}
