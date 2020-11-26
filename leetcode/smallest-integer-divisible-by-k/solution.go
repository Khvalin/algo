func smallestRepunitDivByK(K int) int {
	k := uint64(K)
	if k%2 == 0 {
		return -1
	}

	last := k % 10

	var start uint64
	for i := uint64(1); i < 10; i++ {
		if (last*i)%10 == 1 {
			start = i
			break
		}
	}
	if start == 0 {
		return -1
	}

	res := 1
	rem := start * k
	for rem > 0 {
		rem /= 10

		if rem == 0 {
			break
		}

		next := uint64(42)
		for i := uint64(0); i < 10; i++ {
			if (last*i+rem%10)%10 == 1 {
				next = i
				break
			}
		}

		if next > 9 {
			return -1
		}

		rem += next * k

		res++
	}

	return res
}