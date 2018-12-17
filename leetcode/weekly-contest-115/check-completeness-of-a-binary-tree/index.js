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
var isCompleteTree = function(root) {
  const arr = [ 0 ]

  root.index = 1
  let q = [ root ]
  while (q.length > 0) {
    let node = q.shift()
    const i = node.index
    arr.push(i)

    if (node.left) {
      node.left.index = i * 2
      q.push(node.left)
    }

    if (node.right) {
      node.right.index = i * 2 + 1
      q.push(node.right)
    }
  }

  for (let i = 0; i < arr.length; i++) {
    if (i !== arr[i]) {
      return false
    }
  }

  return true
}

function TreeNode(val) {
  this.val = val
  this.left = this.right = null
}

const root = new TreeNode(1)
root.left = new TreeNode(2)
//root.left.left = new TreeNode(6)
root.left.right = new TreeNode(30)

root.right = new TreeNode(3)

console.log(isCompleteTree(root))
