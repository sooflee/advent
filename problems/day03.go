package problems

import (
	"github.com/sooflee/advent/internal"
)

type day03 struct {
	lines []string
}

func Day03(part int, file string) int {

	in := day03{
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

func (input day03) part_one() int {
	return 0
}

func (input day03) part_two() int {
	return 0
}
