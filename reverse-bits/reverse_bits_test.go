package reverse_bits

import (
	"fmt"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	tcases := []struct {
		num    uint32
		output uint32
	}{
		{
			num:    43261596,
			output: 964176192,
		},
		{
			num:    4294967293,
			output: 3221225471,
		},
	}

	for _, tc := range tcases {
		t.Run(fmt.Sprintf("%d", tc.num), func(t *testing.T) {
			if output := reverseBits(tc.num); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})
	}
}
