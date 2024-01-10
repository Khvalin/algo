/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} start
 * @return {number}
 */
var amountOfTime = function(root, start) {
	const parent = new Map();
	let startNode = root;
    const walk = (node) => {
		if (!node) {
			return;
		}
		
		if (node.val === start) {
			startNode = node;
		}
		
		if (node.left) {
			parent.set(node.left.val, node)
			walk(node.left)
		}
		
		if (node.right) {
			parent.set(node.right.val, node)
			walk(node.right)
		}
	}
	
	walk(root)
	let res = 0
	let visited = new Set();
	let q = [startNode]
	while (q.length) {
		const nodes = []
		for (const node of q) {
			visited.add(node.val)
			
			for (const c of [parent.get(node.val), node.left, node.right]) {
				if (c && !visited.has(c.val)) {
					nodes.push(c)
					visited.add(c.val)
				}
			}
		}
		
	
		res++;
		q = nodes
	}
	
	return res - 1
};
