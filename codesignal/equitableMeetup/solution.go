package equitableMeetup

import (
	"container/heap"
)

type Friend struct {
	dist, fIndex, rIndex int
}

type FriendHeap []Friend

func (h FriendHeap) Len() int           { return len(h) }
func (h FriendHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h FriendHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FriendHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*h = old[0 : n-1]

	return item
}

func (h *FriendHeap) Push(x interface{}) {
	item := x.(Friend)
	*h = append(*h, item)
}

func equitableMeetup(friendsRoutes [][]int) []int {
	res := make([]int, len(friendsRoutes))
	for i := range friendsRoutes {
		res[i] = 1
	}

	maxDist := 0
	friendHeap := make(FriendHeap, len(friendsRoutes))
	for i, f := range friendsRoutes {
		friendHeap[i] = Friend{rIndex: 1, fIndex: i, dist: f[0]}
		if f[0] > maxDist {
			maxDist = f[0]
		}
	}
	heap.Init(&friendHeap)
	minDiff := maxDist - friendHeap[0].dist

	for len(friendHeap) > 1 {
		top := &(friendHeap[0])

		if (*top).rIndex >= len(friendsRoutes[(*top).fIndex]) {
			heap.Remove(&friendHeap, 0)
			break
		}
		(*top).dist += friendsRoutes[(*top).fIndex][(*top).rIndex]
		if (*top).dist > maxDist {
			maxDist = (*top).dist
		}
		(*top).rIndex++

		heap.Fix(&friendHeap, 0)
		diff := maxDist - friendHeap[0].dist
		if minDiff > diff {
			minDiff = diff

			for _, f := range friendHeap {
				res[f.fIndex] = f.rIndex
			}
		}

		//fmt.Println(friendHeap, minDiff)
	}

	return res
}
