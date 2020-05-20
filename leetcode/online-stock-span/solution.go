package main

import "fmt"

type StockSpanner struct {
	stocks []int
	spans  []int
	last   []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stocks: []int{10e5 + 5},
		spans:  []int{1},
		last:   []int{0},
	}
}

func (this *StockSpanner) Next(price int) int {
	for i := len(this.stocks); i > 0; i-- {

	}
	return 1
}

func main() {
	obj := Constructor()
	res := []int{}
	for _, price := range []int{100, 80, 60, 70, 60, 75, 85} {
		res = append(res, obj.Next(price))
	}

	fmt.Println(res)
	fmt.Println([]int{1, 1, 1, 2, 1, 4, 6})
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
