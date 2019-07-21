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
var subtreeWithAllDeepest = function(root) {
  let getDepth = (root) => {
    if (!root) {
      return 0
    }
    return Math.max(getDepth(root.left), getDepth(root.right)) + 1
  }

  var lDepth = getDepth(root.left)
  var rDepth = getDepth(root.right)

  if (lDepth === rDepth) {
    return root
  } else {
    if (lDepth > rDepth) {
      return subtreeWithAllDeepest(root.left)
    } else {
      return subtreeWithAllDeepest(root.right)
    }
  }
}

var tree = require('leetcode').Tree

var t = tree.create([ 3, 5, 1, 6, 2, 0, 8, null, null, 7, 4 ])
console.log(subtreeWithAllDeepest(t))
