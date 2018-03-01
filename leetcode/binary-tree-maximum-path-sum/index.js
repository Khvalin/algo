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
  if (!root) {
    return Number.NEGATIVE_INFINITY;
  }

  if ('weigth' in root) {
    return root.weigth;
  }

  const leftSubtreeMaxSum = maxPathSum(root.left);
  const rightSubtreeMaxSum = maxPathSum(root.right);

  root.weigth = root.val + Math.max(0, leftSubtreeMaxSum, rightSubtreeMaxSum, leftSubtreeMaxSum + rightSubtreeMaxSum);

  return root.weigth;
};

module.exports = { maxPathSum };
