package day1

import (
	"adventofcode2024/util"
	"slices"
	"strconv"
	"strings"
)

var part1TestInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func abs(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func TestPart1() int {
	return Part1(util.SplitFile(part1TestInput, util.N))
}

func Part1(lines []string) int {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, v := range lines {
		s := strings.Split(v, "   ")
		i2, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, i2)

		i2, err = strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		list2 = append(list2, i2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	result := 0
	for i := 0; i < len(list1); i++ {
		i1 := list1[i]
		i2 := list2[i]

		result += abs(i1, i2)
	}

	return result
}

func TestPart2() int {
	return Part2(util.SplitFile(part1TestInput, util.N))
}

func Part2(lines []string) int {
	l1 := make([]int, 0)
	m2 := make(map[int]int)

	for _, v := range lines {
		s := strings.Split(v, "   ")
		i2, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		l1 = append(l1, i2)

		i2, err = strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		m2[i2] = m2[i2] + 1
	}

	similarity := 0
	for _, v := range l1 {
		similarity += v * m2[v]
	}

	return similarity
}
