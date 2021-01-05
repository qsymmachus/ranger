package ranger

// Returns an interval slice of ints, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1.
// If the end value is smaller than the start value, returns an empty slice.
func Int(start, end int) []int {
	if start > end {
		return []int{}
	}

	vals := make([]int, (end-start)+1)
	for i := range vals {
		vals[i] = start
		start++
	}

	return vals
}

// Returns an interval slice of int8s, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1.
// If the end value is smaller than the start value, returns an empty slice.
func Int8(start, end int8) []int8 {
	if start > end {
		return []int8{}
	}

	vals := make([]int8, (end-start)+1)
	for i := range vals {
		vals[i] = start
		start++
	}

	return vals
}

// Returns an interval slice of int16s, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1.
// If the end value is smaller than the start value, returns an empty slice.
func Int16(start, end int16) []int16 {
	if start > end {
		return []int16{}
	}

	vals := make([]int16, (end-start)+1)
	for i := range vals {
		vals[i] = start
		start++
	}

	return vals
}
