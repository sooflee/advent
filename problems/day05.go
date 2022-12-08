package problems

import (
	"strings"

	"github.com/sooflee/advent/internal"
)

type day05 struct {
	lines []string
}

type instruction struct {
	amount int
	from   int
	to     int
}

func read_input(input []string) ([][]string, []instruction) {
	stacks := []string{}
	var instructions []instruction

	second_part := false
	for _, line := range input {
		if !second_part {
			for i, char := range line {
				if i%4 == 1 {
					if string(char) == "1" {
						second_part = true
						break
					}
					stacks = append(stacks, string(char))
				}
			}
		}
		if second_part {
			if strings.HasPrefix(line, "move") {
				amt := strings.Split(line[5:], "from")
				fromto := strings.Split(amt[1], "to")
				instruct := instruction{
					amount: internal.StrToInt(amt[0]),
					from:   internal.StrToInt(fromto[0]),
					to:     internal.StrToInt(fromto[1]),
				}
				instructions = append(instructions, instruct)
			}
		}
	}

	output := [][]string{}
	columns := (len(input[0]) / 4) + 1
	rows := len(stacks) / columns

	for j := 0; j < columns; j++ {
		append_string := []string{}
		for i := rows - 1; i > -1; i-- {
			box := stacks[columns*i+j]
			if box != " " {
				append_string = append(append_string, box)
			}
		}
		output = append(output, append_string)
	}

	return output, instructions
}

func Day05(part int, file string) string {

	in := day05{
		lines: internal.Read(file),
	}

	switch part {
	case 1:
		return in.part_one()
	case 2:
		return in.part_two()
	default:
		return "no"
	}

}

func (input day05) part_one() string {

	stacks, instructions := read_input(input.lines)
	result := []string{}

	for _, instruction := range instructions {
		for i := instruction.amount; i > 0; i-- {
			// pop
			n := len(stacks[instruction.from-1]) - 1
			value := stacks[instruction.from-1][n]
			stacks[instruction.from-1][n] = ""
			stacks[instruction.from-1] = stacks[instruction.from-1][:n]

			// push
			stacks[instruction.to-1] = append(stacks[instruction.to-1], value)
		}
	}

	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}
	return strings.Join(result, "")
}

func (input day05) part_two() string {
	stacks, instructions := read_input(input.lines)
	result := []string{}

	for _, instruction := range instructions {

		// grab instruction.amount of things
		to_grab := []string{}

		for i := instruction.amount; i > 0; i-- {
			n := len(stacks[instruction.from-1]) - 1
			value := stacks[instruction.from-1][n]
			stacks[instruction.from-1][n] = ""
			stacks[instruction.from-1] = stacks[instruction.from-1][:n]

			to_grab = append(to_grab, value)
		}

		for i := instruction.amount - 1; i > -1; i-- {
			// // push
			stacks[instruction.to-1] = append(stacks[instruction.to-1], to_grab[i])
		}
	}

	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}
	return strings.Join(result, "")
}
