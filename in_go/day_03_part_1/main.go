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

	var crossDistances = []int{}
	var visited = map[position]bool{}

	var pos position

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
	visited[pos] = true
	for _, m := range wire1 {
		move := moveMaker(m.direction)
		utils.DoTimes(m.amplitude, func() {
			move()
			visited[pos] = true
		})
	}

	pos = position{x: 0, y: 0}
	for _, m := range wire2 {
		move := moveMaker(m.direction)
		utils.DoTimes(m.amplitude, func() {
			move()
			if visited[pos] {
				crossDistances = append(crossDistances, manhattan(pos))
			}
		})
	}

	sort.Ints(crossDistances)

	var shortest = crossDistances[0]

	fmt.Println("Day 03 - Part 1:", shortest)
}
