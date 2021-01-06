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

func TestUint(t *testing.T) {
	type testVals struct {
		start    uint
		end      uint
		step     int
		expected []uint
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint{7}},
			{10, 1, 1, []uint{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []uint{1, 4, 7, 10}},
			{1, 9, 3, []uint{1, 4, 7}},
			{0, 6, 2, []uint{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Uint(1, 1000000)
	}
}

func TestUint8(t *testing.T) {
	type testVals struct {
		start    uint8
		end      uint8
		step     int
		expected []uint8
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint8{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint8{7}},
			{10, 1, 1, []uint8{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint8(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []uint8{1, 4, 7, 10}},
			{1, 9, 3, []uint8{1, 4, 7}},
			{0, 6, 2, []uint8{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint8(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint8(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Uint8(1, 255)
	}
}

func TestUint16(t *testing.T) {
	type testVals struct {
		start    uint16
		end      uint16
		step     int
		expected []uint16
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []uint16{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint16{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint16{7}},
			{10, 1, 1, []uint16{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint16(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []uint16{1, 4, 7, 10}},
			{1, 9, 3, []uint16{1, 4, 7}},
			{0, 6, 2, []uint16{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint16(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint16(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Uint16(1, 65535)
	}
}

func TestUint64(t *testing.T) {
	type testVals struct {
		start    uint64
		end      uint64
		step     int
		expected []uint64
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{0, 5, 1, []uint64{0, 1, 2, 3, 4, 5}},
			{7, 7, 1, []uint64{7}},
			{10, 1, 1, []uint64{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint64(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, []uint64{1, 4, 7, 10}},
			{1, 9, 3, []uint64{1, 4, 7}},
			{0, 6, 2, []uint64{0, 2, 4, 6}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Uint64(test.start, test.end, Step(test.step)))
		}
	})
}

func BenchmarkUint64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Uint64(1, 1000000)
	}
}

func TestFloat32(t *testing.T) {
	type testVals struct {
		start     float32
		end       float32
		step      int
		floatStep float32
		expected  []float32
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, 1, []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, 1, []float32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, 1, []float32{7}},
			{10, 1, 1, 1, []float32{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Float32(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, 1, []float32{1, 4, 7, 10}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Float32(test.start, test.end, Step(test.step)))
		}
	})

	t.Run("It should generate intervals using the provided 'FloatStep' option", func(t *testing.T) {
		tests := []testVals{
			{1, 2, 1, 0.3, []float32{1, 1.3, 1.6, 1.9}},
			{1, 2.5, 3, 0.5, []float32{1, 1.5, 2, 2.5}},
		}

		for _, test := range tests {
			actual := Float32(test.start, test.end, FloatStep(test.floatStep))

			for i := range actual {
				assert.InDelta(t, test.expected[i], actual[i], 0.0001)
			}
		}
	})
}

func BenchmarkFloat32(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Float32(1, 1000000)
	}
}

func TestFloat64(t *testing.T) {
	type testVals struct {
		start     float64
		end       float64
		step      int
		floatStep float32
		expected  []float64
	}

	t.Run("It should generate intervals with a default step of 1", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 1, 1, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			{-5, 5, 1, 1, []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}},
			{7, 7, 1, 1, []float64{7}},
			{10, 1, 1, 1, []float64{}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Float64(test.start, test.end))
		}
	})

	t.Run("It should generate intervals using the provided 'Step' option", func(t *testing.T) {
		tests := []testVals{
			{1, 10, 3, 1, []float64{1, 4, 7, 10}},
		}

		for _, test := range tests {
			assert.Equal(t, test.expected, Float64(test.start, test.end, Step(test.step)))
		}
	})

	t.Run("It should generate intervals using the provided 'FloatStep' option", func(t *testing.T) {
		tests := []testVals{
			{1, 2, 1, 0.3, []float64{1, 1.3, 1.6, 1.9}},
			{1, 2.5, 3, 0.5, []float64{1, 1.5, 2, 2.5}},
		}

		for _, test := range tests {
			actual := Float64(test.start, test.end, FloatStep(test.floatStep))

			for i := range actual {
				assert.InDelta(t, test.expected[i], actual[i], 0.0001)
			}
		}
	})
}

func BenchmarkFloat64(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Float64(1, 1000000)
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
