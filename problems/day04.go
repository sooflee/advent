package problems

import (
	"strings"

	"github.com/sooflee/advent/internal"
)

type day04 struct {
	lines []string
}

func Day04(part int, file string) int {

	in := day04{
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

func parse_section(line string) (int, int, int, int) {
	pairs := strings.Split(line, ",")
	section_1 := strings.Split(pairs[0], "-")
	section_2 := strings.Split(pairs[1], "-")
	min_1 := internal.StrToInt(section_1[0])
	max_1 := internal.StrToInt(section_1[1])
	min_2 := internal.StrToInt(section_2[0])
	max_2 := internal.StrToInt(section_2[1])
	return min_1, max_1, min_2, max_2
}

func (input day04) part_one() int {

	total_fully_contain := 0

	for _, line := range input.lines {

		min_1, max_1, min_2, max_2 := parse_section(line)
		if (min_1 <= min_2 && max_1 >= max_2) ||
			(min_2 <= min_1 && max_2 >= max_1) {
			total_fully_contain++

		}

	}
	return total_fully_contain
}

func (input day04) part_two() int {
	overlap_count := 0

	for _, line := range input.lines {

		min_1, max_1, min_2, max_2 := parse_section(line)

		if !((min_1 > max_2) || (min_2 > max_1)) {
			overlap_count++
		}

	}
	return overlap_count
}
