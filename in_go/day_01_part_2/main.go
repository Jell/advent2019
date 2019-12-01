package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
)

func main() {
	var lines = utils.ReadLines("../../inputs/day01.txt")
	var fuel = 0

	for _, line := range lines {
		var mass = utils.StrToInt(line)
		var extraFuelMass = mass
		for {
			extraFuelMass = (extraFuelMass / 3) - 2
			if extraFuelMass > 0 {
				fuel += extraFuelMass
			} else {
				break
			}
		}
	}

	fmt.Println("Day 01 - Part 2:", fuel)
}
