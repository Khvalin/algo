/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {number[]}
 */
var getAllElements = function(root1, root2) {
  function* inOrderTraversal(root) {
    if (!root) {
      return;
    }

    function* walk(node) {
      if (node.left) {
        yield* walk(node.left);
      }

      yield node.val;

      if (node.right) {
        yield* walk(node.right);
      }
    }

    for (const i of walk(root)) {
      yield i;
    }
  }

  const res = [];
  const walker1 = inOrderTraversal(root1);
  const walker2 = inOrderTraversal(root2);

  let [ val1, val2 ] = [ null, null ];

  while (true) {
    val1 = val1 || walker1.next();
    val2 = val2 || walker2.next();

    if (val1.done && val2.done) {
      break;
    }

    let n = 0;
    if (val2.done || val1.value < val2.value) {
      n = val1.value;
      val1 = null;
    } else {
      n = val2.value;
      val2 = null;
    }

    res.push(n);
  }

  return res;
};
