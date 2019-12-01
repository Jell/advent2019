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
		fuel += (mass / 3) - 2
	}

	fmt.Println("Day 01 - Part 1:", fuel)
}
