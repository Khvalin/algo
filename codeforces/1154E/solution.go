package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Node ..
type Node struct {
	next *Node
	prev *Node
	val  uint
}

func solve(k int, a []uint) []byte {
	removeNode := func(node *Node) {
		if node == nil {
			return
		}

		prev, next := node.prev, node.next
		if prev != nil {
			prev.next = next
		}

		if next != nil {
			next.prev = prev
		}
	}

	L := len(a)

	indices := map[uint]int{}
	for i, n := range a {
		indices[n] = i
	}

	res := make([]byte, L)
	sorted := make([]uint, L)
	copy(sorted, a)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	orig := map[uint]*Node{}
	var p *Node
	for _, n := range a {
		node := Node{val: n, prev: p}
		if p != nil {
			(*p).next = &node
		}
		p = &node

		orig[n] = &node
	}

	i := L
	team := 1
	for i > 0 {
		i--
		cur := indices[sorted[i]]
		if res[cur] > 0 {
			continue
		}

		t := byte(team)

		node := orig[sorted[i]]
		prev, next := node.prev, node.next
		for j := 0; j < k; j++ {
			if prev != nil {
				res[indices[prev.val]] = t
				removeNode(prev)
				prev = prev.prev
			}

			if next != nil {
				res[indices[next.val]] = t
				removeNode(next)
				next = next.next
			}
		}

		res[cur] = t
		removeNode(node)

		if team == 1 {
			team = 2
		} else {
			team = 1
		}
	}

	return res
}

func toInt(buf []byte) uint {
	n := uint(0)
	for _, v := range buf {
		n = n*10 + uint(v-'0')
	}

	return n
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	a := make([]uint, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := range a {
		scanner.Scan()
		a[i] = toInt(scanner.Bytes())
	}

	ans := solve(k, a)
	for _, n := range ans {
		fmt.Print(n)
	}

	fmt.Println()
}
