package day2

import (
	"adventofcode2024/util"
	"fmt"
	"strconv"
	"strings"
)

var part1TestInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
5 7 3 8 9`

func TestPart1() int {
	return Part1(util.SplitFile(part1TestInput, util.N))
}

func TestPart2() int {
	return Part2(util.SplitFile(part1TestInput, util.N))
	// return Part2([]string{"5 7 3 8 9"})
}

func safeAccessNextValue(i int, l []string) int {
	length := len(l)
	if i == (length - 1) {
		return -1
	} else {
		val, err := strconv.Atoi(l[i+1])
		if err != nil {
			panic(err)
		}
		return val
	}
}

func safeAccessPrevValue(i int, l []string) int {
	if i == 0 {
		return -1
	} else {
		val, err := strconv.Atoi(l[i-1])
		if err != nil {
			panic(err)
		}
		return val
	}
}

func abs(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func Part1(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		safe := true
		l := strings.Split(line, " ")
		increasing := true
		for i := 0; i < (len(l) - 1); i++ {
			val, err := strconv.Atoi(l[i])
			if err != nil {
				panic(err)
			}
			curr := val
			val, err = strconv.Atoi(l[i+1])
			if err != nil {
				panic(err)
			}
			next := val

			if abs(curr, next) > 3 || abs(curr, next) < 1 {
				safe = false
				break
			}

			if i == 0 {
				increasing = next > curr
			} else {
				if increasing != (next > curr) {
					// fmt.Println(line, "UNSAFE")
					safe = false
					break
				}
			}
		}

		if safe {
			safeCount += 1
		}

	}

	return safeCount
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func isSafe(l []string) bool {
	increasing := true
	for i := 0; i < (len(l) - 1); i++ {
		val, err := strconv.Atoi(l[i])
		if err != nil {
			panic(err)
		}
		curr := val
		val, err = strconv.Atoi(l[i+1])
		if err != nil {
			panic(err)
		}
		next := val

		if abs(curr, next) > 3 || abs(curr, next) < 1 {
			return false
		}

		if i == 0 {
			increasing = next > curr
		} else {
			if increasing != (next > curr) {
				return false
			}
		}
	}

	return true
}

func Part2(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		l := strings.Split(line, " ")
		safe := isSafe(l)

		variant := make([]string, len(l)-1)
		if !safe {
			for i := range l {
				copy(variant, l[:i])
				copy(variant[i:], l[i+1:])
				safe = isSafe(variant)
				if safe {
					break
				}
			}
		}

		if safe {
			fmt.Println(line, "SAFE")
			safeCount += 1
		} else {
			fmt.Println(line, "UNSAFE")
		}
	}

	return safeCount
}

func calculateSafe(curr int, next int, increasing bool) bool {
	if abs(curr, next) > 3 || abs(curr, next) < 1 {
		return false
	}

	if increasing != (next > curr) {
		// fmt.Println(line, "UNSAFE")
		return false
	}
	return true
}
