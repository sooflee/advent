package problems

import (
	"github.com/sooflee/advent/internal"
)

type day02InputStruct struct {
	lines []string
}

func Day02(part int, file string) int {

	in := day02InputStruct{
		lines: internal.Read(file),
	}

	switch part {
	case 1:
		return in.part_one()
	case 2:
		return in.part_two()
	default:
		return 0
	}

}

func (input day02InputStruct) part_one() int {

	total_score := 0

	for _, line := range input.lines {
		elf := string(line[0])
		me := string(line[2])

		// tie
		if (elf == "A" && me == "X") || (elf == "B" && me == "Y") || (elf == "C" && me == "Z") {
			total_score = total_score + 3
		}

		// I win
		if (elf == "A" && me == "Y") || (elf == "B" && me == "Z") || (elf == "C" && me == "X") {
			total_score = total_score + 6
		}

		if me == "X" {
			total_score = total_score + 1
		}

		if me == "Y" {
			total_score = total_score + 2
		}

		if me == "Z" {
			total_score = total_score + 3
		}
	}
	return total_score
}

func (input day02InputStruct) part_two() int {

	total_score := 0

	for _, line := range input.lines {
		elf := string(line[0])
		me := string(line[2])

		// we need to lose
		if me == "X" {
			if elf == "A" {
				total_score = total_score + 3
			}

			if elf == "B" {
				total_score = total_score + 1
			}

			if elf == "C" {
				total_score = total_score + 2
			}
		}

		// we need to tie
		if me == "Y" {
			if elf == "A" {
				total_score = total_score + 1
			}

			if elf == "B" {
				total_score = total_score + 2
			}

			if elf == "C" {
				total_score = total_score + 3
			}

			total_score = total_score + 3
		}

		// we need to win
		if me == "Z" {
			if elf == "A" {
				total_score = total_score + 2
			}

			if elf == "B" {
				total_score = total_score + 3
			}

			if elf == "C" {
				total_score = total_score + 1
			}

			total_score = total_score + 6
		}
	}
	return total_score
}
