/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 */
var FindElements = function(root) {
  this.vals = {}

  const nodes = [ [ root, 0 ] ]
  while (nodes.length > 0) {
    const [ node, val ] = nodes.shift()

    if (node) {
      node.val = val
      this.vals[val] = true

      nodes.push([ node.left, 2 * val + 1 ])
      nodes.push([ node.right, 2 * val + 2 ])
    }
  }
}

/** 
 * @param {number} target
 * @return {boolean}
 */
FindElements.prototype.find = function(target) {
  return !!this.vals[target]
}

/** 
 * Your FindElements object will be instantiated and called as such:
 * var obj = new FindElements(root)
 * var param_1 = obj.find(target)
 */
