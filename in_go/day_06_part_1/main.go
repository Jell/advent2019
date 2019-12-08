package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
	"strings"
)

type planet struct {
	name     string
	parent   *planet
	children []*planet
}

func getOrbits(p planet) int {
	if p.parent == nil {
		return 0
	}
	return 1 + getOrbits(*p.parent)
}

func main() {
	var lines = utils.ReadLines("../../inputs/day06.txt")
	var planets = map[string]*planet{}

	for _, line := range lines {
		var parts = strings.Split(line, ")")
		var name1 = parts[0]
		var name2 = parts[1]
		var planet1 = planets[name1]
		if planet1 == nil {
			planet1 = &planet{
				name:     name1,
				children: []*planet{},
			}
		}

		var planet2 = planets[name2]
		if planet2 == nil {
			planet2 = &planet{
				name:     name2,
				children: []*planet{},
			}
		}

		planet1.children = append(planet1.children, planet2)
		planet2.parent = planet1

		planets[name1] = planet1
		planets[name2] = planet2
	}

	var orbits = 0
	for _, planet := range planets {
		orbits += getOrbits(*planet)
	}

	fmt.Println("Day 06 - Part 1:", orbits)
}
