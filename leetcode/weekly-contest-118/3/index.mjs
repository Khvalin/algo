//import '../../leetcode-tools'
import { createTree } from '../../leetcode-tools.mjs'

/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number[]} voyage
 * @return {number[]}
 */
var flipMatchVoyage = function(root, voyage) {
  const visit = (node, rightNodes = [], i = 0) => {
    node = node || rightNodes.shift()
    if ((node && i > voyage.length) || (!node && i < voyage.length)) {
      return null
    }

    //console.log(node.val, voyage[i])

    if ((!node || !node.val) && i >= voyage.length) {
      return []
    }

    if (node.val === voyage[i]) {
      let noSwap, swap

      if (node.left && node.right && node.right.val === voyage[i + 1]) {
        swap = visit(node.right, node.left ? [ node.left, ...rightNodes ] : [ ...rightNodes ], i + 1)
      } else {
        noSwap = visit(node.left, node.right ? [ node.right, ...rightNodes ] : [ ...rightNodes ], i + 1)
      }

      if (!noSwap) {
        return swap ? [ node.val, ...swap ] : null
      }

      if (!swap) {
        return noSwap || null
      }

      swap.unshift(node.val)

      return swap.length < noSwap.length ? swap : noSwap
    }

    return null
  }

  return visit(root) || [ -1 ]
}

/* console.log(flipMatchVoyage(createTree([ 1, 2, 3 ]), [ 1, 3, 2 ]))
console.log(flipMatchVoyage(createTree([ 1, 2, 3 ]), [ 1, 2, 3 ]))
console.log(flipMatchVoyage(createTree([ 1, 2 ]), [ 2, 1 ]))
 */
console.log(flipMatchVoyage(createTree([ 5, 1, 2, null, null, 4, 3 ]), [ 5, 2, 3, 4, 1 ]))

console.log(
  flipMatchVoyage(
    createTree([
      49,
      9,
      41,
      40,
      11,
      4,
      7,
      37,
      25,
      46,
      23,
      8,
      19,
      18,
      38,
      null,
      12,
      null,
      null,
      5,
      44,
      null,
      29,
      47,
      null,
      null,
      27,
      null,
      null,
      13,
      6,
      50,
      null,
      null,
      34,
      45,
      33,
      28,
      20,
      null,
      21,
      15,
      null,
      14,
      31,
      48,
      null,
      null,
      null,
      null,
      null,
      null,
      null,
      36,
      null,
      26,
      10,
      null,
      null,
      null,
      2,
      null,
      null,
      null,
      null,
      null,
      16,
      17,
      3,
      null,
      null,
      null,
      null,
      null,
      null,
      42,
      39,
      null,
      32,
      43,
      1,
      null,
      null,
      null,
      null,
      null,
      null,
      null,
      null,
      24,
      null,
      22,
      null,
      null,
      30,
      null,
      35
    ]),
    [
      49,
      41,
      7,
      38,
      6,
      48,
      17,
      1,
      22,
      35,
      43,
      24,
      30,
      3,
      13,
      31,
      16,
      32,
      14,
      18,
      4,
      19,
      27,
      15,
      8,
      47,
      21,
      2,
      42,
      39,
      9,
      40,
      25,
      37,
      12,
      50,
      11,
      23,
      29,
      20,
      28,
      26,
      10,
      46,
      5,
      34,
      44,
      33,
      36,
      45
    ]
  )
)
