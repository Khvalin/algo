import util from 'util'

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
var minCameraCover = function(root) {
  let res = 0

  const setCam = (node, parent) => {
    if (!node.hasCam) {
      res++
      node.covered = true
      let { left, right } = node

      parent && (parent.covered = true)
      right && (right.covered = true)
      left && (left.covered = true)
    }
  }

  const dfs = (node, parent) => {
    if (!node) {
      return
    }

    node.val = parent ? parent.val + 1 : 0

    let { left, right } = node

    left && dfs(left, node)
    right && dfs(right, node)

    if (!node.covered) {
      if ((!left || !left.covered) && (!right || !right.covered)) {
        setCam(parent || node, parent)
      }
    }
  }

  dfs(root)

  return res
}

//console.log(minCameraCover(createTree([ 0, 0, null, 0, 0 ])))
//console.log(minCameraCover(createTree([ 0 ])))
//console.log(minCameraCover(createTree([ 0, 0, 0, null, null, null, 0 ])))
console.log(2, minCameraCover(createTree([ 0, 0, null, null, 0, 0, null, null, 0, 0 ])))
console.log(3, minCameraCover(createTree([ 0, null, 0, 0, 0, null, null, null, 0, 0, null, 0, null, null, 0 ])))
