package task9

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
	"sort"
)

func inSlice(value string, slice []string) bool {
	for _, node := range slice {
		if value == node {
			return true
		}
	}

	return false
}

func cloneSlice(slice []string) []string {
	return append([]string{}, slice...)
}

func findRoute(path []string, routes map[string]map[string]int, totalDistance int) map[string]int {
	results := map[string]int{}

	if len(path) == 0 {
		for start, _:= range routes {
			found := findRoute(append(cloneSlice(path), start), routes, 0)
			for k, v := range found {
				results[k] = v
			}
		}

		return results
	}

	source := path[len(path)-1]
	for dest, dist := range routes[source] {
		if !inSlice(dest, path) {
			for k, v := range findRoute(append(cloneSlice(path), dest), routes, totalDistance + dist) {
				results[k] = v
			}
		}
	}

	// Every destination has been visited
	if len(path) == len(routes) {
		results[strings.Join(path, " -> ")] = totalDistance
	}

	return results
}

func Solve() {
	file, err := os.Open("res/task9.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	routes := map[string]map[string]int{};

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		route := scanner.Text()
		parts := strings.Split(route, " to ")
		from := parts[0]
		parts2 := strings.Split(parts[1], " = ")
		to := parts2[0]
		distance, _ := strconv.Atoi(parts2[1])

		if routes[from] == nil {
			routes[from] = map[string]int{to: distance}
		} else {
			routes[from][to] = distance
		}

		if routes[to] == nil {
			routes[to] = map[string]int{from: distance}
		} else {
			routes[to][from] = distance
		}
	}

	distances := []int{}
	for _, dist := range findRoute([]string{}, routes, 0) {
		distances = append(distances, dist)
	}

	sort.Ints(distances)

	for route, dist := range findRoute([]string{}, routes, 0) {
		if dist == distances[0] {
			fmt.Println("Shortest route", route, dist)
		}

		if dist == distances[len(distances)-1] {
			fmt.Println("Shortest route", route, dist)
		}
	}
}