package ranger

// Encapsulates options that are applied to functions that generate intervals.
type options struct {
	step int
}

// An Option function mutates the state of the underlying options struct.
type Option func(*options)

// Applies a list of option functions to the underlying options struct.
func (o *options) mergeOptions(opts ...Option) {
	// Initialize defaults
	o.step = 1

	for _, opt := range opts {
		opt(o)
	}
}

// The Step option changes the increment ("step") between values in generated
// intervals. If this option is omitted, the step defaults to 1.
func Step(size int) func(*options) {
	return func(opts *options) {
		opts.step = size
	}
}

// Returns an interval slice of ints, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1, unless a
// different step option is provided. If the end value is smaller than the
// start value, returns an empty slice.
func Int(start, end int, opts ...Option) []int {
	if start > end {
		return []int{}
	}

	options := options{}
	options.mergeOptions(opts...)
	vals := make([]int, ((end-start)/options.step)+1)
	for i := range vals {
		vals[i] = start
		start += options.step
	}

	return vals
}

// Returns an interval slice of int8s, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1, unless a
// different step option is provided. If the end value is smaller than the
// start value, returns an empty slice.
func Int8(start, end int8, opts ...Option) []int8 {
	if start > end {
		return []int8{}
	}

	options := options{}
	options.mergeOptions(opts...)
	vals := make([]int8, ((end-start)/int8(options.step))+1)
	for i := range vals {
		vals[i] = start
		start += int8(options.step)
	}

	return vals
}

// Returns an interval slice of int16s, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1, unless a
// different step option is provided. If the end value is smaller than the
// start value, returns an empty slice.
func Int16(start, end int16, opts ...Option) []int16 {
	if start > end {
		return []int16{}
	}

	options := options{}
	options.mergeOptions(opts...)
	vals := make([]int16, ((end-start)/int16(options.step))+1)
	for i := range vals {
		vals[i] = start
		start += int16(options.step)
	}

	return vals
}
