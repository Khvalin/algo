package main

import (
	"fmt"
	"reflect"
)

func longestConsecutive(nums []int) int {
	has := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		has[nums[i]] = true
	}

	max := 0

	for len(has) > 0 {
		// Get any key from map:
		keys := reflect.ValueOf(has).MapKeys()
		key := keys[0].Interface().(int)
		delete(has, key)

		count := 1
		for l, ok := key-1, true; ok; l-- {
			_, ok = has[l]
			if ok {
				delete(has, l)
				count++
			}
		}

		for r, ok := key+1, true; ok; r++ {
			_, ok = has[r]
			if ok {
				delete(has, r)
				count++
			}
		}

		if count > max {
			max = count
		}
	}

	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

	fmt.Println(longestConsecutive([]int{}))
}
