package booAnalysis

func wailingGhosts(sounds string) bool {
	o, u := 0, false
	for _, s := range sounds {
		if s == 'o' {
			if u {
				o--
				if 0 == o {
					u = false
				}
			} else {
				o++
			}
		}

		if s == 'u' {
			if o < 1 {
				return false
			}
			u = true
		}
	}

	return !u && o == 0
}

func booAnalysis(sounds string) int {

}
