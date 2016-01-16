package task14

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

type reindeer struct {
	name  string
	speed int
	time  int
	rest  int
}

func (r reindeer) distanceAfter(seconds int) int {
	distance := 0
	for seconds > 0 {
		for i := r.time; i > 0 && seconds > 0; i-- {
			distance += r.speed
			seconds--
		}
		seconds -= r.rest
	}
	return distance
}



func Solve() {
	file, err := os.Open("res/task14.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reindeers := []reindeer{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		r := reindeer{}
		r.name = fields[0]
		r.speed, _ = strconv.Atoi(fields[3])
		r.time, _ = strconv.Atoi(fields[6])
		r.rest, _ = strconv.Atoi(fields[13])
		reindeers = append(reindeers, r)
	}

	for _, r := range reindeers {
		fmt.Println(r.name, "traveled", r.distanceAfter(2503))
	}

	score := map[string]int{}
	for _, r := range reindeers {
		score[r.name] = 0
	}

	for time := 1; time < 2503; time++ {
		travelled := map[string]int{}
		for _, r := range reindeers {
			travelled[r.name] = r.distanceAfter(time)
		}

		max := 0
		for _, went := range travelled {
			if went > max {
				max = went
			}
		}

		for name, went := range travelled {
			if went == max {
				score[name]++
			}
		}
	}

	fmt.Println(score)
}

googleapps
mail@viktoras.de
(NkH1ai!a@3T