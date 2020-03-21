

//
// Binary trees are already defined with this interface:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
func deleteFromBST(t *Tree, queries []int) *Tree {
	for _, n := range queries {
		var parent *Tree
		node := t

		for node != nil && node.Value != n {
			parent = node
			if node.Value.(int) > n {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		if node == nil {
			continue
		}

		if node.Left == nil {
			if parent == nil {
				t = node.Right
				continue
			}
			if parent.Value.(int) > n {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
			continue
		}

		parent = node
		next := node.Left
		for next.Right != nil {
			parent = next
			next = next.Right
		}

		node.Value = next.Value
		if node.Left.Right != nil {
			parent.Right = next.Left
		} else {
			parent.Left = next.Left
		}

	}

	return t
}


