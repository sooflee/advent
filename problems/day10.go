package problems

import (
	"github.com/sooflee/advent/internal"
)

type day10 struct {
	lines []string
}

func Day10(part int, file string) int {

	in := day10{
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

func (input day10) part_one() int {
	return 0
}

func (input day10) part_two() int {
	return 0
}
