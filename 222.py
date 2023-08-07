""" main.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from typing import Optional

# Definition for a binary tree node.
class TreeNode:
     def __init__(self, val=0, left=None, right=None):
         self.val = val
         self.left = left
         self.right = right

class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        if root == None:
            return 0
        return 1 + self.countNodes(root.left) + self.countNodes(root.right) 


if __name__ == '__main__':
    tree6 = TreeNode(6, None, None)
    tree5 = TreeNode(5, None, None)
    tree4 = TreeNode(4, None, None)
    tree3 = TreeNode(3, tree6, None)
    tree2 = TreeNode(2, tree4, tree5)
    tree = TreeNode(1, tree2, tree3)
    print(Solution().countNodes(root=tree))
    print('ok')

