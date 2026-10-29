func groupAnagrams(strs []string) [][]string {
	groupsCount := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, r := range str {
			count[r-'a']++
		}
		groupsCount[count] = append(groupsCount[count], str)
	}

	var result [][]string
	for _, group := range groupsCount {
		result = append(result, group)
	}

	return result
}



