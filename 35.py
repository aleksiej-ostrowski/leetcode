""" 35.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from typing import List

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        nums = [(i, el) for i, el in enumerate(nums)]
        while True:
            # print(nums)
            len_nums = len(nums)
            mid_nums = len_nums >> 1
            mid = nums[mid_nums][1]
            if mid == target:
                return nums[mid_nums][0]
            elif target < mid:
                if len_nums == 1:
                    return nums[mid_nums][0]
                nums = nums[:mid_nums]
            else:
                if len_nums == 1:
                    return nums[mid_nums][0] + 1 
                nums = nums[mid_nums:]

if __name__ == '__main__':
    # Input: nums = [1,3,5,6], target = 5
    # Output: 2

    nums = [1,3,5,7]
    target = 6

    res = Solution().searchInsert(nums, target)
    print(res)

    print('ok')

