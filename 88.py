""" 88.py """

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
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        # print(nums1)
        # print(nums2)
        ind1 = m - 1
        ind2 = n - 1
        cur  = n + m - 1
        while True:
            flag1 = ind1 >= 0
            flag2 = ind2 >= 0 
            if flag1:
                el1 = nums1[ind1]
            if flag2:
                el2 = nums2[ind2]
            if flag1 and flag2:    
                if el1 < el2:
                    nums1[cur] = el2
                    ind2 -= 1
                else:    
                    nums1[cur] = el1
                    ind1 -= 1
            else:
                if flag1:
                    nums1[cur] = el1
                    ind1 -= 1
                else:
                    nums1[cur] = el2
                    ind2 -= 1     
            cur -= 1
            if cur < 0:
                break
        # print(nums1)    
  

if __name__ == "__main__":
    print(Solution().merge([1, 30, 40, 50, 0, 0, 0], 4, [7, 8, 10], 3))
    # print(Solution().merge([0], 0, [1], 1))

    
