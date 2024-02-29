# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isEvenOddTree(self, root: Optional[TreeNode]) -> bool:
        nodes = [root]

        i = 0
        while len(nodes):
            a = []
            prev = None
            for node in nodes:
                val = node.val
                if i % 2 == 0:
                    if val % 2 == 0 or ((prev is not None) and prev >= val):
                        return False

                if i % 2 == 1:
                    if val % 2 == 1 or ((prev is not None) and prev <= val):
                        return False

                prev = val
                if node.left:
                    a.append(node.left)
                if node.right:
                    a.append(node.right)
            i += 1
            nodes = a

        return True


Solution().isEvenOddTree({val: 1})
