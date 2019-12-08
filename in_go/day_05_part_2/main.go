package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
	"strings"
)

type operation struct {
	code      int
	arity     int
	immediate []bool
}

func parseOp(in int) operation {
	var arity int
	var immediates []bool
	var op = in % 100
	in -= op
	var immediate1 = in % 1000
	in -= immediate1
	var immediate2 = in % 10000
	in -= immediate2

	switch op {
	case 99:
		arity = 0
		immediates = []bool{}
	case 1, 2, 7, 8:
		arity = 3
		immediates = []bool{
			immediate1 > 0,
			immediate2 > 0,
			true,
		}
	case 3:
		arity = 1
		immediates = []bool{
			true,
		}
	case 4:
		arity = 1
		immediates = []bool{
			immediate1 > 0,
		}
	case 5, 6:
		arity = 2
		immediates = []bool{
			immediate1 > 0,
			immediate2 > 0,
		}
	}
	return operation{
		code:      op,
		arity:     arity,
		immediate: immediates,
	}
}

func main() {
	var prog = utils.ReadFile("../../inputs/day05.txt")
	prog = strings.TrimSuffix(prog, "\n")
	var tape = utils.StrsToInts(strings.Split(prog, ","))

	var position = 0

	var input = 5
	var outputs = []int{}
	var op operation
	var args []int

Loop:
	for {
		op = parseOp(tape[position])
		args = []int{}
		for a := 1; a <= op.arity; a++ {
			var val int
			if op.immediate[a-1] {
				val = tape[position+a]
			} else {
				val = tape[tape[position+a]]
			}
			args = append(args, val)
		}

		switch op.code {
		case 1:
			tape[args[2]] = args[0] + args[1]
		case 2:
			tape[args[2]] = args[0] * args[1]
		case 3:
			tape[args[0]] = input
		case 4:
			outputs = append(outputs, args[0])
		case 5:
			if args[0] != 0 {
				position = args[1]
				continue
			}
		case 6:
			if args[0] == 0 {
				position = args[1]
				continue
			}
		case 7:
			if args[0] < args[1] {
				tape[args[2]] = 1
			} else {
				tape[args[2]] = 0
			}
		case 8:
			if args[0] == args[1] {
				tape[args[2]] = 1
			} else {
				tape[args[2]] = 0
			}
		case 99:
			break Loop
		default:
			panic("oh noes!")
		}
		position += (op.arity + 1)
	}

	fmt.Println("Day 05 - Part 2:", outputs[len(outputs)-1])
}
