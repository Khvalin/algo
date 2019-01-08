
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
  const dfs = (node, parent) => {
    if (!node) {
      return
    }

    node.val = parent? parent.val + 1:0

    let {left, right} = node
    
    left && dfs(left, node)
    right && dfs(right, node)

    if (!left && !right && !node.hasCam && (!parent || !parent.hasCam)) {
      (parent || node).hasCam = true
    } else {
      node.hasCam = node.hasCam || (left || right) && !(left && left.hasCam || right && right.hasCam)
    }

    res += node.hasCam? 1: 0
  }

  dfs(root)
  //console.log(root)

  return res
}

//console.log(minCameraCover(createTree([ 0, 0, null, 0, 0 ])))
//console.log(minCameraCover(createTree([ 0 ])))
console.log(minCameraCover(createTree([0,0,0,null,null,null,0])))
console.log(minCameraCover(createTree([0,0,null,null,0,0,null,null,0,0])))
