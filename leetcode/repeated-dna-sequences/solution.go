func findRepeatedDnaSequences(s string) []string {
	m := map[uint64] int{}
	var res []string
	var  sum uint64
	
	for i := 0; i < len(s); i++{
			ch := s[i]

			var d uint64
			switch(ch) {
					case ('C'):
							d = 1
					case ('G'):
							d = 2
					case ('T'):
							d = 3
			}
			sum = sum << 2
			sum += d
			sum %= 1 << 20
			
			if i >= 9 {
					fmt.Println(sum, s[i-9:i+1])
					m[sum]++
					if  m[sum] == 2 {
							res = append(res, s[i-9:i+1])
					}
			}
	}

	return res
}