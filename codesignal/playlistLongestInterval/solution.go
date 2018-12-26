package playlistlongestinterval

import (
	"fmt"
	"regexp"
	"strconv"
)

func playlistLongestInterval(songs []string) (res int) {
	total := 0
	songs = append(songs, "PLAYLISTEND (00:00)")
	songDuration := map[string]int{}
	re := regexp.MustCompile(`^(\w+)\s\((\d?\d):(\d\d)\)$`)
	for _, s := range songs {
		matches := re.FindStringSubmatch(s)
		name := matches[1]
		h, _ := strconv.Atoi(matches[2])
		m, _ := strconv.Atoi(matches[3])

		_, found := songDuration[name]

		fmt.Println(name, total)

		if res < total {
			res = total
		}

		d := h*60 + m

		total += d

		if found {
			fmt.Println("  clear")
			total -= songDuration[name]
			songDuration = map[string]int{}
		}
		songDuration[name] = total
	}

	fmt.Println(songDuration)

	return res
}
