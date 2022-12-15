package problems

import (
	"strings"

	"github.com/sooflee/advent/internal"
)

type day09 struct {
	lines []string
}

func Day09(part int, file string) int {

	in := day09{
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

type coordinate struct {
	x int
	y int
}

func new_t_curr(h_curr *coordinate, t_curr *coordinate) *coordinate {
	x := h_curr.x - t_curr.x
	y := h_curr.y - t_curr.y

	if (x <= 1 && x >= -1) && (y <= 1 && y >= -1) {
		return t_curr
	}

	if x > 0 {
		t_curr.x++
	} else if x < 0 {
		t_curr.x--
	}

	if y > 0 {
		t_curr.y++
	} else if y < 0 {
		t_curr.y--
	}

	return t_curr
}

func get_dir(direction string) coordinate {
	dir := coordinate{}
	switch direction {
	case "R":
		dir = coordinate{x: 1, y: 0}
	case "L":
		dir = coordinate{x: -1, y: 0}
	case "U":
		dir = coordinate{x: 0, y: 1}
	case "D":
		dir = coordinate{x: 0, y: -1}
	default:
		panic("invalid input")
	}
	return dir
}

func (input day09) part_one() int {
	t_visited := make(map[coordinate]bool)

	h_curr := &coordinate{
		x: 0,
		y: 0,
	}

	t_curr := &coordinate{
		x: 0,
		y: 0,
	}

	for _, line := range input.lines {
		parsed := strings.Split(line, " ")
		direction := parsed[0]
		steps := internal.StrToInt(parsed[1])
		dir := get_dir(direction)

		for i := 0; i < steps; i++ {
			h_curr.x = h_curr.x + dir.x
			h_curr.y = h_curr.y + dir.y
			t_curr = new_t_curr(h_curr, t_curr)
			t_visited[*t_curr] = true
		}
	}

	return len(t_visited)
}

func (input day09) part_two() int {
	t_visited := make(map[coordinate]bool)

	// rope[0] is H and rope[9] is being kept track by t_visited
	rope := make([]*coordinate, 10)
	for i := range rope {
		rope[i] = &coordinate{
			x: 0,
			y: 0,
		}
	}

	for _, line := range input.lines {
		parsed := strings.Split(line, " ")
		direction := parsed[0]
		steps := internal.StrToInt(parsed[1])
		dir := get_dir(direction)

		for i := 0; i < steps; i++ {
			rope[0].x = rope[0].x + dir.x
			rope[0].y = rope[0].y + dir.y
			for i := range rope {
				if i == 0 {
					continue
				}
				rope[i] = new_t_curr(rope[i-1], rope[i])
			}
			t_visited[*rope[9]] = true
		}
	}
	return len(t_visited)
}
