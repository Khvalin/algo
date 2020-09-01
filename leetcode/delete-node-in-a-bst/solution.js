/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} key
 * @return {TreeNode}
 */
var deleteNode = function(root, key) {
  const findNode = (node, key) => {
    if (!node || node.val === key) {
      return [ null, '', node ];
    }

    if (node.val > key) {
      if (node.left && node.left.val === key) {
        return [ node, 'left', node.left ];
      }
      return findNode(node.left, key);
    }

    if (node.right && node.right.val === key) {
      return [ node, 'right', node.right ];
    }

    return findNode(node.right, key);
  };

  const [ parent, dir, target ] = findNode(root, key);

  if (!target) {
    return root;
  }

  if (!target.left || !target.right) {
    const next = target.left || target.right || null;
    if (root === target) {
      root = next;
    } else {
      parent[dir] = next;
    }

    return root;
  }

  const grandChild = target.left.right;
  target.left.right = target.right;
  if (parent) {
    parent[dir] = target.left;
  } else {
    root = target.left;
  }

  let newParent = target.right;
  while (newParent.left) {
    newParent = newParent.left;
  }
  newParent.left = grandChild;

  return root;
};
