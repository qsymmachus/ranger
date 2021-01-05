package ranger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	type testVals struct {
		start    int
		end      int
		step     int
		expected []int
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int{7}},
			{10, 1, 1, []int{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []int{1, 4, 7, 10}},
			{1, 9, 3, []int{1, 4, 7}},
			{-6, 6, 2, []int{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Int(1, 1000000)
	}
}

func TestInt8(t *testing.T) {
	type testVals struct {
		start    int8
		end      int8
		step     int
		expected []int8
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int8{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int8{7}},
			{10, 1, 1, []int8{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int8(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []int8{1, 4, 7, 10}},
			{1, 9, 3, []int8{1, 4, 7}},
			{-6, 6, 2, []int8{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int8(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt8(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Int8(1, 127)
	}
}

func TestInt16(t *testing.T) {
	type testVals struct {
		start    int16
		end      int16
		step     int
		expected []int16
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int16{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int16{7}},
			{10, 1, 1, []int16{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int16(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []int16{1, 4, 7, 10}},
			{1, 9, 3, []int16{1, 4, 7}},
			{-6, 6, 2, []int16{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int16(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt16(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Int16(1, 32767)
	}
}

func TestInt64(t *testing.T) {
	type testVals struct {
		start    int64
		end      int64
		step     int
		expected []int64
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int64{7}},
			{10, 1, 1, []int64{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int64(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []int64{1, 4, 7, 10}},
			{1, 9, 3, []int64{1, 4, 7}},
			{-6, 6, 2, []int64{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Int64(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Int64(1, 1000000)
	}
}
