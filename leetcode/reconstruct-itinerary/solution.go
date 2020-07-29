package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	adj := map[string][]string{}

	for _, pair := range tickets {
		adj[pair[0]] = append(adj[pair[0]], pair[1])
	}

	for k := range adj {
		sort.Strings(adj[k])
	}

	used := map[string]bool{}
	pos := map[string]int{}
	path := []string{"JFK"}

	outer:
	for len(path) > 0 {
		l := len(path)
		if l > len(tickets) {
			return path
		}
			fmt.Println(path)

		cur := path[l-1]
		i := pos[cur]
		for i < pos[cur]+len(adj[cur]) {
			next := adj[cur][i%len(adj[cur])]
			if !used[cur+next] {
				used[cur+next] = true
				path = append(path, next)
				pos[cur]=  i+1

				continue outer
			}

			i++
		}

		pos[cur] = 0
		l--
		path = path[:l]
		if len(path) > 0 {
			used[path[l-1]+cur] = false
			//		pos[path[l-1]]++
		}
	}

	return nil
}

func main() {
	var tickets [][]string
	// tickets = [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	// fmt.Println(findItinerary(tickets))

	// tickets = [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	// fmt.Println(findItinerary(tickets))

	tickets = [][]string{{"AXA", "EZE"}, {"EZE", "AUA"}, {"ADL", "JFK"}, {"ADL", "TIA"}, {"AUA", "AXA"}, {"EZE", "TIA"}, {"EZE", "TIA"}, {"AXA", "EZE"}, {"EZE", "ADL"}, {"ANU", "EZE"}, {"TIA", "EZE"}, {"JFK", "ADL"}, {"AUA", "JFK"}, {"JFK", "EZE"}, {"EZE", "ANU"}, {"ADL", "AUA"}, {"ANU", "AXA"}, {"AXA", "ADL"}, {"AUA", "JFK"}, {"EZE", "ADL"}, {"ANU", "TIA"}, {"AUA", "JFK"}, {"TIA", "JFK"}, {"EZE", "AUA"}, {"AXA", "EZE"}, {"AUA", "ANU"}, {"ADL", "AXA"}, {"EZE", "ADL"}, {"AUA", "ANU"}, {"AXA", "EZE"}, {"TIA", "AUA"}, {"AXA", "EZE"}, {"AUA", "SYD"}, {"ADL", "JFK"}, {"EZE", "AUA"}, {"ADL", "ANU"}, {"AUA", "TIA"}, {"ADL", "EZE"}, {"TIA", "JFK"}, {"AXA", "ANU"}, {"JFK", "AXA"}, {"JFK", "ADL"}, {"ADL", "EZE"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"ADL", "EZE"}, {"JFK", "ADL"}, {"ADL", "AXA"}, {"TIA", "AUA"}, {"AXA", "JFK"}, {"ADL", "AUA"}, {"TIA", "JFK"}, {"JFK", "ADL"}, {"JFK", "ADL"}, {"ANU", "AXA"}, {"TIA", "AXA"}, {"EZE", "JFK"}, {"EZE", "AXA"}, {"ADL", "TIA"}, {"JFK", "AUA"}, {"TIA", "EZE"}, {"EZE", "ADL"}, {"JFK", "ANU"}, {"TIA", "AUA"}, {"EZE", "ADL"}, {"ADL", "JFK"}, {"ANU", "AXA"}, {"AUA", "AXA"}, {"ANU", "EZE"}, {"ADL", "AXA"}, {"ANU", "AXA"}, {"TIA", "ADL"}, {"JFK", "ADL"}, {"JFK", "TIA"}, {"AUA", "ADL"}, {"AUA", "TIA"}, {"TIA", "JFK"}, {"EZE", "JFK"}, {"AUA", "ADL"}, {"ADL", "AUA"}, {"EZE", "ANU"}, {"ADL", "ANU"}, {"AUA", "AXA"}, {"AXA", "TIA"}, {"AXA", "TIA"}, {"ADL", "AXA"}, {"EZE", "AXA"}, {"AXA", "JFK"}, {"JFK", "AUA"}, {"ANU", "ADL"}, {"AXA", "TIA"}, {"ANU", "AUA"}, {"JFK", "EZE"}, {"AXA", "ADL"}, {"TIA", "EZE"}, {"JFK", "AXA"}, {"AXA", "ADL"}, {"EZE", "AUA"}, {"AXA", "ANU"}, {"ADL", "EZE"}, {"AUA", "EZE"}}
	fmt.Println(findItinerary(tickets))

	tickets = [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}
	fmt.Println(findItinerary(tickets))
}
