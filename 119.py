
""" 119.py """

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
    def getRow(self, rowIndex: int) -> List[int]:
        if rowIndex == 0:
            return [1]
        start = [1,1]
        if rowIndex == 1:
            return start
        for cic in range(rowIndex-1):
            start = [1] + [start[x] + start[x-1] for x in range(1, len(start))] + [1]
        return start

if __name__ == '__main__':
    print(Solution().getRow(4))

