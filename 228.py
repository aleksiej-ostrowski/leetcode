
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

    def chunks2(self, lst):
        # print(lst)
        return [i for i in range(0, len(lst)-1, 1) if (el:=lst[i:i+2])[1] != el[0] +1]

    def mysplit(self, lst, spl):
        res = []
        minus = 0
        for idx in spl:
            idx -= minus
            res.append(lst[:idx+1])
            minus += len(lst[:idx+1])
            lst = lst[idx+1:]
        res.append(lst)
        return res

    def summaryRanges(self, nums: List[int]) -> List[str]:
        len_nums = len(nums)
        if len_nums == 0:
            return []
        if len_nums == 1:
            return [f"{nums[0]}"]
        # print(nums)
        res = self.mysplit(lst=nums, spl=self.chunks2(lst=nums))
        res = [f"{el[0]}" if len(el)==1 else f"{el[0]}->{el[-1]}" for el in res]
        return res

if __name__ == '__main__':
    print(Solution().summaryRanges(nums=[0,1,2,4,5,7]))
    print('ok')

