package main

import (
	"adventofcode2024/day1"
	"adventofcode2024/day2"
	"adventofcode2024/day3"
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

func Day2() {
	result := day2.TestPart1()
	Assert(result, 2)
	result = day2.Part1(util.GetInput("resources/day2part1"))
	PrintAnswer(result)
	result = day2.TestPart2()
	Assert(result, 5)
	result = day2.Part2(util.GetInput("resources/day2part1"))
	PrintAnswer(result)
}

func Day3() {
	result := day3.TestPartOne()
	Assert(result, 161)
	result = day3.PartOne(util.ReadFile("resources/day3part1"))
	PrintAnswer(result)
}

func main() {
	// Day1()
	// Day2()
	Day3()
}
