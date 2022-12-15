package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day05(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert string
		part_two_assert string
	}{
		{
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
				"",
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			part_one_assert: "CMZ",
			part_two_assert: "MCD",
		},
	}

	for _, test := range test_cases {
		in := day05{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
