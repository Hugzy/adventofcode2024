package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	rule   string
	first  int
	second int
}

type pageUpdates struct{}

var (
	partOneInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	partTwoInput = ``
)

func convToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		panic(err)
	}

	return i
}

func parseRules(input string) []rule {
	lines := strings.Split(input, "\n")
	rules := []rule{}
	for _, line := range lines {
		parsed := strings.Split(line, "|")
		rule := rule{
			rule:   line,
			first:  convToInt(parsed[0]),
			second: convToInt(parsed[1]),
		}
		rules = append(rules, rule)
	}
	return rules
}

func PartOneTest() {
	PartOne(partOneInput)
}

func PartOne(input string) {
	split := strings.Split(input, "\n\n")
	rules := parseRules(split[0])
	fmt.Printf("%v", rules)
}
