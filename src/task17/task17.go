package task17

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

const eggnog = 150

func sliceSum(slice []int) int {
	result := 0

	for _, val := range slice {
		result += val
	}

	return result
}

func combinations(current []int, src []int) [][]int {
	result := [][]int{}

	for key, value := range src {
		newCurrent := append(current, value)
		newTotal := sliceSum(newCurrent)

		if newTotal == eggnog {
			result = append(result, newCurrent)
			continue
		}

		if newTotal < eggnog {
			result = append(
				result,
				combinations(newCurrent, src[key+1:])...
			)
		}
	}

	return result
}

func Solve() {
	file, err := os.Open("res/task17.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	containers := []int{}

	// Read the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		container, _ := strconv.Atoi(scanner.Text())
		containers = append(containers, container)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(containers)))

	comb := combinations([]int{}, containers)
	fmt.Println("Part 1 Result: ", len(comb))

	min := 100
	for _, slice := range comb {
		if min > len(slice) {
			min = len(slice)
		}
	}

	result := 0
	for _, slice := range comb {
		if min == len(slice) {
			result++
		}
	}

	fmt.Println("Part 2 Result: ", result)
}
