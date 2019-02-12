package subarraysWithKDistinct

func subarraysWithKDistinct(A []int, K int) int {
	calc := func(start, index int) int {
		c := index - start + 1

		if c >= K {
			d := c - K + 1
			return d * (d + 1) >> 1
		}
		return 0
	}
	res := 0
	m := map[int]int{}

	i, s := 0, 0
	for i < len(A) {

		m[A[i]] = i

		if len(m) > K {
			res += calc(s, i-1)

			for len(m) > K {
				if m[A[s]] == s {
					delete(m, A[s])
				}
				s++
			}
		}
		i++
	}

	if len(m) == K {
		res += calc(s, i-1)
	}

	return res
}
