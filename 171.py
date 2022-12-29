
""" main.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

# copy from 168.py

import string
import math

class Solution:
    def __init__(self):
        self.STR_ASCII = [*string.ascii_uppercase]
        self.DICT_ASCII = dict( (e,k+1) for k, e in enumerate(self.STR_ASCII))
    def titleToNumber(self, columnTitle: str) -> int:    
        res = [self.DICT_ASCII[el] for el in columnTitle]
        return round(sum([el * math.pow(26, ind) for ind, el in enumerate(res[::-1])]))

if __name__ == '__main__':
    print(Solution().titleToNumber("ZY")) #52 = AZ,  701 = ZY
    # print(Solution().DICT_ASCII)
    print('ok')

