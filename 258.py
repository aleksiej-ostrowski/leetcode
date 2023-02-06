
""" 258.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

# It's not my solution
# https://leetcode.com/problems/add-digits/solutions/666062/add-digits/

class Solution:
    def addDigits(self, num: int) -> int:
        return 1 + (num - 1) % 9 if num else 0

if __name__ == '__main__':
    print(Solution().addDigits(38))

