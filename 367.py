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

def mySqrt(x):
    if x == 1:
        return 1
    x0 = x // 2
    while True:
        xn = .5 * (x0 + x / x0)
        if abs(xn - x0) < 1e-5:
            break
        x0 = xn
    return xn

class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if num <= 0:    
            return False    
        eps = 10_000_000_000.   
        res = mySqrt(num)
        return round(trunc(res*eps)) == round(trunc(res)*eps)

if __name__ == '__main__':
    print(Solution().isPerfectSquare(16))
    print(Solution().isPerfectSquare(9))
    print(Solution().isPerfectSquare(27))
    print('ok')

