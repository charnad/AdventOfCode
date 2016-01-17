package task16

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type aunt struct {
	number      int
	children    int
	cats        int
	samoyeds    int
	pomeranians int
	akitas      int
	vizslas     int
	goldfish    int
	trees       int
	cars        int
	perfumes    int
}

func fromString(str string) aunt {
	str = strings.Replace(str, ":", "", -1)
	str = strings.Replace(str, ",", "", -1)
	fields := strings.Fields(str)

	a := aunt{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	for i := 0; i < 7; i += 2 {
		amount, _ := strconv.Atoi(fields[i+1])
		switch fields[i] {
		case "Sue":
			a.number = amount
		case "children":
			a.children = amount
		case "cats":
			a.cats = amount
		case "samoyeds":
			a.samoyeds = amount
		case "pomeranians":
			a.pomeranians = amount
		case "akitas":
			a.akitas = amount
		case "vizslas":
			a.vizslas = amount
		case "goldfish":
			a.goldfish = amount
		case "trees":
			a.trees = amount
		case "cars":
			a.cars = amount
		case "perfumes":
			a.perfumes = amount
		}
	}

	return a
}

func (a aunt) check(children int, cats int, samoyeds int, pomeranians int, akitas int, vizslas int, goldfish int, trees int, cars int, perfumes int) bool {
	if a.children >= 0 && a.children != children {
		return false
	}

	if a.cats >= 0 && a.cats != cats {
		return false
	}

	if a.samoyeds >= 0 && a.samoyeds != samoyeds {
		return false
	}

	if a.pomeranians >= 0 && a.pomeranians != pomeranians {
		return false
	}

	if a.akitas >= 0 && a.akitas != akitas {
		return false
	}

	if a.vizslas >= 0 && a.vizslas != vizslas {
		return false
	}

	if a.goldfish >= 0 && a.goldfish != goldfish {
		return false
	}

	if a.trees >= 0 && a.trees != trees {
		return false
	}

	if a.cars >= 0 && a.cars != cars {
		return false
	}

	if a.perfumes >= 0 && a.perfumes != perfumes {
		return false
	}

	return true
}

func (a aunt) check2(children int, cats int, samoyeds int, pomeranians int, akitas int, vizslas int, goldfish int, trees int, cars int, perfumes int) bool {
	if a.children >= 0 && a.children != children {
		return false
	}

	if a.cats >= 0 && a.cats <= cats {
		return false
	}

	if a.samoyeds >= 0 && a.samoyeds != samoyeds {
		return false
	}

	if a.pomeranians >= 0 && a.pomeranians >= pomeranians {
		return false
	}

	if a.akitas >= 0 && a.akitas != akitas {
		return false
	}

	if a.vizslas >= 0 && a.vizslas != vizslas {
		return false
	}

	if a.goldfish >= 0 && a.goldfish >= goldfish {
		return false
	}

	if a.trees >= 0 && a.trees <= trees {
		return false
	}

	if a.cars >= 0 && a.cars != cars {
		return false
	}

	if a.perfumes >= 0 && a.perfumes != perfumes {
		return false
	}

	return true
}

func Solve() {

	file, err := os.Open("res/task16.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		aunt := fromString(line)
		if aunt.check(3, 7, 2, 3, 0, 0, 5, 3, 2, 1) {
			fmt.Printf("Aunt #%d first check\n", aunt.number)
		}

		if aunt.check2(3, 7, 2, 3, 0, 0, 5, 3, 2, 1) {
			fmt.Printf("Aunt #%d second check\n", aunt.number)
		}
	}
}
