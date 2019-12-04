package main

import (
	"fmt"
	"github.com/Jell/advent2019/utils"
	"strings"
)

type password = [6]int

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func compareInts(left []int, right []int) string {
	var leftLen = len(left)
	var rightLen = len(right)
	var sharedLen = minInt(leftLen, rightLen)
	for i := 0; i < sharedLen; i++ {
		if left[i] < right[i] {
			return "LT"
		}
		if left[i] > right[i] {
			return "GT"
		}
	}
	if rightLen > leftLen {
		return "GT"
	}
	if rightLen < leftLen {
		return "LT"
	}
	return "EQ"
}
func lessThanInts(left []int, right []int) bool {
	return compareInts(left, right) == "LT"
}
func greaterThanInts(left []int, right []int) bool {
	return compareInts(left, right) == "GT"
}
func lessThanOrEqualToInts(left []int, right []int) bool {
	return !greaterThanInts(left, right)
}
func greaterThanOrEqualToInts(left []int, right []int) bool {
	return !lessThanInts(left, right)
}

func parsePassword(partStr string) [6]int {
	var password = password{}
	for i, c := range partStr {
		password[i] = utils.StrToInt(string(c))
	}
	return password
}

func partitionInts(in []int) [][]int {
	var out = [][]int{}
	var currentValue int
	var currentPartition []int
	for _, i := range in {
		if currentValue == i {
			currentPartition = append(currentPartition, i)
		} else {
			if currentPartition != nil {
				out = append(out, currentPartition)
			}
			currentPartition = []int{i}
		}
		currentValue = i
	}
	return append(out, currentPartition)
}

func hasDoubles(pass password) bool {
	var partitions = partitionInts(pass[:])
	for _, p := range partitions {
		if len(p) == 2 {
			return true
		}
	}
	return false
}

func main() {
	var input = utils.ReadFile("../../inputs/day04.txt")
	input = strings.TrimSuffix(input, "\n")
	var parts = strings.Split(input, "-")
	var startPassword = parsePassword(parts[0])
	var endPassword = parsePassword(parts[1])
	var testPassword password
	var validPasswords = []password{}

	var withinBounds = func(p password) bool {
		return greaterThanOrEqualToInts(testPassword[:], startPassword[:]) &&
			lessThanOrEqualToInts(testPassword[:], endPassword[:])
	}

	for p1 := 0; p1 < 10; p1++ {
		for p2 := p1; p2 < 10; p2++ {
			for p3 := p2; p3 < 10; p3++ {
				for p4 := p3; p4 < 10; p4++ {
					for p5 := p4; p5 < 10; p5++ {
						for p6 := p5; p6 < 10; p6++ {
							testPassword = password{p1, p2, p3, p4, p5, p6}
							if withinBounds(testPassword) && hasDoubles(testPassword) {
								validPasswords = append(validPasswords, testPassword)
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Day 04 - Part 2:", len(validPasswords))
}
