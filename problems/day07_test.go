package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day07(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"$ cd /",
				"$ ls",
				"dir a",
				"14848514 b.txt",
				"8504156 c.dat",
				"dir d",
				"$ cd a",
				"$ ls",
				"dir e",
				"29116 f",
				"2557 g",
				"62596 h.lst",
				"$ cd e",
				"$ ls",
				"584 i",
				"$ cd ..",
				"$ cd ..",
				"$ cd d",
				"$ ls",
				"4060174 j",
				"8033020 d.log",
				"5626152 d.ext",
				"7214296 k",
			},
			part_one_assert: 95437,
			part_two_assert: 24933642,
		},
	}

	for _, test := range test_cases {
		in := day07{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
