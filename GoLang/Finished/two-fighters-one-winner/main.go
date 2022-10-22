package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func main() {
	fighterOne := Fighter{
		Name:            "Harald",
		Health:          20,
		DamagePerAttack: 5,
	}
	fighterTwo := Fighter{
		Name:            "Harry",
		Health:          5,
		DamagePerAttack: 4,
	}

	fmt.Println(DeclareWinner(fighterOne, fighterTwo, "Harry"))
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	if firstAttacker == fighter1.Name {
		fighter2.Health -= fighter1.DamagePerAttack
		if fighter2.Health <= 0 {
			return fighter1.Name
		}
		fighter1.Health -= fighter2.DamagePerAttack
		if fighter1.Health <= 0 {
			return fighter2.Name
		}
		return DeclareWinner(fighter1, fighter2, fighter1.Name)
	} else {
		fighter1.Health -= fighter2.DamagePerAttack
		if fighter1.Health <= 0 {
			return fighter2.Name
		}
		fighter2.Health -= fighter1.DamagePerAttack
		if fighter2.Health <= 0 {
			return fighter1.Name
		}
		return DeclareWinner(fighter1, fighter2, fighter2.Name)
	}
}
