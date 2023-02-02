
""" 228.py """

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
    def summaryRanges(self, nums: List[int]) -> List[str]:
        # [0,1,2,4,5,7]
        len_nums = len(nums)
        if len_nums == 0:
            return []
        el0 = nums[0]
        if len_nums == 1:
            return [str(el0)]
        start = 0
        end = 0
        res = []
        for idx, el in enumerate(nums[1:]):
            flag = False
            if (el != el0 + 1) or (idx + 1 == len_nums - 1):
                flag = True
                r = ""
                if start == end:
                    r = str(nums[start])
                else:
                    r = f"{nums[start]}->{nums[end]}"
                res.append(r)
                start = idx + 1
            if flag and (idx + 1 == len_nums - 1):
                res.append(str(el))    
            end = idx + 1
            el0 = el
        return res        


if __name__ == '__main__':
    print(Solution().summaryRanges(nums=[0,1,2,4,5,6]))
    print('ok')

