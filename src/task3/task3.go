package task3

import (
	"fmt"
	"io/ioutil"
)

func Solve() {
	data, err := ioutil.ReadFile("res/task3.txt")
	if err != nil {
		panic(err)
	}

	cx, cy, ox, oy := 0, 0, 0, 0

	var houses = make(map[int]map[int]bool)

	if houses[cx] == nil {
		houses[cx] = make(map[int]bool)
	}
	houses[cx][cy] = true

	for _, symbol := range data {
		switch string(symbol) {
		case "^":
			cy += 1
		case "v":
			cy -= 1
		case "<":
			cx -= 1
		case ">":
			cx += 1
		}

		if houses[cx] == nil {
			houses[cx] = make(map[int]bool)
		}
		houses[cx][cy] = true

		cx, cy, ox, oy = ox, oy, cx, cy
	}

	total := 0
	for _, i := range houses {
		total += len(i)
	}
	fmt.Println("Houses: ", total)
}
