package booAnalysis

func booAnalysis(sounds string) (r int) {
	sounds += "u"
	u := 0
	leftStreak, rightStreak := 0, 0

	save := func() {
		if rightStreak > 0 {
			if u > 0 {
				min := leftStreak
				if rightStreak < min {
					min = rightStreak
				}
				r += 2*min + u
			}
			u = 0
			leftStreak = 0
			rightStreak = 0
		}
	}

	for _, s := range sounds {
		if s == 'o' {
			if u > 0 {
				rightStreak++
				if leftStreak == rightStreak {
					save()
				}
			} else {
				leftStreak++
			}
		}

		if s == 'u' {
			if leftStreak > 0 {
				save()
				u++
			}
		}
	}

	return
}
