package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day01(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"1000",
				"2000",
				"3000",
				"",
				"4000",
				"",
				"5000",
				"6000",
				"",
				"7000",
				"8000",
				"9000",
				"",
				"10000",
				"",
			},
			part_one_assert: 24000,
			part_two_assert: 45000,
		},
	}

	for _, test := range test_cases {
		in := day01InputStruct{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
