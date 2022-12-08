package problems

import (
	"github.com/golang-collections/collections/set"
	"github.com/sooflee/advent/internal"
)

type day06 struct {
	lines []string
}

func Day06(part int, file string) int {

	in := day06{
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

func run_input_for_size(size int, lines []string) int {
	for i := size - 1; i <= len(lines[0])-1; i++ {

		set := set.New()

		for j := 0; j <= size-1; j++ {
			set.Insert(lines[0][i-j])
		}

		if set.Len() == size {
			return i + 1
		}
	}

	return 0
}

func (input day06) part_one() int {
	return run_input_for_size(4, input.lines)
}

func (input day06) part_two() int {
	return run_input_for_size(14, input.lines)
}
