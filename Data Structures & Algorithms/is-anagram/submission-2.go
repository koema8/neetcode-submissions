func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	temp := make(map[rune]int, len(s))

	for _, r := range s {
		temp[r]++
	}

	for _, r := range t {
		temp[r]--
		if temp[r] < 0 {
			return false
		}
	}

	return true
}