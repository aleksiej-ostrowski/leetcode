""" 495.py """

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
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        memory = [[timeSeries[0], timeSeries[0] + duration - 1]]
        for a2 in timeSeries[1:]:
            b2 = a2 + duration - 1
            a1, b1 = memory[-1]
            if a2 > b1:
                memory.append([a2,b2])
            if (b1 >= a2) and (b2 > b1):
                memory[-1][1] = b2     
        sum_e = 0
        for a, b in memory:
            sum_e += (b - a) + 1

        return sum_e    


if __name__ == '__main__':
    # print(Solution().findPoisonedDuration(timeSeries=[1,2], duration=2))
    # print(Solution().findPoisonedDuration(timeSeries=[1,4], duration=2))
    print(Solution().findPoisonedDuration(timeSeries=[1,2,3,4,5,6,7,8,9], duration=10))
    print('ok')

