""" 561.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from functools import reduce
from typing import List

def chunks(lst, n):
    return [lst[i:i+n] for i in range(0, len(lst), n)]

class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        return sum(map(lambda x: min(x[0],x[1]), chunks(sorted(nums),2)))

if __name__ == '__main__':
    print(Solution().arrayPairSum([1,4,3,2]))
    print('ok')

