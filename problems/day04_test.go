package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day04(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,3-7",
				"6-6,4-6",
				"2-6,4-8",
			},
			part_one_assert: 2,
			part_two_assert: 4,
		},
	}

	for _, test := range test_cases {
		in := day04{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
