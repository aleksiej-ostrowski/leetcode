""" 521.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from collections import defaultdict

# this is not my solution
class Solution:
    def findLUSlength(self, a: str, b: str) -> int:
        if a == b:
            return -1
        return max(len(a),len(b))


if __name__ == '__main__':
    print(Solution().findLUSlength("aba", "cdc"))
    print('ok')

