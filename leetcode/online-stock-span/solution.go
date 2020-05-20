type StockSpanner struct {
    stocks []int
    spans  []int
    last   []int
}


func Constructor() StockSpanner {
    return StockSpanner{
        stocks: []int{10e5+5},
        spans:  []int{1},
        last:   []int{0},
    }
}


func (this *StockSpanner) Next(price int) int {
    for i := len(this.stocks); i > 0; i-- {
        
    }
    return 1
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
