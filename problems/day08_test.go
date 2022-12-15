package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day08(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			part_one_assert: 21,
			part_two_assert: 8,
		},
	}

	for _, test := range test_cases {
		in := day08{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
