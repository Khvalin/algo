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
var sumRootToLeaf = function(root) {
  const MOD = 1000000007

  let a = []
  const walk = (node, n) => {
    n = n * 2
    n += node.val

    n %= MOD

    if (!node.left && !node.right) {
      a.push(n)
      return
    }

    node.left && walk(node.left, n)
    node.right && walk(node.right, n)
  }

  walk(root, 0)

  return a.reduce((n, sum) => {
    return (sum + n) % MOD
  }, 0)
}
