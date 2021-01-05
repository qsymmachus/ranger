package ranger

// Returns an interval slice of ints, with the specified start and end value.
// Values between the start and end are incremented ("stepped") by 1.
//
//   ints := ranger.Int(1, 5)
//   fmt.Printf("%v", ints) // [1 2 3 4 5]
//
// If the end value is smaller than the start, returns an empty slice.
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
//
//   int8s := ranger.Int8(4, 8)
//   fmt.Printf("%v", int8s) // [4 5 6 7 8]
//
// If the end value is smaller than the start, returns an empty slice.
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
