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
var minCameraCover = function(root) {
  const countAndSetCams = (node, setCam) => {
    let c = 0
    if (node) {
      c = 1 * setCam + countAndSetCams(node.left, !setCam) + countAndSetCams(node.right, !setCam)
    }

    return c
  }

  let min = 0
  if (root) {
    min = 1
  }

  return Math.max(min, Math.min(countAndSetCams(root, false), countAndSetCams(root, true)))
}
