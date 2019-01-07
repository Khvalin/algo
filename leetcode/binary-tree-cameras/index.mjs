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
var minCameraCover = function(node, needsCam = false) {
  if (null == node) {
    return needsCam ? 9999 : 0
  }

  let withCam = 1 + minCameraCover(node.left, false) + minCameraCover(node.right, false)
  if (needsCam) {
    return withCam
  }

  let noCam = Math.min(
    minCameraCover(node.left, true) + minCameraCover(node.right, false),
    minCameraCover(node.left, false) + minCameraCover(node.right, true)
  )

  return Math.min(withCam, noCam)
}

console.log(minCameraCover(createTree([ 0, 0, null, 0, 0 ])))
