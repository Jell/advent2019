package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
	"sort"
	"strings"
)

type move struct {
	direction string
	amplitude int
}

type position struct {
	x int
	y int
}

type stepCounter struct {
	wire1 int
	wire2 int
}

func parseMoves(in string) []move {
	var out []move
	for _, s := range strings.Split(in, ",") {
		out = append(out, move{
			direction: s[0:1],
			amplitude: utils.StrToInt(s[1:]),
		})
	}
	return out
}

func manhattan(pos position) int {
	return utils.AbsInt(pos.x) + utils.AbsInt(pos.y)
}

func main() {
	var input = utils.ReadLines("../../inputs/day03.txt")

	var wire1 = parseMoves(input[0])
	var wire2 = parseMoves(input[1])

	var crossPositions = map[position]stepCounter{}
	var visited1 = map[position]bool{}
	var visited2 = map[position]bool{}

	var pos position
	var steps int

	var moveMaker = func(dir string) func() {
		switch dir {
		case "U":
			return func() {
				pos = position{pos.x, pos.y + 1}
			}
		case "D":
			return func() {
				pos = position{pos.x, pos.y - 1}
			}
		case "R":
			return func() {
				pos = position{pos.x + 1, pos.y}
			}
		case "L":
			return func() {
				pos = position{pos.x - 1, pos.y}
			}
		}
		panic("oops")
	}

	pos = position{x: 0, y: 0}
	for _, m := range wire1 {
		move := moveMaker(m.direction)
		utils.DoTimes(m.amplitude, func() {
			move()
			visited1[pos] = true
		})
	}

	pos = position{x: 0, y: 0}
	steps = 0
	for _, m := range wire2 {
		move := moveMaker(m.direction)
		utils.DoTimes(m.amplitude, func() {
			move()
			visited2[pos] = true
			steps++
			if visited1[pos] {
				if _, ok := crossPositions[pos]; !ok {
					crossPositions[pos] = stepCounter{wire2: steps}
				}

			}
		})
	}

	pos = position{x: 0, y: 0}
	steps = 0
	for _, m := range wire1 {
		move := moveMaker(m.direction)
		utils.DoTimes(m.amplitude, func() {
			move()
			steps++
			if visited2[pos] {
				if counter := crossPositions[pos]; counter.wire1 == 0 {
					crossPositions[pos] = stepCounter{
						wire1: steps,
						wire2: counter.wire2,
					}
				}

			}
		})
	}

	var crossDistances = []int{}
	for _, cp := range crossPositions {
		crossDistances = append(crossDistances, cp.wire1+cp.wire2)
	}

	sort.Ints(crossDistances)

	var shortest = crossDistances[0]

	fmt.Println("Day 03 - Part 2:", shortest)
}
