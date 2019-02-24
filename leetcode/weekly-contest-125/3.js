/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} val
 * @return {TreeNode}
 */

function TreeNode(val) {
  this.val = val
  this.left = this.right = null
}

var insertIntoMaxTree = function(root, val) {
  if (!root) {
    return new TreeNode(val)
  }

  if (root.val < val) {
    const res = new TreeNode(val)
    res.left = root
    return res
  }

  root.right = insertIntoMaxTree(root.right, val)
  return root
}
