package ranger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	type testVals struct {
		start    int
		end      int
		expected []int
	}

	tests := []testVals{
		{1, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{-5, 5, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
		{7, 7, []int{7}},
		{10, 1, []int{}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Int(test.start, test.end))
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Int(1, 1000000)
	}
}
