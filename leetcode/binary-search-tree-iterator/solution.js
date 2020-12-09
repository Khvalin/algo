/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

const unfold = (node) => {
  const res = [];

  let cur = node;
  while (cur) {
    if (cur.right) {
      res.push([cur.right, true]);
    }
    res.push([cur, false]);
    cur = cur.left;
  }

  return res;
};

/**
 * @param {TreeNode} root
 */
var BSTIterator = function (root) {
  this.q = [[root, true]];
};

/**
 * @return {number}
 */
BSTIterator.prototype.next = function () {
  if (!this.q.length) {
    return null;
  }

  const next = this.q.pop();

  if (!next[1]) {
    return next[0].val;
  }

  this.q.push(...unfold(next[0]));

  return this.next();
};

/**
 * @return {boolean}
 */
BSTIterator.prototype.hasNext = function () {
  return this.q.length > 0;
};

/**
 * Your BSTIterator object will be instantiated and called as such:
 * var obj = new BSTIterator(root)
 * var param_1 = obj.next()
 * var param_2 = obj.hasNext()
 */
