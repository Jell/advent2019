package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
	"strings"
)

func main() {
	var prog = utils.ReadFile("../../inputs/day02.txt")
	prog = strings.TrimSuffix(prog, "\n")
	var tape = utils.StrsToInts(strings.Split(prog, ","))

	tape[1] = 12
	tape[2] = 2

	var position = 0

	var aStar int
	var bStar int
	var cStar int
Loop:
	for {
		switch tape[position] {
		case 1:
			aStar = tape[position+1]
			bStar = tape[position+2]
			cStar = tape[position+3]
			tape[cStar] = tape[aStar] + tape[bStar]
		case 2:
			aStar = tape[position+1]
			bStar = tape[position+2]
			cStar = tape[position+3]
			tape[cStar] = tape[aStar] * tape[bStar]
		case 99:
			break Loop
		default:
			panic("oh noes!")
		}
		position += 4
	}

	fmt.Println("Day 02 - Part 1:", tape[0])
}
