/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var balanceBST = function (root) {
  const toSortedArray = (node) => {
    if (!node) {
      return [];
    }

    return [
      ...toSortedArray(node.left),
      node.val,
      ...toSortedArray(node.right),
    ];
  };

  const buildBalancedBST = (a, left, right) => {
    if (left >= right) {
      return null;
    }

    const mid = (left + right) >> 1;

    const root = new TreeNode(a[mid]);
    root.left = buildBalancedBST(a, left, mid);
    root.right = buildBalancedBST(a, mid + 1, right);

    return root;
  };

  const a = toSortedArray(root);

  return buildBalancedBST(a, 0, a.length);
};
