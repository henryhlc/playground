package algo

// Implementing the KMP string search algorithm [1].
// This function returns the index of the start of at most n
// occurrances of substr in s. When n <= 0, this returns
// all occurrances, which can be called more conveniently
// with AllSubstrOccurrances.
//
// [1] https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
func FirstNSubstrOccurrances(s, substr string, n int) []int {
	jt := make([]int, len(substr)+1)
	jt[0] = -1
	for i := 1; i < len(jt); i++ {
		next := jt[i-1]
		for next >= 0 && substr[next] != substr[i-1] {
			next = jt[next]
		}
		jt[i] = next + 1
	}

	occurrances := []int{}
	next := 0
	for i := range len(s) {
		for next >= 0 && substr[next] != s[i] {
			next = jt[next]
		}
		next++
		if next == len(substr) {
			occurrances = append(occurrances, i-len(substr)+1)
			if n > 0 && len(occurrances) == n {
				return occurrances
			}
			next = jt[next]
		}
	}
	return occurrances
}

// The index of the start of the first occurance of substr in s.
// If substr is not in s, returns -1.
func FirstSubstrOccurrance(s, substr string) int {
	occurrance := FirstNSubstrOccurrances(s, substr, 1)
	if len(occurrance) == 0 {
		return -1
	}
	return occurrance[0]
}

// The indices of the start of the occurrances of substr in s.
func AllSubstrOccurrances(s, substr string) []int {
	return FirstNSubstrOccurrances(s, substr, 0)
}
