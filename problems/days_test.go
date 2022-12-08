package problems

import (
	"testing"

	"github.com/sooflee/advent/internal"
	"github.com/stretchr/testify/assert"
)

func Test_All_Days(t *testing.T) {
	test_cases := []struct {
		day    int
		part   int
		assert interface{}
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
		{
			day:    2,
			part:   1,
			assert: 13005,
		},
		{
			day:    2,
			part:   2,
			assert: 11373,
		},
		{
			day:    3,
			part:   1,
			assert: 7727,
		},
		{
			day:    3,
			part:   2,
			assert: 2609,
		},
		{
			day:    4,
			part:   1,
			assert: 475,
		},
		{
			day:    4,
			part:   2,
			assert: 825,
		},
	}

	for _, test := range test_cases {
		answer, _ := internal.Call(Dailies, internal.GetFunc(test.day), test.part, internal.GetFile(test.day))
		assert.Equal(t, test.assert, answer[0].Interface().(int))
	}
}
