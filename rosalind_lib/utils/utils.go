package utils

// HammingDistance calculates the number of differing
// characters between two strings s and t.
func HammingDistance(s, t string) int {
	if len(s) != len(t) {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			count++
		}
	}
	return count
}
