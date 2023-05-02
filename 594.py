""" 594.py """

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
from collections import defaultdict


class Solution:
    def findLHS(self, nums: List[int]) -> int:
        dict_nums = defaultdict(lambda: 0)
        for el in nums:
            dict_nums[el] += 1
        sort_nums = sorted(list(set(nums)))
        # print(dict_nums)
        # print(sort_nums)
        idx = 0
        sum_glb = 0
        while True:
            if idx + 1 > len(sort_nums) - 1:
                break
            if abs(sort_nums[idx] - sort_nums[idx + 1]) == 1:
                if (
                    sum_cur := dict_nums[sort_nums[idx]] + dict_nums[sort_nums[idx + 1]]
                ) > sum_glb:
                    sum_glb = sum_cur
            idx += 1
        return sum_glb


if __name__ == "__main__":
    print(Solution().findLHS(nums=[1, 3, 2, 2, 5, 2, 3, 7]))
    print("ok")
