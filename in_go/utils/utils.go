package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Check ...
// panic if error
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadFile ...
func ReadFile(path string) string {
	raw, err := ioutil.ReadFile(path)
	Check(err)
	return string(raw)
}

// ReadLines ...
func ReadLines(path string) []string {
	var lines = strings.Split(ReadFile(path), "\n")
	return lines[:len(lines)-1]
}

// StrToInt ...
func StrToInt(input string) int {
	i, err := strconv.Atoi(input)
	Check(err)
	return i
}

// IntToStr ...
func IntToStr(input int) string {
	return strconv.FormatInt(int64(input), 10)
}

// StrsToInts ...
func StrsToInts(input []string) []int {
	var out []int
	for _, str := range input {
		out = append(out, StrToInt(str))
	}
	return out
}

// AbsInt ...
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// DoTimes ...
func DoTimes(n int, fn func()) {
	for i := 0; i < n; i++ {
		fn()
	}
}
