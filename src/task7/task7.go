package task7

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type wire struct {
	lefts  string // left value source
	left   uint16 // left int value

	op     string // operation

	rights string // right value source
	right  uint16 // right int value
}

func getSignal(wires map[string]wire, name string) uint16 {
	if _, ok := wires[name]; !ok {
		return 0
	}

	wire := wires[name]
	var left, right, result uint16 = 0, 0, 0

	if (wire.lefts != "") {
		left = getSignal(wires, wire.lefts)
		wire.left = left
		wire.lefts = ""
	} else {
		left = wire.left
	}

	if (wire.rights != "") {
		right = getSignal(wires, wire.rights)
		wire.right = right
		wire.rights = ""
	} else {
		right = wire.right
	}

	switch wire.op {
	case "AND":
		result = left & right
	case "OR":
		result = left | right
	case "LSHIFT":
		result = left << right
	case "RSHIFT":
		result = left >> right
	case "NOT":
		result = ^left
	case "NOP":
		result = left
	case "ASG":
		result = left
	}

	wires[name] = wire

	return uint16(result)
}

func strToUint16(s string) uint16 {
	i, _ := strconv.Atoi(s)
	return uint16(i)
}

func isIntString(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func fromString(s string) (string, wire) {
	parts := strings.Split(s, " -> ")
	operands := strings.Fields(parts[0])

	w := wire{}

	switch len(operands) {
		case 3:
			if isIntString(operands[0]) {
				w.lefts = ""
				w.left = strToUint16(operands[0])
			} else {
				w.lefts = operands[0]
				w.left = 0
			}
			w.op   = operands[1]
			if isIntString(operands[2]) {
				w.rights = ""
				w.right = strToUint16(operands[2])
			} else {
				w.rights = operands[2]
				w.right = 0
			}

		case 2:
			if isIntString(operands[1]) {
				w.lefts = ""
				w.left = strToUint16(operands[1])
			} else {
				w.lefts = operands[1]
				w.left = 0
			}
			w.op = operands[0]

		case 1:
			if i, err := strconv.Atoi(operands[0]); err == nil {
				w.left = uint16(i)
				w.lefts = ""
			} else {
				w.lefts = operands[0]
				w.left = 0
			}
			w.op = "NOP"
	}

	return parts[1], w
}

func Solve() {
	file, err := os.Open("res/task7.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	wires := map[string]wire{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name, w := fromString(scanner.Text())
		wires[name] = w
	}

//	fmt.Println(getSignal(wires, "a"))

	wireb := wires["b"]
	wireb.left = 16076
	wires["b"] = wireb

	fmt.Println("Signal at wire a is", getSignal(wires, "a"))
}