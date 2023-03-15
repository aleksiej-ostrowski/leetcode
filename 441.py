""" 441.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

# https://oeis.org/A000217

import math

class Solution:
    def arrangeCoins(self, n: int) -> int:
        return math.trunc((-1 + (1 + 8 * n) ** 0.5) * 0.5)

if __name__ == '__main__':
    print(Solution().arrangeCoins(5))
    print('ok')

