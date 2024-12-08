package util

import (
	"os"
	"strconv"
	"strings"
)

// Newline character
var N = "\n"

var res = "resources"

func GetInput(filename string) []string {
	file := ReadFile(filename)
	file = file[:len(file)-1]
	return SplitFile(file, "\n")
}

func ReadFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func SplitFile(file string, predicate string) []string {
	return strings.Split(file, predicate)
}

func GetInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}
