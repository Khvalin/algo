/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isUnivalTree = function(root) {
  const vals = new Set()
  const visit = (node) => {
    if (node) {
      vals.add(node.val)
      visit(node.left)
      visit(node.right)
    }
  }
  visit(root)

  return vals.size < 2
}
