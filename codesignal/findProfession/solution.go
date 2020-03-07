func findProfession(l, p int) (r string) {
	r = "Doctor"
	if l > 1 && findProfession(l-1, p/2+p%2) == r {
		p++
	}

	if p%2 > 0 {
		r = "Engineer"
	}

	return
}
