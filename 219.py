
""" 219.py """

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
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        memory = defaultdict(set)
        for ind1, el in enumerate(nums):
            for ind2 in memory[el]: 
                if abs(ind2 - ind1) <= k:
                    return True
            memory[el].add(ind1)        
        return False
                
if __name__ == '__main__':
    # print(Solution().containsNearbyDuplicate(nums=[1,0,1,1], k=1))
    print(Solution().containsNearbyDuplicate(nums=[1,2,3,1,2,3], k=2))
    print('ok')


