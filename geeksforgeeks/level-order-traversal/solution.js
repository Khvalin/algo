
/**
 * @param {Node} node
 * @returns {number[]}
*/

/*
class Node{
    constructor(data){
        this.data = data;
        this.left = null;
        this.right = null;
    }
}
*/

class Solution {
    //Function to return the level order traversal of a tree.
    levelOrder(node)
    {
        const q = [node]
        const res = []
        let i = 0
        while (i < q.length) {
            const node = q[i++]
            if (!node) {
                continue
            }

            res.push(node.data)

            if (node.left) {
                q.push(node.left)
            }

            if (node.right) {
                q.push(node.right)
            }
        }

        return res
    }
}
