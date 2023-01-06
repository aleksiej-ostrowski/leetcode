
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
        memory = {}
        while True:
            num = digits % 10
            # print(num)
            sm_q += num * num
            # print("sm_q = ", sm_q)
            if digits < 10:
                if sm_q == 0:
                    return False
                if sm_q == 1:
                    return True
                if sm_q in memory:
                    return False
                memory[sm_q] = True
                digits = sm_q
                sm_q = 0
            else:    
                digits = math.trunc(digits / 10.0)
        return False   

if __name__ == '__main__':
    print(Solution().isHappy(19))
    print('ok')

