package task13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inSlice(s string, sl []string) bool {
	for _, item := range sl {
		if item == s {
			return true
		}
	}

	return false
}

func sliceDiff(a []string, b []string) []string {
	result := []string{}
	for _, itemA := range a {
		if !inSlice(itemA, b) {
			result = append(result, itemA)
		}
	}
	return result
}

func permutations(persons []string) [][]string {
	permutations := [][]string{}
	permutations = append(permutations, []string{})
	for i := len(persons); i > 0; i-- {
		newPermutations := [][]string{}

		for _, permutation := range permutations {
			leftToSeat := sliceDiff(persons, permutation)
			for _, item := range leftToSeat {
				newPermutations = append(newPermutations, append(append([]string{}, permutation...), item))
			}
		}

		if len(newPermutations) > 0 {
			permutations = append(newPermutations)
		}
	}

	return permutations
}

func calculateTotalChange(seatings [][]string, changes map[string]int) int {
	total := 0
	max := 0
	for _, permutation := range seatings {
		permutation = append(permutation, permutation[0])
		total = 0
		for i := 0; i < len(permutation)-1; i++ {
			total += changes[permutation[i]+permutation[i+1]] + changes[permutation[i+1]+permutation[i]]
		}

		if total > max {
			max = total
		}
	}

	return max
}

func Solve() {
	file, err := os.Open("res/task13.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	persons := []string{}
	changes := map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(strings.Trim(scanner.Text(), "."))
		person1 := parts[0]
		person2 := parts[10]
		change, _ := strconv.Atoi(parts[3])
		if parts[2] == "lose" {
			change *= -1
		}
		if !inSlice(person1, persons) {
			persons = append(persons, person1)
		}
		changes[person1+person2] = change
	}

	fmt.Println("Total change", calculateTotalChange(permutations(persons), changes))
	fmt.Println("Total change including me", calculateTotalChange(permutations(append(persons, "me")), changes))
}
