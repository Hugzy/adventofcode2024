package day3

import (
	"adventofcode2024/util"
	"fmt"
	"regexp"
)

var partOneInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestPartOne() int {
	return PartOne(partOneInput)
}

type memory struct {
	input    string
	commands []command
}

type command struct {
	operation string
	x         int
	y         int
	disabled  bool
}

func (m *memory) findAllOccurrences(regex string) {
	r := regexp.MustCompile(regex)
	matches := r.FindAllString(m.input, -1)

	commands := make([]command, 0)
	rd := regexp.MustCompile(`\d{1,3}`)
	for _, v := range matches {
		nums := rd.FindAllString(v, -1)
		cmd := command{
			operation: "mul",
			x:         util.GetInt(nums[0]),
			y:         util.GetInt(nums[1]),
		}
		commands = append(commands, cmd)
	}

	m.commands = commands
}

func (cmd command) Execute() int {
	switch cmd.operation {
	case "mul":
		return (cmd.x * cmd.y)
	}

	return -1
}

func PartOne(line string) int {
	memory := memory{input: line}
	memory.findAllOccurrences("mul\\(\\d{1,3},\\d{1,3}\\)")

	result := 0
	for _, cmd := range memory.commands {
		fmt.Println(cmd)
		result += cmd.Execute()
	}

	return result
}
