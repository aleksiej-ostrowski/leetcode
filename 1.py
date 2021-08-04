#-----------------------------------------#
#                                         #
#  version 0.0.1                          #
#  https://leetcode.com/problems/two-sum/ #
#                                         #
#  Aleksiej Ostrowski, 2020               #
#                                         #
#  https://aleksiej.com                   #
#                                         #
#-----------------------------------------#

from itertools import combinations

class Solution(object):

    def __init__(self, nums, target):
        self.nums = nums
        self.target = target

    def twoSum(self):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        new_nums = sorted(list(enumerate(self.nums)), key = lambda x: x[1])

        for x in combinations(range(len(new_nums)), 2):
            if new_nums[x[0]][1] + new_nums[x[1]][1] == self.target:
                return list((new_nums[x[0]][0], new_nums[x[1]][0]))

'''
nums = [0,4,3,0]
target = 0

nums = [2, 20, 100, 7, 11, 15]
target = 9
'''

nums = [-3,4,3,90]
target = 0

print(nums)
print(target)

sol = Solution(nums, target)
print(sol.twoSum())
