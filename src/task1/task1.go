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

	var floor int16 = 0

	for _, symbol := range data {
		fmt.Println(floor)
		switch string(symbol) {
		case ")":
			floor -= 1
		case "(":
			floor += 1
		}
	}

	fmt.Println("Floor: ", floor)
}
