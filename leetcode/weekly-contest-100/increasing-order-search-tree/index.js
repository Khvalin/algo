/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var increasingBST = function(root) {
  function inorder(node) {
    let result = []
    if (node) {
      return result.concat(inorder(node.left), [ node.val ], inorder(node.right))
    }
    return result
  }

  let nums = inorder(root)

  let cur = null
  for (let i = 0; i < nums.length; i++) {
    if (!cur) {
      cur = root
    } else {
      cur.right = { left: null, right: null }
      cur = cur.right
    }
    cur.right = null
    cur.left = null
    cur.val = nums[i]
  }

  return root
}
