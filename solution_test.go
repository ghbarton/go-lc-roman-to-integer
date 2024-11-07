package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name string
		data string
		exp  int
	}{
		{
			name: "returns 3",
			data: "III",
			exp:  3,
		},
		{
			name: "returns 58",
			data: "LVIII",
			exp:  58,
		},
		{
			name: "returns 1994",
			data: "MCMXCIV",
			exp:  1994,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
