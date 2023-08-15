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

from typing import Optional, List

# Definition for a binary tree node.

def checkSum2(lst: List, k: int) -> bool:
    for ind1, el1 in enumerate(lst):
        for _, el2 in enumerate(lst[ind1+1:]):
            if el1 + el2 == k:
                return True
    return False         

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def __init__(self):
        self.res = []
 
    def getVals(self, root: Optional[TreeNode]):
        if root != None:
            left = self.getVals(root.left)
            right = self.getVals(root.right)
            self.res.append(root.val)
            if left != None:
                self.res.append(left)
            if right != None:
                self.res.append(right)
        # return []

    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        self.getVals(root)
        return checkSum2(self.res, k)

if __name__ == '__main__':
    tree5 = TreeNode(7, None, None)
    tree4 = TreeNode(4, None, None)
    tree3 = TreeNode(2, None, None)
    tree2 = TreeNode(6, None, tree5)
    tree1 = TreeNode(3, tree3, tree4)
    tree = TreeNode(5, tree1, tree2)
    print(Solution().findTarget(tree, 9))
    print('ok')

