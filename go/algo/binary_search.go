package algo

/*
Returns
  - low: index to the last element where cond is false,
    or -1 if cond(e) is false for all elements of s.
  - high: index to the first element where cond is true,
    or len(s) if cond(e) is true for all elements of s.
*/
func BinarySearch[E any](s []E, cond func(E) bool) (int, int) {
	low := -1
	high := len(s)
	for high-low > 1 {
		mid := (high + low) / 2
		if cond(s[mid]) {
			high = mid
		} else {
			low = mid
		}
	}
	return low, high
}
