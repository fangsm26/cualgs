package algs

// package algs contains some commonly used helper functions.

// Min returns the minimum value of two integers
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Max returns the maximum value of two integers
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
