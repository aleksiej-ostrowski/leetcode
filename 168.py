
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

import string

class Solution:
    def __init__(self):
        self.STR_ASCII = [*string.ascii_uppercase]
        self.LEN_ASCII = len(self.STR_ASCII)
        # print(self.STR_ASCII)
    def convertToTitle(self, columnNumber: int) -> str:
        num = columnNumber
        res = []
        while True:
            # print("num = ", num)
            res.append(self.STR_ASCII[num % self.LEN_ASCII - 1])
            if num <= self.LEN_ASCII:
                break
            num = (num - 1) // self.LEN_ASCII
        return "".join(res[::-1])

if __name__ == '__main__':
    print(Solution().convertToTitle(701)) #52 = AZ,  701 = ZY
    print('ok')

