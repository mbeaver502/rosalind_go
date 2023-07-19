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

// FindAllSubstrings finds all overlapping instances
// of the substring t in the string s.
// The zero-based indices are returned as a slice of int.
func FindAllSubstrings(s, t string) []int {
	result := []int{}
	k := len(t)
	for i := 0; i < len(s)-k; i++ {
		if s[i:i+k] == t {
			result = append(result, i)
		}
	}
	return result
}
