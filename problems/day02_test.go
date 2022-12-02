package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day02(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"A Y",
				"B X",
				"C Z",
			},
			part_one_assert: 15,
			part_two_assert: 12,
		},
	}

	for _, test := range test_cases {
		in := day02InputStruct{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
