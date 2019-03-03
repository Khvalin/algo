package powersOfPrime23

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func powersOfPrime23(p int) (r []int) {
	p2 := 1
	for i := 0; i < 100 && p2 > 0; i++ {
		p3 := 1
		for j := 0; j < 100 && p3 > 0; j++ {
			if abs(p2-p3) == p {
				return []int{i, j}
			}
			p3 *= 3
		}
		p2 *= 2
	}

	return
}
