package equitableMeetup

//import "container/heap";

type Friend struct {
	dist, index int
}

type FriendHeap []Friend

func (h FriendHeap) Len() int           { return len(h) }
func (h FriendHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h FriendHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func equitableMeetup(friendsRoutes [][]int) []int {
	res := []int{}
	return res
}
