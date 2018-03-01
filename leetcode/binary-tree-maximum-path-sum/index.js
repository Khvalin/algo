'use strict'

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
var maxPathSum = function(root) {
  let result = Number.NEGATIVE_INFINITY

  const visit = function(root) {
    if (!root) {
      return Number.NEGATIVE_INFINITY
    }

    if ('weigth' in root) {
      return root.weigth
    }

    const leftSubtreeMaxSum = visit(root.left)
    const rightSubtreeMaxSum = visit(root.right)

    root.weigth = root.val + Math.max(0, leftSubtreeMaxSum, rightSubtreeMaxSum)

    result = Math.max(result, root.weigth, leftSubtreeMaxSum + rightSubtreeMaxSum + root.val)

    //    console.log('--', leftSubtreeMaxSum, rightSubtreeMaxSum)
    //    console.log('-- --', root.val, root.weigth, result)
    return root.weigth
  }

  visit(root)

  return result
}

module.exports = { maxPathSum }
