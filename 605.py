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
import math

class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        nn_0, nn_1 = 0, 0
        next = -1  
        new_flo = []
        for ind, el in enumerate(flowerbed):
            if ind + 1 < len(flowerbed):
                next = flowerbed[ind+1]
            else:
                next = -1    
            if el == 0:
                nn_0 += 1
                if next != 0:
                    new_flo.append(nn_0)
                    nn_0 = 0
            else:
                nn_1 += 1
                if next != 1:
                    new_flo.append(-1)
                    nn_1 = 0
        # print(new_flo)
        for ind, el in enumerate(new_flo):
            if el != -1:
                sm = math.ceil(el / 2)
                pp = 0
                if ind - 1 >= 0:
                    pp += 1
                if ind + 1 < len(new_flo):
                    pp += 1
                if pp == 0:
                    n -= sm
                if pp == 1: 
                    if el in [1,3]:
                        n -= sm - 1
                    else:
                        n -= sm
                if pp == 2:
                    n -= sm - 1
            if n <= 0:
                return True
        return False                            


if __name__ == '__main__':
    flowerbed = [1,0,0,0,1,0,0] 
    n = 2
    print(Solution().canPlaceFlowers(flowerbed=flowerbed,n=n))
    print('ok')