package day3

import (
	"adventofcode2024/util"
	"regexp"
)

var (
	partOneInput = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	partTwoInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
)

func TestPartOne() int {
	return PartOne(partOneInput)
}

type memory struct {
	input      string
	commands   []command
	isDisabled bool
}

type command struct {
	operation string
	x         int
	y         int
}

func (m *memory) split(regex string) {
}

func (m *memory) findAllOccurrences(regex string) {
	r := regexp.MustCompile(regex)
	matches := r.FindAllString(m.input, -1)
	commands := make([]command, 0)
	rd := regexp.MustCompile(`\d{1,3}`)
	for _, v := range matches {
		cmd := command{}
		switch v {
		case "don't()":
			cmd = command{
				operation: "dont",
				x:         0,
				y:         0,
			}
		case "do()":
			cmd = command{
				operation: "do",
				x:         0,
				y:         0,
			}
		default:
			nums := rd.FindAllString(v, -1)
			cmd = command{
				operation: "mul",
				x:         util.GetInt(nums[0]),
				y:         util.GetInt(nums[1]),
			}
		}
		commands = append(commands, cmd)
	}

	m.commands = commands
}

func (cmd command) execute() int {
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
		result += cmd.execute()
	}

	return result
}

func TestPartTwo() int {
	return PartTwo(partTwoInput)
}

func (m *memory) execute2() int {
	sum := 0
	for _, cmd := range m.commands {
		switch cmd.operation {
		case "dont":
			m.isDisabled = true
		case "do":
			m.isDisabled = false
		case "mul":
			if m.isDisabled {
				continue
			}
			sum += (cmd.x * cmd.y)
		}
	}
	return sum
}

func PartTwo(line string) int {
	memory := memory{input: line, isDisabled: false}
	memory.findAllOccurrences(`mul\(\d{1,3},\d{1,3}\)|(don't\(\)|do\(\))`)
	return memory.execute2()
}
