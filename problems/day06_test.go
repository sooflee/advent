package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day06(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{
		{
			input: []string{
				"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			part_one_assert: 7,
			part_two_assert: 19,
		},
		{
			input: []string{
				"bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			part_one_assert: 5,
			part_two_assert: 23,
		},
		{
			input: []string{
				"nppdvjthqldpwncqszvftbrmjlhg",
			},
			part_one_assert: 6,
			part_two_assert: 23,
		},
		{
			input: []string{
				"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			part_one_assert: 10,
			part_two_assert: 29,
		},
		{
			input: []string{
				"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			part_one_assert: 11,
			part_two_assert: 26,
		},
	}

	for _, test := range test_cases {
		in := day06{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
