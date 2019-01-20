import { createTree } from '../leetcode-tools.mjs'

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
var distributeCoins = function(root) {
  const dfs = (node) => {
    if (node === null) {
      return
    }
    dfs(node.left)
    dfs(node.right)
    if (node.val == 0) {
    }
  }
}

console.log(distributeCoins(createTree([ 3, 0, 0 ])))
