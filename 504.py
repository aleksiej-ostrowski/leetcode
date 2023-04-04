""" 504.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

class Solution:
    def convertToBase7(self, num: int) -> str:
        res = []
        zn = "-" if num < 0 else "" 
        num = abs(num)
        while True:
            if num < 7:
                break
            res.append(num % 7) 
            num //= 7
        res.append(num) 
        return zn + "".join(map(str, res[::-1]))

if __name__ == '__main__':
    # print(Solution().convertToBase7(100))
    # print(Solution().convertToBase7(77))
    print(Solution().convertToBase7(7))
    print('ok')

