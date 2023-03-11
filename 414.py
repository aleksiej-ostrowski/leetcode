""" 414.py """

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

class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        nums = list(set(nums))
        match len(nums):
            case 1:
                return nums[0]
            case 2:
                return max(nums)
            case 3:
                return min(nums)
        a, b, c = nums[0], nums[1], nums[2]  
        min_abc = min(a, b, c)
        max_abc = max(a, b, c)
        md = a + b + c - min_abc - max_abc       
        a, b, c = max_abc, md, min_abc
        for el in nums[3:]:
            if el > a:
                c = b
                b = a
                a = el
            elif (el > b) and (el < a):
                c = b 
                b = el
            elif (el > c) and (el < b):
                c = el
        return c  

if __name__ == '__main__':
    print(Solution().thirdMax(nums=[-4,-5,-3,-2,-1]))
    print('ok')

