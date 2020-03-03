//
// Binary trees are already defined with this interface:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
func kthSmallestInBST(t *Tree, k int) int {
	r := 0

	var walk func(t *Tree, c int) int
	walk = func(t *Tree, c int) int {
		if c == k {
			r = t.Value.(int)
		}

		if t.Left != nil {
			c = walk(t.Left, c)
			c++
		}

		if c == k {
			r = t.Value.(int)
		}

		if t.Right != nil {
			c = walk(t.Right, c+1)
		}

		return c
	}

	walk(t, 1)

	return r
}
