
""" 674.py """

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
    def findLengthOfLCIS(self, nums: List[int]) -> int:    
        gl_sm = -1  
        ind = 0  
        while True:    
            if ind >= len(nums):
                break
            sm = 1    
            last = nums[ind]    
            for ind2 in range(ind+1, len(nums)):    
                el2 = nums[ind2]
                if last < el2:    
                    sm += 1    
                    last = el2    
                else: 
                    ind = ind2 - 1   
                    break    
            if sm > 0 and sm > gl_sm:    
                gl_sm = sm    
            ind += 1    
        return gl_sm  

if __name__ == '__main__':
    # print(Solution().findLengthOfLCIS([1,3,5,4,7]))
    # print(Solution().findLengthOfLCIS([2,2,2,2,2]))
    print(Solution().findLengthOfLCIS([1,3,5,7]))
    print('ok')

