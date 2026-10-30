func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// freq = index
	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets)-1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	return result[:k]
}