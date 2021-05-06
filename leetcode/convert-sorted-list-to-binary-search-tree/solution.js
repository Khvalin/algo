/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {ListNode} head
 * @return {TreeNode}
 */
var sortedListToBST = function(head) {
    let vals = []
    
    let cur = head
    while (cur) {
        vals.push(cur.val)
        cur = cur.next
    }
    if (!vals.length) {
        return null
    }
    
    const convert = (left, right) => {
        if (left >= right) {
            return null
        }
        
        const mid = (left + right) >> 1
        return new TreeNode(vals[mid], convert(left, mid), convert(mid + 1, right))
    }
    
    return convert(0, vals.length)
};
