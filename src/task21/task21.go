package task21

import (
	"fmt"
)

type unit struct {
	hp int
	damage int
	armor int
	equipment
}

func (u unit) getDamage() int {
	return u.damage + u.equipment.getDamage();
}

func (u unit) getArmor() int {
	return u.armor + u.equipment.getArmor();
}

func (this *unit) attack(that *unit) {
	that.hp -= max(1, this.getDamage() - that.getArmor())
}

type item struct {
	name string
	cost int
	damage int
	armor int
}

type equipment struct {
	weapon item
	armor item
	ring1 item
	ring2 item
}

func (e equipment) getCost() int {
	return e.weapon.cost + e.armor.cost + e.ring1.cost + e.ring2.cost
}

func (e equipment) getDamage() int {
	return e.weapon.damage + e.armor.damage + e.ring1.damage + e.ring2.damage
}

func (e equipment) getArmor() int {
	return e.weapon.armor + e.armor.armor + e.ring1.armor + e.ring2.armor
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fight(a, b unit) bool {
	current, other := &a, &b

	for a.hp > 0 && b.hp > 0 {
		current.attack(other)
		current, other = other,current
	}

	return a.hp > 0 && b.hp <= 0
}

func Solve() {

	boss := unit{100, 8, 2, equipment{}}
	hero := unit{100, 0, 0, equipment{}}

	noitem := item{}

	weapons := []item{
		item{"dagger", 8, 4, 0},
		item{"shortsword", 10, 5, 0},
		item{"warhammer", 25, 6, 0},
		item{"longsword", 40, 7, 0},
		item{"greataxe", 74, 8, 0},
	}

	armors := []item{
		item{"leather", 13, 0, 1},
		item{"chainmail", 31, 0, 2},
		item{"splintmail", 53, 0, 3},
		item{"bandedmail", 75, 0, 4},
		item{"platemail", 102, 0, 5},
	}

	rings := []item{
		item{"damage +1", 25, 1, 0},
		item{"damage +2", 50, 2, 0},
		item{"damage +3", 100, 3, 0},
		item{"defense +1", 20, 0, 1},
		item{"defense +2", 40, 0, 2},
		item{"defense +3", 80, 0, 3},
	}

	equipments := []equipment{}

	for _, weapon := range weapons {
		equipments = append(equipments, equipment{weapon, noitem, noitem, noitem})
	}

	for _, eq := range equipments {
		for _, armor := range armors {
			equipments = append(equipments, equipment{eq.weapon, armor, noitem, noitem})
		}
	}

	for _, eq := range equipments {
		for _, ring := range rings {
			equipments = append(equipments, equipment{eq.weapon, eq.armor, ring, noitem})
		}
	}

	for _, eq := range equipments {
		for _, ring := range rings {
			if ring != eq.ring1 {
				equipments = append(equipments, equipment{eq.weapon, eq.armor, eq.ring1, ring})
			}
		}
	}

	minGold := 0
	var minEq equipment
	for _, eq := range equipments {
		hero.equipment = eq
		if fight(hero, boss) && (minGold == 0 || eq.getCost() < minGold) {
			minGold = eq.getCost()
			minEq = eq
		}
	}

	fmt.Println("Least amount of gold to win is", minGold, minEq)

	maxGold := 0
	var maxEq equipment
	for _, eq := range equipments {
		hero.equipment = eq
		if !fight(hero, boss) && eq.getCost() > maxGold {
			maxGold = eq.getCost()
			maxEq = eq
		}
	}

	fmt.Println("Most amount of gold to lose is", maxGold, maxEq)
}
