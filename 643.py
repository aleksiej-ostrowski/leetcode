""" main.py """

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
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        maxf = sum(nums[:k])
        newsum = maxf
        for idx, _ in enumerate(nums[1:len(nums)-k+1]):
            newsum = newsum - nums[idx] + nums[idx+k]
            if newsum > maxf:
                maxf = newsum
        return maxf / k

if __name__ == '__main__':
    # print(Solution().findMaxAverage(nums=[1,12,-5,-6,50,3], k=4))
    print(Solution().findMaxAverage(nums=[0,4,0,3,2], k=1))
    print('ok')

