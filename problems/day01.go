package problems

import (
	"container/heap"
	"strconv"

	"github.com/sooflee/advent/internal"
)

type day01 struct {
	lines []string
}

func Day01(part int, file string) int {

	in := day01{
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

func (input day01) part_one() int {
	current_largest := -1
	current_value := 0

	for _, line := range input.lines {
		if line == "" {
			if current_value > current_largest {
				current_largest = current_value
			}
			current_value = 0
			continue
		}
		n, _ := strconv.Atoi(line)
		current_value = current_value + n
	}

	return current_largest
}

func (input day01) part_two() int {
	pq := &internal.IntHeap{}
	heap.Init(pq)

	current_value := 0
	for _, line := range input.lines {
		if line == "" {
			heap.Push(pq, current_value)
			current_value = 0
		} else {
			n, _ := strconv.Atoi(line)
			current_value = current_value + (n * -1)
		}
	}

	largest_three := (heap.Pop(pq).(int) + heap.Pop(pq).(int) + heap.Pop(pq).(int)) * -1
	return largest_three
}
