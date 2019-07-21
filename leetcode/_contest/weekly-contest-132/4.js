// Definition for a binary tree node.
function TreeNode(val) {
  this.val = val
  this.left = this.right = null
}

/**
 * @param {string} S
 * @return {TreeNode}
 */
var recoverFromPreorder = function(S) {
  const readNum = (j) => {
    let n = 0
    while (S[j] <= '9' && S[j] >= '0') {
      n *= 10
      n += S[j] - '0'
      j++
    }

    return [ n, j ]
  }

  let [ n, i ] = readNum(0)
  const root = new TreeNode(n)
  const nodes = { 0: [ root ] }

  //i++
  while (i < S.length) {
    let d = 0
    let j = i
    while (S[j] == '-') {
      j++
      d++
    }

    ;[ n, j ] = readNum(j)

    const node = new TreeNode(n)
    nodes[d] = nodes[d] || []
    nodes[d].push(node)

    //const len = nodes[d].length - 1
    //  console.log(j, n, d)
    const p = nodes[d - 1][nodes[d - 1].length - 1]

    if (p.left) {
      p.right = node
    } else {
      p.left = node
    }

    i = j
  }

  // console.log(nodes)

  return root
}

//console.log(recoverFromPreorder('1-2--3--4-5--6--7'))
let r = JSON.stringify(recoverFromPreorder('10-7--8'), null, 2)
console.log(r)
