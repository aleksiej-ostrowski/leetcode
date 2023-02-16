""" 342.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from math import log2, trunc

class Solution:    
    def isPowerOfFour(self, n: int) -> bool:    
        if n <= 0:    
            return False    
        eps = 10_000_000_000.    
        num = log2(n) / log2(4)    
        return round(trunc(num*eps)) == round(trunc(num)*eps)

if __name__ == '__main__':
    print(Solution().isPowerOfFour(16))
    # print(Solution().isPowerOfThree(-1))
    print('ok')

