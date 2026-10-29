func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))

	for idx, num := range nums {
		if index, ok := sumMap[num]; ok{
			return []int{index, idx}
		}
		sumMap[target - num] = idx
	}

	return nil
}

