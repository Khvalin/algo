package luckykids

import (
	"fmt"
	"sort"
)

// Kid is a kid
type Kid struct {
	age, behavior, naughtyEldersCount int
}

func (k Kid) String() string {
	return fmt.Sprintf("{ AGE: %v, BEH: %v, count: %v }", k.age, k.behavior, k.naughtyEldersCount)
}

//TreeNode is a tree node
type TreeNode struct {
	value *Kid
	left  *TreeNode //left
	right *TreeNode //right
}

func (n TreeNode) String() string {
	left, right := "", ""
	if n.left != nil {
		left = fmt.Sprintf("LEFT: %v", n.left)
	}
	if n.right != nil {
		right = fmt.Sprintf("RIGHT: %v", n.right)
	}

	return fmt.Sprintf("{ VAL: %v %v %v }", n.value, right, left)
}

//Insert is insert
func (n *TreeNode) Insert(value *Kid) error {

	//If the data value is greater than the current node's value,
	// do the same but for the right subtree.
	if value.behavior <= n.value.behavior {
		// "ages"  are distinct
		if value.behavior < n.value.behavior && value.age > n.value.age {
			n.value.naughtyEldersCount++
		}

		if n.right == nil {
			n.right = &TreeNode{value: value}
			return nil
		}
		return n.right.Insert(value)
	}

	//If the data value is less than the current node's value, and if the left child node is nil, insert a new left child node.
	// Else call Insert on the left subtree.
	if n.left == nil {
		n.left = &TreeNode{value: value}
		return nil
	}

	return n.left.Insert(value)
}

func luckyKids(behaviors []int) int {
	kids := []Kid{}
	L := len(behaviors)

	for i, beh := range behaviors {
		kids = append(kids, Kid{age: i, behavior: beh})
	}

	sort.Slice(kids, func(i, j int) bool {
		if kids[i].behavior == kids[j].behavior {
			return kids[i].age > kids[j].age
		}

		return kids[i].behavior < kids[j].behavior
	})

	pivot := L - 1
	for pivot > L>>1 && kids[pivot].behavior == kids[pivot-1].behavior {
		pivot--
	}
	root := TreeNode{value: &(kids[pivot])}
	for i := L - 1; i >= 0; i-- {
		if i != pivot {
			root.Insert(&(kids[i]))
		}
	}

	fmt.Println(root)

	res := 0
	for _, kid := range kids {
		minCount := (L - kid.age) >> 1
		//	fmt.Println(L, kid.naughtyEldersCount, kid.age, minCount)
		if kid.naughtyEldersCount >= minCount {
			res++
		}
	}

	return res
}
