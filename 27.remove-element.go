func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i, j := 0, 0
	for j < n {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
