/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {string}
 */
var smallestFromLeaf = function(root) {
  const ABC = 'abcdefghijklmnopqrstuvwxyz'

  let strings = []
  let minLen = 1000
  const walk = (node, path = '') => {
    if (node) {
      path = ABC[node.val] + path
      node.left && walk(node.left, path)
      node.right && walk(node.right, path)

      if (!node.left && !node.right) {
        strings.push(path)
        minLen = Math.min(minLen, path.length)
      }
    }
  }

  walk(root)
  strings = strings.filter((s) => s.length === minLen)

  strings.sort((s1, s2) => {
    return s1.localeCompare(s2)
  })

  return strings[0]
}
