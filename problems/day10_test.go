package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day10(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"noop",
				"addx 3",
				"addx -5",
			},
			part_one_assert: 13,
			part_two_assert: 1,
		},
	}

	for _, test := range test_cases {
		in := day10{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		// assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
