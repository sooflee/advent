package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Too lazy for the custom test cases
	Testing against real answers
*/

func Test_Days(t *testing.T) {
	test_cases := []struct {
		day    int
		part   int
		assert int
	}{
		{
			day:    1,
			part:   1,
			assert: 70613,
		},
		{
			day:    1,
			part:   2,
			assert: 205805,
		},
	}

	for _, test := range test_cases {
		answer, _ := Call(Dailies, GetFunc(test.day), test.part, GetFile(test.day))
		assert.Equal(t, test.assert, answer[0].Interface().(int))
	}
}
