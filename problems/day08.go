package problems

import (
	"github.com/sooflee/advent/internal"
)

type day08 struct {
	lines []string
}

func Day08(part int, file string) int {

	in := day08{
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

func read_grid(lines []string) [][]int {
	result := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			row = append(row, internal.StrToInt(string(char)))
		}
		result = append(result, row)
	}
	return result
}

// I copied this code to rotate a grid
func rotate_grid(grid [][]int) [][]int {
	tp := make([][]int, len(grid[0]))
	for i := range tp {
		tp[i] = make([]int, len(grid))
		for j := range tp[i] {
			tp[i][j] = grid[j][i]
		}
	}

	for i := range tp {
		for j, k := 0, len(tp[i])-1; j < k; j, k = j+1, k-1 {
			tp[i][j], tp[i][k] = tp[i][k], tp[i][j]
		}
	}
	return tp
}

func (input day08) part_one() int {
	grid := read_grid(input.lines)
	marker := make([][]int, len(grid))
	for i := range marker {
		marker[i] = make([]int, len(grid))
	}

	for a := 0; a < 4; a++ {
		grid = rotate_grid(grid)
		marker = rotate_grid(marker)
		for j := 1; j < len(grid[0])-1; j++ {
			curr_max := grid[0][j]
			for i := 1; i < len(grid)-1; i++ {
				if grid[i][j] > curr_max {
					marker[i][j] = 1
					curr_max = grid[i][j]
				}
			}
		}
	}

	total_count := len(grid) + len(grid) + len(grid[0]) + len(grid[0]) - 4
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			total_count = total_count + marker[i][j]
		}
	}
	return total_count
}

func scenic_score(grid [][]int, i int, j int) int {
	up_score := 0
	for n := i - 1; n >= 0; n-- {
		up_score++
		if !(grid[n][j] < grid[i][j]) {
			break
		}
	}

	down_score := 0
	for n := i + 1; n < len(grid); n++ {
		down_score++
		if !(grid[n][j] < grid[i][j]) {
			break
		}
	}

	left_score := 0
	for n := j - 1; n >= 0; n-- {
		left_score++
		if !(grid[i][n] < grid[i][j]) {
			break
		}
	}

	right_score := 0
	for n := j + 1; n < len(grid[i]); n++ {
		right_score++
		if !(grid[i][n] < grid[i][j]) {
			break
		}
	}

	return up_score * down_score * left_score * right_score
}

func (input day08) part_two() int {
	grid := read_grid(input.lines)
	max_scenic_score := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			scenic_score := scenic_score(grid, i, j)
			if scenic_score > max_scenic_score {
				max_scenic_score = scenic_score
			}
		}
	}

	return max_scenic_score
}
