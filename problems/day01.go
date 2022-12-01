package problems

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

func Day01(part int, file string) int {

	f, _ := os.Open(file)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	if part == 1 {

		current_largest := -1
		current_value := 0

		for scanner.Scan() {
			v := scanner.Text()
			if v == "" {
				if current_value > current_largest {
					current_largest = current_value
				}
				current_value = 0
				continue
			}
			n, _ := strconv.Atoi(v)
			current_value = current_value + n
		}

		return current_largest

	} else {

		pq := &IntHeap{}
		heap.Init(pq)

		current_value := 0
		for scanner.Scan() {
			v := scanner.Text()
			if v == "" {
				heap.Push(pq, current_value)
				current_value = 0
			} else {
				n, _ := strconv.Atoi(v)
				current_value = current_value + (n * -1)
			}
		}

		largest_three := (heap.Pop(pq).(int) + heap.Pop(pq).(int) + heap.Pop(pq).(int)) * -1
		return largest_three

	}

}
