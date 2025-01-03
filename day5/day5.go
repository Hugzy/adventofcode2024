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

type pageUpdate struct {
	page []int
}

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

func convListToInt(s []string) []int {
	result := make([]int, 0)
	for _, v := range s {
		result = append(result, convToInt(v))
	}

	return result
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

func parsePages(input string) []pageUpdate {
	lines := strings.Split(input, "\n")
	pageUpdates := []pageUpdate{}
	for _, line := range lines {
		parsed := strings.Split(line, ",")
		pageUpdate := pageUpdate{
			page: convListToInt(parsed),
		}
		pageUpdates = append(pageUpdates, pageUpdate)
	}
	return pageUpdates
}

func PartOneTest() {
	PartOne(partOneInput)
}

func PartOne(input string) {
	split := strings.Split(input, "\n\n")
	rules := parseRules(split[0])
	fmt.Printf("%v", rules)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	pageUpdates := parsePages(split[1])
	fmt.Printf("%v", pageUpdates)
}
