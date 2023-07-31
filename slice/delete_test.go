package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteByIndex(t *testing.T) {
	testCases := []struct {
		name      string
		srcSlice  []int
		index     int
		wantSlice []int
		wantVal   int
		wantError error
	}{
		{
			name:      "index first",
			srcSlice:  []int{11, 12, 13, 14},
			index:     0,
			wantSlice: []int{12, 13, 14},
			wantVal:   11,
			wantError: nil,
		},
		{
			name:      "index middle",
			srcSlice:  []int{97, 98, 99, 100},
			index:     2,
			wantSlice: []int{97, 98, 100},
			wantVal:   99,
			wantError: nil,
		},
		{
			name:      "index last",
			srcSlice:  []int{90, 91, 92, 93},
			index:     3,
			wantSlice: []int{90, 91, 92},
			wantVal:   93,
			wantError: nil,
		},
		{
			name:      "index out of range",
			srcSlice:  []int{21, 22, 23, 24},
			index:     5,
			wantSlice: nil,
			wantVal:   0,
			wantError: IndexOutOfRange,
		},
		{
			name:      "index less than 0",
			srcSlice:  []int{31, 32, 33, 34},
			index:     -1,
			wantSlice: nil,
			wantVal:   0,
			wantError: IndexOutOfRange,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, val, err := DeleteByIndex(tc.srcSlice, tc.index)
			assert.Equal(t, tc.wantError, err)
			if err == nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}
