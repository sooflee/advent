package problems

import (
	"strings"

	"github.com/sooflee/advent/internal"
)

type day07 struct {
	lines []string
}

type Node struct {
	Data     string
	Children []*Node
	Parent   *Node
}

func Day07(part int, file string) int {

	in := day07{
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

func create_tree(lines []string) *Node {

	root := &Node{
		Data:     "/",
		Parent:   nil,
		Children: []*Node{},
	}

	curr_node := root

	for _, line := range lines {

		if !strings.HasPrefix(line, "$") {

			// we are ls-ing, so let's add a child
			child := &Node{}

			if strings.HasPrefix(line, "dir") {
				// dir
				child.Data = strings.Split(line, "dir ")[1]
				child.Parent = curr_node
				child.Children = []*Node{}

			} else {
				// file
				child.Data = strings.Split(line, " ")[0]
				child.Parent = curr_node
				child.Children = nil
			}
			curr_node.Children = append(curr_node.Children, child)

		} else if strings.HasPrefix(line, "$ cd") {
			dir := strings.Split(line, "$ cd ")[1]
			if dir == "/" {
				continue
			} else if dir == ".." {
				curr_node = curr_node.Parent
			} else {
				changed_directory := false
				for _, child := range curr_node.Children {
					if dir == child.Data {
						curr_node = child
						changed_directory = true
					}
				}
				if !changed_directory {
					panic("didn't change directory")
				}
			}

		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else {
			panic(line)
		}
	}
	return root
}

var file_sizes []int = []int{}

func postorder_sums(root *Node) int {
	total_children_value := 0

	for _, child := range root.Children {

		if child.Children != nil {
			total_children_value = total_children_value + postorder_sums(child)
		}
		if child.Children == nil {
			total_children_value = total_children_value + internal.StrToInt(child.Data)
		}

	}
	file_sizes = append(file_sizes, total_children_value)
	return total_children_value
}

func (input day07) part_one() int {

	root := create_tree(input.lines)
	postorder_sums(root)

	total := 0

	for _, size := range file_sizes {
		if size <= 100000 {
			total = total + size
		}
	}

	file_sizes = []int{}
	return total
}

func (input day07) part_two() int {
	root := create_tree(input.lines)
	total_size := postorder_sums(root)

	min_size := 70000000

	for _, size := range file_sizes {
		if total_size-size+30000000 < 70000000 {
			if size < min_size {
				min_size = size
			}
		}
	}

	file_sizes = []int{}
	return min_size
}
