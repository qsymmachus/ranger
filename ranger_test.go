package ranger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testVals[N Number] struct {
	start    N
	end      N
	step     N
	expected []N
}

func TestInt(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[int]{
			{1, 10, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int{7}},
			{10, 1, 1, []int{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[int]{
			{1, 10, 3, []int{1, 4, 7, 10}},
			{1, 9, 3, []int{1, 4, 7}},
			{-6, 6, 2, []int{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[int](1, 1000000)
	}
}

func TestInt8(t *testing.T) {

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[int8]{
			{1, 10, 1, []int8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int8{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int8{7}},
			{10, 1, 1, []int8{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int8](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[int8]{
			{1, 10, 3, []int8{1, 4, 7, 10}},
			{1, 9, 3, []int8{1, 4, 7}},
			{-6, 6, 2, []int8{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int8](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt8(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[int8](1, 127)
	}
}

func TestInt16(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[int16]{
			{1, 10, 1, []int16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int16{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int16{7}},
			{10, 1, 1, []int16{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int16](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[int16]{
			{1, 10, 3, []int16{1, 4, 7, 10}},
			{1, 9, 3, []int16{1, 4, 7}},
			{-6, 6, 2, []int16{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int16](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt16(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[int16](1, 32767)
	}
}

func TestInt64(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[int64]{
			{1, 10, 1, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []int64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []int64{7}},
			{10, 1, 1, []int64{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int64](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[int64]{
			{1, 10, 3, []int64{1, 4, 7, 10}},
			{1, 9, 3, []int64{1, 4, 7}},
			{-6, 6, 2, []int64{-6, -4, -2, 0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[int64](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkInt64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[int64](1, 1000000)
	}
}

func TestUint(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[uint]{
			{1, 10, 1, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint{7}},
			{10, 1, 1, []uint{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[uint]{
			{1, 10, 3, []uint{1, 4, 7, 10}},
			{1, 9, 3, []uint{1, 4, 7}},
			{0, 6, 2, []uint{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[uint](1, 1000000)
	}
}

func TestUint8(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[uint8]{
			{1, 10, 1, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint8{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint8{7}},
			{10, 1, 1, []uint8{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint8](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[uint8]{
			{1, 10, 3, []uint8{1, 4, 7, 10}},
			{1, 9, 3, []uint8{1, 4, 7}},
			{0, 6, 2, []uint8{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint8](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint8(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[uint8](1, 255)
	}
}

func TestUint16(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[uint16]{
			{1, 10, 1, []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint16{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint16{7}},
			{10, 1, 1, []uint16{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint16](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[uint16]{
			{1, 10, 3, []uint16{1, 4, 7, 10}},
			{1, 9, 3, []uint16{1, 4, 7}},
			{0, 6, 2, []uint16{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint16](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint16(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[uint16](1, 65535)
	}
}

func TestUint64(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[uint64]{
			{1, 10, 1, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint64{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint64{7}},
			{10, 1, 1, []uint64{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint64](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[uint64]{
			{1, 10, 3, []uint64{1, 4, 7, 10}},
			{1, 9, 3, []uint64{1, 4, 7}},
			{0, 6, 2, []uint64{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[uint64](test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[uint64](1, 1000000)
	}
}

func TestFloat32(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[float32]{
			{1, 10, 1, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []float32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []float32{7}},
			{10, 1, 1, []float32{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[float32](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[float32]{
			{1, 10, 3, []float32{1, 4, 7, 10}},
			{1, 2, 0.3, []float32{1, 1.3, 1.6, 1.9}},
			{1, 2.5, 0.5, []float32{1, 1.5, 2, 2.5}},
		}

		for _, test := range tests {
			actual := Range[float32](test.start, test.end, Step(test.step))

			for i := range actual {
				assert.InDelta(t, test.expected[i], actual[i], 0.0001)
			}
		}
	})
}

func BenchmarkFloat32(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[float32](1, 1000000)
	}
}

func TestFloat64(t *testing.T) {
	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals[float64]{
			{1, 10, 1, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []float64{7}},
			{10, 1, 1, []float64{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Range[float64](test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals[float64]{
			{1, 10, 3, []float64{1, 4, 7, 10}},
			{1, 2, 0.3, []float64{1, 1.3, 1.6, 1.9}},
			{1, 2.5, 0.5, []float64{1, 1.5, 2, 2.5}},
		}

		for _, test := range tests {
			actual := Range[float64](test.start, test.end, Step(test.step))

			for i := range actual {
				assert.InDelta(t, test.expected[i], actual[i], 0.0001)
			}
		}
	})
}

func BenchmarkFloat64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Range[float64](1, 1000000)
	}
}

func TestRune(t *testing.T) {
	type testVals struct {
		start    rune
		end      rune
		step     int
		expected []rune
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{'a', 'k', 1, []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Rune(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{'a', 'g', 2, []rune{'a', 'c', 'e', 'g'}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Rune(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkRune(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Rune('a', 'z')
	}
}

func TestString(t *testing.T) {
	type testVals struct {
		start    string
		end      string
		step     int
		expected []string
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{"a", "k", 1, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}},
		}

		for _, test := range tests {
			actual, err := String(test.start, test.end)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{"a", "g", 2, []string{"a", "c", "e", "g"}},
		}

		for _, test := range tests {
			actual, err := String(test.start, test.end, Step(test.step))
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		}
	})

	t.Run("It should return an error if either string has more than one character", func(t *testing.T) {
		strings, err := String("abc", "z")
		assert.Empty(t, strings)
		assert.Error(t, err)

		strings, err = String("a", "xyz")
		assert.Empty(t, strings)
		assert.Error(t, err)
	})

	t.Run("It should return an error if either string is empty", func(t *testing.T) {
		strings, err := String("", "z")
		assert.Empty(t, strings)
		assert.Error(t, err)

		strings, err = String("a", "")
		assert.Empty(t, strings)
		assert.Error(t, err)
	})
}

func BenchmarkString(b *testing.B) {
	for i := 1; i < b.N; i++ {
		String("a", "z")
	}
}
