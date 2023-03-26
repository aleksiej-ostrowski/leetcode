""" 485.py """

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
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        sum_cur = 0
        max_sum = 0
        for el in nums:
            if el == 1:
                sum_cur += 1
                if sum_cur > max_sum:
                    max_sum = sum_cur
            else:
                sum_cur = 0
        return max_sum        

if __name__ == '__main__':
    print(Solution().findMaxConsecutiveOnes(nums=[1,0,1,1,0,1]))
    print('ok')

