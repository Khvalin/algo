package playlistlongestinterval

import (
	"regexp"
	"strconv"
)

func playlistLongestInterval(songs []string) (res int) {
	type entry struct {
		name     string
		duration int
	}

	total := 0

	plPos := map[string]int{}
	stack := []entry{}
	head := 0

	re := regexp.MustCompile(`^(\w+)\s\((\d?\d):(\d\d)\)$`)

	for _, s := range songs {
		matches := re.FindStringSubmatch(s)
		name := s //matches[1]
		m, _ := strconv.Atoi(matches[2])
		ss, _ := strconv.Atoi(matches[3])

		duration := m*60 + ss

		pos, found := plPos[name]
		if found {
			for k := head; k <= pos; k++ {
				total -= stack[k].duration
				delete(plPos, stack[k].name)
			}
			head = pos + 1
		}
		total += duration

		plPos[name] = len(stack)
		stack = append(stack, entry{name, duration})

		if res < total {
			res = total
		}
	}

	return res
}
