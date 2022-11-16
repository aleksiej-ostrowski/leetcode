""" 66.py """

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
    def plusOne(self, digits: List[int]) -> List[int]:
        overflow = 0
        i = len(digits) - 1
        first = True
        while i >= 0:    
            new_x = digits[i] + overflow
            overflow = 0
            if first:
                new_x += 1
                first = False
            if new_x <= 9:
                digits[i] = new_x
                if overflow == 0:
                    return digits
            else:
                digits[i] = new_x - 10
                overflow += 1
            i -= 1    
        if overflow > 0:
            digits.insert(0, overflow)
        return digits    

if __name__ == '__main__':
    digits = [1,2,3,4]
    print(digits)
    print(Solution().plusOne(digits))

