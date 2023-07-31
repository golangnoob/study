package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShrink(t *testing.T) {
	testCases := []struct {
		name    string
		length  int
		oldCap  int
		wantCap int
	}{
		{
			name:    "小于256",
			oldCap:  128,
			length:  10,
			wantCap: 128,
		},
		{
			name:    "小于1024，c/l小于3",
			oldCap:  1000,
			length:  600,
			wantCap: 1000,
		},
		{
			name:    "小于1024，c/l大于3",
			oldCap:  1000,
			length:  200,
			wantCap: 500,
		},
		{
			name:    "大于1024，c/l小于2",
			oldCap:  2000,
			length:  1200,
			wantCap: 2000,
		},
		{
			name:    "大于1024，c/l大于2",
			oldCap:  2000,
			length:  900,
			wantCap: 1500,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testSlice := make([]int, tc.length, tc.oldCap)

			s := Shrink[int](testSlice)
			assert.Equal(t, tc.wantCap, cap(s))
		})
	}
}
