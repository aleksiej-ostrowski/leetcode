
""" 217.py """

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
    """
    def containsDuplicate2(self, nums: List[int]) -> bool:
        return sorted(list(set(nums))) != sorted(nums)
    """
    def containsDuplicate(self, nums: List[int]) -> bool:
        memory = {}
        for el in nums:
            if el in memory:
                return True
            else:
                memory[el] = True
        return False
        
if __name__ == '__main__':
    print(Solution().containsDuplicate(nums=[1,1,1,3,3,4,3,2,4,2]))
    print('ok')


