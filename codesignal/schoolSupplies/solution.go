package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func schoolSupplies(supplies []string, prices []int) int {
	re := regexp.MustCompile("([0-9]+) (.*$)")
	priceMap := map[string]int{}

	i, sum := 0, 0
	for _, pr := range supplies {
		res := re.FindAllStringSubmatch(pr, -1)
		count, _ := strconv.Atoi(res[0][1])
		prod := res[0][2]
		if 1 == count {
			prod += "s"
		}
		price, ok := priceMap[prod]

		if !ok {
			price = prices[i]
			priceMap[prod] = price
			i++
		}

		fmt.Println(prod, count, price)

		sum += count * price
	}

	return sum
}

func main() {
	fmt.Println(schoolSupplies([]string{"2 binders",
		"5 black pens",
		"3 notebooks",
		"1 calculator",
		"4 notebooks",
		"1 binder",
		"3 folders"}, []int{750, 150, 200, 1875, 250}))
}
