package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
	"strings"
)

func try(tape []int, target int) bool {
	var position = 0

	var aStar int
	var bStar int
	var cStar int
Loop:
	for {
		if tape[0] >= target {
			break Loop
		}
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

	return tape[0] == target
}

func main() {
	var prog = utils.ReadFile("../../inputs/day02.txt")
	prog = strings.TrimSuffix(prog, "\n")
	var tape = utils.StrsToInts(strings.Split(prog, ","))
	var target = 19690720
	var noun = 0
	var verb = 0

Loop:
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			var tapeTmp = make([]int, len(tape))
			copy(tapeTmp, tape)
			tapeTmp[1] = noun
			tapeTmp[2] = verb
			if try(tapeTmp, target) {
				break Loop
			}
		}
	}

	fmt.Println("Day 02 - Part 2:", noun*100+verb)
}
