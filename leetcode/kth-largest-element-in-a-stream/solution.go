import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Peek() int     { return h[0] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
    c int
    pq *IntHeap
}


func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{ }
	heap.Init(h)
    
    for _, val := range nums { // O(n)
		heap.Push(h, val) // O(log n)
	}
    
    for h.Len() > k {
        heap.Pop(h)
    }
    
    return KthLargest{c:k, pq:h}
}


func (this *KthLargest) Add(val int) int {
    if this.pq.Len() < this.c {
        heap.Push(this.pq, val)
    } else {
        if val >  this.pq.Peek() {
            heap.Pop(this.pq)   
            heap.Push(this.pq, val)
        }
    }
    
    return this.pq.Peek()
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */