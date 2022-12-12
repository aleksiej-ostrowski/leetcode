
""" 118.py """

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
    def generate(self, numRows: int) -> List[List[int]]:
        res = [[1]]
        if numRows == 1:
            return res
        start = [1,1]
        res.append(start)
        if numRows == 2:
            return res
        for cic in range(numRows-2):
            start = [1] + [start[x] + start[x-1] for x in range(1, len(start))] + [1]
            res.append(start)
        return res

if __name__ == '__main__':
    print(Solution().generate(5))

