package holidayPotluck

import "fmt"
import "sort"

func holidayPotluck(_dishes [][]int, capacity int) int {
	type dish struct {
		quality, quantity int
	}

	L := len(_dishes)

	dishes := make([]dish, L)
	for i, arr := range _dishes {
		dishes[i] = dish{quality: arr[0], quantity: arr[1]}
	}

	sort.Slice(dishes, func(i, j int) bool {
		return dishes[i].quality > dishes[j].quality
	})

	fmt.Println(dishes)

	count, res := 0, 0
	for i := 0; i < L && count < capacity; i++ {
		d := dishes[i]
		c := capacity - count
		if d.quantity < c {
			c = d.quantity
		}
		res += c * d.quality
		count += c
	}

	return res
}
