/**
 * @param {number} label
 * @return {number[]}
 */
var pathInZigZagTree = function(l) {
  function Node() {
    return {
      path: [],
      label: -1,
      left: {},
      right: {},
      parent: null
    }
  }
  var root = new Node()
  let n = 0

  let nodes = [ root ]
  let odd = true
  let res = null

  while (n < l) {
    const next = []
    for (const node of nodes) {
      n++
      node.label = n
      //console.log(node.path)

      if (n === l) {
        res = []
        let cur = node
        while (cur) {
          res.unshift(cur.label)
          cur = cur.parent
        }
        break
      }

      node.left = new Node()
      node.left.parent = node
      //   node.left.path = [ ...node.path ]

      node.right = new Node()
      node.right.parent = node
      //    node.right.path = [ ...node.path ]

      next.unshift(node.left, node.right)
    }

    nodes = next
    odd = !odd
  }

  //console.log(JSON.stringify(root))
  return res
}
/* 
console.log(pathInZigZagTree(14))

console.log(pathInZigZagTree(26)) */

console.log(pathInZigZagTree(196348))
