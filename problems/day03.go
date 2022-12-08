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

func calculate_priorities(count map[rune]int) int {
	total_priorities := 0
	for key, element := range count {
		var priority int
		if int(key) > 64 && int(key) < 91 {
			priority = ((int(key) - 38) * element)
		} else if int(key) > 96 && int(key) < 123 {
			priority = ((int(key) - 96) * element)
		}
		total_priorities = total_priorities + priority
	}
	return total_priorities
}

func (input day03) part_one() int {
	count := make(map[rune]int)
	for _, line := range input.lines {
		exist := make(map[rune]bool)
		for _, letter := range line[0 : len(line)/2] {
			exist[letter] = true
		}
		for _, letter := range line[len(line)/2:] {
			if _, ok := exist[letter]; ok {
				if _, ok = count[letter]; ok {
					count[letter]++
					break
				} else {
					count[letter] = 1
					break
				}
			}
		}
	}
	return calculate_priorities(count)
}

func (input day03) part_two() int {

	var filter_1 map[rune]bool
	var filter_2 map[rune]bool
	count := make(map[rune]int)

	for i, line := range input.lines {

		if i%3 == 0 {
			filter_1 = make(map[rune]bool)
			for _, letter := range line {
				filter_1[letter] = true
			}
		}

		if i%3 == 1 {
			filter_2 = make(map[rune]bool)
			for _, letter := range line {
				if filter_1[letter] {
					filter_2[letter] = true
				}
			}
		}

		if i%3 == 2 {
			for _, letter := range line {
				if filter_2[letter] {
					if _, ok := count[letter]; ok {
						count[letter]++
						break
					} else {
						count[letter] = 1
						break
					}
				}
			}
		}
	}
	return calculate_priorities(count)
}
