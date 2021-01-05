package ranger

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
