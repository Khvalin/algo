/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {boolean}
 */
var flipEquiv = function(root1, root2) {
  if (root1.val !== root2.val) {
    return false
  }

  if (!root1.left && !root1.right && !root2.left && !root2.right) {
    return true
  }

  const [ l1, r1, l2, r2 ] = [ root1.left || {}, root1.right || {}, root2.left || {}, root2.right || {} ]
  if (l1.val === l2.val && r1.val === r2.val) {
    return flipEquiv(l1, l2) && flipEquiv(r1, r2)
  }

  if (l1.val === r2.val && r1.val === l2.val) {
    return flipEquiv(l1, r2) && flipEquiv(r1, l2)
  }

  return false
}
