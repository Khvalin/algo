package theSmallestStringCipher

func theSmallestStringCipher(key string, message string) string {
	result := ""

	i := 0
	j := 0
	for i < len(key) || j < len(message) {
		pushMessageChar := false
		if j >= len(message) ||
			(i < len(key) && (key[i] <= message[j])) {
			if j < len(message) && i < len(key) && key[i] == message[j] {
				i1 := i
				j1 := j

				for i1 < len(key) && j1 < len(message) && key[i1] == message[j1] {
					i1++
					j1++
				}

				pushMessageChar = i1 >= len(key) || j1 < len(message) && message[j1] < key[i1]
			}
		} else {
			pushMessageChar = true
		}
		if pushMessageChar {
			result += string(message[j])
			j++
		} else {
			result += string(key[i])
			i++
		}
	}

	return result
}
