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
var leafSimilar = function(root1, root2) {
  const leafValueSequence = (node) => {
    if (!node) {
      return []
    }

    if (node.left || node.right) {
      return [].concat(leafValueSequence(node.left), leafValueSequence(node.right))
    }

    return [ node.val ]
  }

  return leafValueSequence(root1).join(',') === leafValueSequence(root2).join(',')
}

console.log(leafSimilar())
