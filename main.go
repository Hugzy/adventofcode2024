package main

import (
	"adventofcode2024/day1"
	"adventofcode2024/util"
)

func Day1() {
	result := day1.TestPart1()
	Assert(result, 11)
	result = day1.Part1(util.GetInput("resources/day1part1"))
	PrintAnswer(result)
	result = day1.TestPart2()
	Assert(result, 31)
	result = day1.Part2(util.GetInput("resources/day1part2"))
	PrintAnswer(result)
}

func main() {
	Day1()
}
