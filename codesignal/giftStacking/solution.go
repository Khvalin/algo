package giftstacking

import (
	"sort"
)

func giftStacking(_boxes [][]int) int {
	type box struct {
		weight, strength, index int
	}

	memo := map[box]int{}

	boxes := []box{}
	for i, b := range _boxes {
		boxes = append(boxes, box{strength: b[0], weight: b[1], index: i})
	}

	sort.Slice(boxes, func(i, j int) bool {
		if boxes[i].strength == boxes[j].strength {
			return boxes[i].weight > boxes[j].weight
		}
		return boxes[i].strength < boxes[j].strength
	})

	//fmt.Println(boxes)

	var calc func(bottom box) int
	calc = func(bottom box) int {
		res, found := memo[bottom]
		if found {
			return res
		}

		pileHeight := 1
		for i := bottom.index + 1; i < len(boxes); i++ {
			minStrength := bottom.strength - boxes[i].weight
			if minStrength > boxes[i].strength {
				minStrength = boxes[i].strength
			}

			next := box{strength: minStrength, weight: boxes[i].weight, index: i}
			height := calc(next) + 1
			if height > pileHeight {
				pileHeight = height
			}
		}

		memo[bottom] = pileHeight

		return pileHeight
	}

	res := 1
	for _, b := range boxes {
		r, found := memo[b]
		if !found {
			r = calc(b)
		}

		if r > res {
			res = r
		}
	}

	return res
}
