package task15

import (
	"fmt"
)

/**
 * Sprinkles: capacity 2, durability 0, flavor -2, texture 0, calories 3
 * Butterscotch: capacity 0, durability 5, flavor -3, texture 0, calories 3
 * Chocolate: capacity 0, durability 0, flavor 5, texture -1, calories 8
 * Candy: capacity 0, durability -1, flavor 0, texture 5, calories 8
 */

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var sprinkles ingredient = ingredient{2, 0, -2, 0, 3}
var butterscotch ingredient = ingredient{0, 5, -3, 0, 3}
var chocolate ingredient = ingredient{0, 0, 5, -1, 8}
var candy ingredient = ingredient{0, -1, 0, 5, 8}

const teaspoons = 100

type recepie struct {
	sprinkles     int
	butterschotch int
	chocolate     int
	candy         int
}

func (r recepie) score() int {
	capacity := r.sprinkles * sprinkles.capacity
	capacity += r.butterschotch * butterscotch.capacity
	capacity += r.chocolate * chocolate.capacity
	capacity += r.candy * candy.capacity
	if capacity < 0 {
		capacity = 0
	}

	durability := r.sprinkles * sprinkles.durability
	durability += r.butterschotch * butterscotch.durability
	durability += r.chocolate * chocolate.durability
	durability += r.candy * candy.durability
	if durability < 0 {
		durability = 0
	}

	flavor := r.sprinkles * sprinkles.flavor
	flavor += r.butterschotch * butterscotch.flavor
	flavor += r.chocolate * chocolate.flavor
	flavor += r.candy * candy.flavor
	if flavor < 0 {
		flavor = 0
	}

	texture := r.sprinkles * sprinkles.texture
	texture += r.butterschotch * butterscotch.texture
	texture += r.chocolate * chocolate.texture
	texture += r.candy * candy.texture
	if texture < 0 {
		texture = 0
	}

	return capacity * durability * flavor * texture
}

func (r recepie) calories() int {
	return r.sprinkles*sprinkles.calories +
		r.butterschotch*butterscotch.calories +
		r.chocolate*chocolate.calories +
		r.candy*candy.calories
}

func Solve() {
	max := 0
	maxWithCalories := 0

	for sp := 0; sp <= 100; sp++ {
		for bu := 0; bu <= 100; bu++ {
			for ch := 0; ch <= 100; ch++ {
				for ca := 0; ca <= 100; ca++ {
					if sp+bu+ch+ca != 100 {
						continue
					}

					r := recepie{sp, bu, ch, ca}
					score := r.score()
					if score > max {
						max = score
					}

					if r.calories() != 500 {
						continue
					}

					if score > maxWithCalories {
						maxWithCalories = score
					}
				}
			}
		}
	}

	fmt.Println("Max score is", max, "max score for 500 calories is", maxWithCalories)
}
