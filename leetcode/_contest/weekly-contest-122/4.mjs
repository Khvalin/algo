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
 * @return {number[][]}
 */
var verticalTraversal = function(root) {
  let nodes = []
  const walk = (node, x = 0, y = 0) => {
    if (node) {
      nodes.push({
        v: node.val,
        x,
        y
      })

      node.left && walk(node.left, x - 1, y - 1)
      node.right && walk(node.right, x + 1, y - 1)
    }
  }

  walk(root)

  nodes.sort((a, b) => {
    if (a.x === b.x) {
      if (a.y === b.y) {
        return a.v - b.v
      }
      return b.y - a.y
    }

    return a.x - b.x
  })

  const res = []

  for (let i = 0; i < nodes.length; i++) {
    if (i == 0 || nodes[i].x !== nodes[i - 1].x) {
      res.push([ nodes[i].v ])
    } else {
      res[res.length - 1].push(nodes[i].v)
    }
  }

  return res
}

console.log(verticalTraversal(createTree([ 0, 2, 1, 3, null, null, null, 4, 5, null, 7, 6, null, 10, 8, 11, 9 ])))
