// Range functions for Go. Easily generate intervals of integers, floats, runes, and strings.
package ranger

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Options applied to all range functions in this package.
type options[N Number] struct {
	step N
}

// An Option is just a function that changes the provided options struct.
// All range functions in this package accept Options as variadic arguments,
// so you can pass (optional!) options like step size.
type Option[N Number] func(*options[N])

// The Step option changes the increment ("step") between values in generated
// intervals. If this option is omitted, the step defaults to 1.
func Step[N Number](size N) func(*options[N]) {
	return func(opts *options[N]) {
		opts.step = size
	}
}

// Applies a list of option functions to the underlying options struct.
func (o *options[N]) mergeOptions(opts ...Option[N]) {
	// Initialize defaults
	o.step = 1

	for _, opt := range opts {
		opt(o)
	}
}

// Returns an interval slice of numbers, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1, unless a
// different step option is provided. If the end value is smaller than the
// start value, returns an empty slice.
func Range[N Number](start, end N, opts ...Option[N]) []N {
	if start > end {
		return []N{}
	}

	options := options[N]{}
	options.mergeOptions(opts...)
	vals := make([]N, int(((end-start)/options.step)+1))
	for i := range vals {
		vals[i] = start
		start += options.step
	}

	return vals
}

// Returns an interval slice of runes, with the specified start and end value.
// Runes are incremented by adding an integer step value to the rune's integer
// value, and casting that value back into a rune. The step value is 1, unless
// a different step option is provided. If the end value is smaller than the
// start value, returns an empty slice.
func Rune(start, end rune, opts ...Option[int]) []rune {
	if start > end {
		return []rune{}
	}

	options := options[int]{}
	options.mergeOptions(opts...)
	vals := make([]rune, ((int(end)-int(start))/options.step)+1)
	for i := range vals {
		vals[i] = start
		start = rune(int(start) + options.step)
	}

	return vals
}

// Returns an interval slice of strings. Under the hood, this generates an
// interval of runes and casts them to strings. For that reason, starting
// and ending strings cannot be empty and must contain a single character,
// or an error is returned. The step value is 1, unless a different step
// option is provided. If the end value is smaller than the start value,
// returns an empty slice.
func String(start, end string, opts ...Option[int]) ([]string, error) {
	startRunes := []rune(start)
	if len(startRunes) > 1 {
		return []string{}, fmt.Errorf("Start string '%s' contains more than one character", start)
	} else if len(startRunes) == 0 {
		return []string{}, fmt.Errorf("Start string is empty")
	}

	endRunes := []rune(end)
	if len(endRunes) > 1 {
		return []string{}, fmt.Errorf("End string '%s' contains more than one character", start)
	} else if len(endRunes) == 0 {
		return []string{}, fmt.Errorf("End string is empty")
	}

	runeVals := Rune(startRunes[0], endRunes[0], opts...)
	stringVals := make([]string, len(runeVals))
	for i, r := range runeVals {
		stringVals[i] = string(r)
	}

	return stringVals, nil
}
