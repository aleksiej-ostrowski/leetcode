
""" 268.py """

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

'''
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        len_nums = len(nums) + 1
        return ((len_nums * (len_nums - 1)) >> 1) - sum(nums)
'''

class Solution:    
    def missingNumber(self, nums: List[int]) -> int:    
        return (( ((lenn:=len(nums)) + 1 ) * lenn) >> 1) - sum(nums) 

if __name__ == '__main__':
    # print(Solution().missingNumber([9,6,4,2,3,5,7,0,1]))
    print(Solution().missingNumber([0,1]))
    print('ok')
   
