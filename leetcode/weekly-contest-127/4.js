/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} preorder
 * @return {TreeNode}
 */
var bstFromPreorder = function(preorder) {
  function TreeNode(val) {
    this.val = val
    this.left = this.right = null
  }

  let insert = (node, n) => {
    if (node) {
      if (node.val > n) {
        node.left = insert(node.left, n)
      } else {
        node.right = insert(node.right, n)
      }
    } else {
      node = new TreeNode(n)
    }

    return node
  }

  const root = new TreeNode(preorder.shift())
  for (const i of preorder) {
    insert(root, i)
  }

  return root
}

console.log(bstFromPreorder([ 8, 5, 1, 7, 10, 12 ]))
