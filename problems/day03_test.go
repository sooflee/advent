package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day03(t *testing.T) {
	test_cases := []struct {
		input           []string
		part_one_assert int
		part_two_assert int
	}{}

	for _, test := range test_cases {
		in := day03InputStruct{
			lines: test.input,
		}
		assert.Equal(t, test.part_one_assert, in.part_one())
		assert.Equal(t, test.part_two_assert, in.part_two())
	}
}
