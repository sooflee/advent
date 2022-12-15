package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day09(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			part_one_assert: 13,
			part_two_assert: 1,
		},
		{
			input: []string{
				"R 5",
				"U 8",
				"L 8",
				"D 3",
				"R 17",
				"D 10",
				"L 25",
				"U 20",
			},
			part_one_assert: 88,
			part_two_assert: 36,
		},
	}

	for _, test := range test_cases {
		in := day09{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
