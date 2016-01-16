package task6

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type lights struct {
	lights [1000][1000]int
}

func (l *lights) turnOn(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			l.lights[x][y]++
		}
	}
}

func (l *lights) turnOff(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if (l.lights[x][y] > 0) {
				l.lights[x][y]--
			}
		}
	}
}

func (l *lights) toggle(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			l.lights[x][y]+=2
		}
	}
}

func (l *lights) countOn() int {
	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if (l.lights[x][y] > 0) {
				count++
			}
		}
	}
	return count
}

func (l *lights) countBrightness() int {
	var count int = 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			count += l.lights[x][y]
		}
	}
	return count
}

func Solve() {
	file, err := os.Open("res/task6.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	l := lights{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands := strings.Fields(scanner.Text()[5:])

		coords := strings.Split(commands[1], ",")
		x1, _ := strconv.Atoi(coords[0])
		y1, _ := strconv.Atoi(coords[1])

		coords = strings.Split(commands[3], ",")
		x2, _ := strconv.Atoi(coords[0])
		y2, _ := strconv.Atoi(coords[1])

		switch commands[0] {
			case "e":
				l.toggle(x1, y1, x2, y2)
			case "off":
				l.turnOff(x1, y1, x2, y2)
			case "on":
				l.turnOn(x1, y1, x2, y2)
		}
	}

	fmt.Println("Lights on: ", l.countOn())
	fmt.Println("Brightness: ", l.countBrightness())
}