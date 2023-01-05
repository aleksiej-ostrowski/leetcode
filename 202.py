
""" 202.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

import math 

class Solution:
    def isHappy(self, n: int) -> bool:
        digits = n
        sm_q = .0
        while True:
            num = digits % 10
            print(num)
            sm_q += num * num
            if digits < 10:
                break
            digits = math.trunc(digits / 10.0)
        return True    

if __name__ == '__main__':
    print(Solution().isHappy(19))
    print('ok')

