package twosums

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for k, v := range nums {
		r := target - v
		if k2, v2 := m[r]; v2 {
			return []int{k2, k}
		}
		m[v] = k
	}

	return nil
}
