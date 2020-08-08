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
 * @param {number} sum
 * @return {number}
 */
var pathSum = function(root, sum) {
  let res = 0
  const walk = (node, sums) => {
    if (!node) {
      return
    }

    const c = new Map()

    for (const [k, v] of sums.entries()) {
      c.set(k + node.val, v)
    }
    c.set(node.val, 1 + (c.get(node.val) || 0) )
    
    res += c.get(sum) || 0

    if (node.left) {
      walk(node.left, c)
    }

    if (node.right) {
      walk(node.right, c)
    }
  }
  
  walk(root, new Map())
  
  return res
};


