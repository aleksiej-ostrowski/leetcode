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

from typing import List
from functools import reduce

# https://leetcode.com/problems/maximum-product-of-three-numbers/solutions/3579502/two-java-solution-easy-explanation/

class Solution:
    def maximumProduct(self, nums: List[int]) -> int:
        nums.sort()
        return max(nums[0]*nums[1]*nums[-1], nums[-1]*nums[-2]*nums[-3])

if __name__ == '__main__':
    print(Solution().maximumProduct([1,2,3]))
    print('ok')

