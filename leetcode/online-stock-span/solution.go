package main

import "fmt"

type StockSpanner struct {
	stocks [][2]int32
}

func Constructor() StockSpanner {
	nmax := 10000
	spanner := StockSpanner{
		stocks: make([][2]int32, 1, nmax),
	}
	spanner.stocks[0] = [2]int32{100001, 0}

	return spanner
}

func (this *StockSpanner) Next(price int) int {
	p := int32(price)
	var res int32 = 1
	i := int32(len(this.stocks) - 1)
	for i > 0 && this.stocks[i][0] <= p {
		res += this.stocks[i][1]
		i -= this.stocks[i][1]
	}

	this.stocks = append(this.stocks, [2]int32{p, res})

	return int(res)
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
