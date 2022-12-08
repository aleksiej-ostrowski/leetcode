
from typing import List, Optional

# Definition for a binary tree node.

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/solutions/2404906/easy-solution-beginner-friendly-best-method/

class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        return self.helper(nums=nums, left=0, right=len(nums)-1)
    
    def helper(self, nums: List[int], left: int, right: int) -> Optional[TreeNode]:
        if left > right:
            return None
        mid = (left + right) >> 1
        temp = TreeNode(val = nums[mid])
        temp.left = self.helper(nums=nums, left=left, right=mid-1)
        temp.right = self.helper(nums=nums, left=mid+1, right=right)
        return temp

    def printS(self, t: Optional[TreeNode]):
        if t != None:
            print("(", t.val, end=' ')
        if t.left != None:
            self.printS(t.left)
        else:
            print("null", end=' ')
        if t.right != None:
            self.printS(t.right)
        else:
            print("null", end=' ')
        print(")", end=' ')    


if __name__ == '__main__':
    nums = [-10,-3,0,5,9]
    s = Solution()
    t = s.sortedArrayToBST(nums)
    s.printS(t)
    print('ok')
