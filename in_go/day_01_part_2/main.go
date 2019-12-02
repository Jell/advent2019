package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
)

func massToFuel(mass int) int {
	var fuel = (mass / 3) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + massToFuel(fuel)
}

func main() {
	var lines = utils.ReadLines("../../inputs/day01.txt")
	var fuel = 0

	for _, line := range lines {
		var mass = utils.StrToInt(line)
		fuel += massToFuel(mass)
	}

	fmt.Println("Day 01 - Part 2:", fuel)
}
