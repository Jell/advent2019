package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
)

func part1(lines []string) {
	var fuel = 0

	for _, line := range lines {
		var mass = utils.StrToInt(line)
		fuel += (mass / 3) - 2
	}

	fmt.Println("Day 01 - Part 1:", fuel)
}

func part2(lines []string) {
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

func main() {
	var lines = utils.ReadLines("../../inputs/day01.txt")
	part1(lines)
	part2(lines)
}
