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
  let getDepth = () => {}

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
