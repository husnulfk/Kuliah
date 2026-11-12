package main

import (
	"fmt"
)

func main() {
	var kill, death, damage int
	var spikePlanted, unggulDuel, damageTinggi, roundImpact bool

	fmt.Scan(&kill, &death, &damage, &spikePlanted)

	unggulDuel = kill > death
	damageTinggi = damage >= 2000
	roundImpact = (unggulDuel == true && damageTinggi == true) || spikePlanted == true

	fmt.Println(unggulDuel, damageTinggi, roundImpact)

}
