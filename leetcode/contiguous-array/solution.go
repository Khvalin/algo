func findMaxLength(nums []int) int {
	m := map[int] int{0:-1}
	s, r := 0, 0

	for i := range nums {
			s++
			if nums[i] == 0 {
					s -= 2
			}

			if v, f := m[s]; f {
					if r < i - v {
							r = i - v
					}
			} else {
					m[s] = i
			}
	}

	return r
}