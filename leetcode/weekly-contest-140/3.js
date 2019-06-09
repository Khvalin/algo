/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} limit
 * @return {TreeNode}
 */
var sufficientSubset = function(root, limit) {
  const calcDown = (node, cur) => {
    if (!node) {
      return
    }

    cur += node.val
    calcDown(node.left, cur)
    calcDown(node.right, cur)

    node.sum = cur
  }

  const walkDown = (node) => {
    if (!node) {
      return false
    }

    if (!node.left && !node.right) {
      return node.sum >= limit
    }

    let leftRes = walkDown(node.left)
    if (!leftRes) {
      node.left = null
    }

    let rightRes = walkDown(node.right)
    if (!rightRes) {
      node.right = null
    }

    let res = leftRes || rightRes

    return res
  }

  calcDown(root, 0)
  walkDown(root)

  if (!root.left && !root.right && root.val < limit) {
    root = null
  }

  return root
}
