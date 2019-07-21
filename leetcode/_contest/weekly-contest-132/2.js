/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxAncestorDiff = function(root) {
  let res = -1
  const walk = (node, min, max) => {
    if (node == null) {
      return
    }

    min = Math.min(min, node.val)
    max = Math.max(max, node.val)

    res = Math.max(res, Math.abs(min - max))

    walk(node.left, min, max)
    walk(node.right, min, max)
  }

  walk(root, root.val, root.val)

  return res
}
