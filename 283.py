
""" 283.py """

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
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        print(nums)
        cnt_zero = 1
        len_nums = len(nums)
        idx = 0
        while True:
            if nums[0] == 0:
                cnt_zero += 1
            else:
                nums.append(nums[0])
            del nums[0]
            idx += 1
            if idx >= len_nums:
                break
        if cnt_zero > 1:
            nums += [0] * (cnt_zero - 1)
        print(nums)        

if __name__ == '__main__':
    print(Solution().moveZeroes(nums=[0,1,0,3,12]))
    # print(Solution().moveZeroes(nums=[1,2]))
    # print(Solution().moveZeroes(nums=[1]))
    print('ok')

