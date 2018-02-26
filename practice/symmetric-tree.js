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

var isSymmetric = function(root) {
  if (root) {
    let leftQ = [ root.left ];
    let rightQ = [ root.right ];

    while (leftQ.length > 0 || rightQ.length > 0) {
      let leftNode = leftQ.shift();
      let leftVal = leftNode && leftNode.val;

      let rightNode = rightQ.shift();
      let rightVal = rightNode && rightNode.val;

      if (rightVal !== leftVal) {
        return false;
      }

      if (leftNode) {
        leftQ.push(leftNode.left);
        leftQ.push(leftNode.right);
      }

      if (rightNode) {
        rightQ.push(rightNode.right);
        rightQ.push(rightNode.left);
      }
    }
  }

  return true;
};
