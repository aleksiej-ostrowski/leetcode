""" 338.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from typing import List

# https://oeis.org/A000120/a000120.py.txt
def a000120_list(n):
    """Returns a list of the first n >= 0 terms."""
    if n < 0: 
        raise ValueError
    a = [0] * n
    if n < 2:
        return a
    a[1] = t = 1
    while t < n:
        for m in range(t, t+t):
            if m + m < n: 
                a[m + m] = a[m]
            else: 
                break
            if m + m + 1 < n: 
                a[m + m + 1] = a[m] + 1
            else: 
                break
        t += t
    return a


def a000120_list2(n):
    return [el.bit_count() for el in range(n)]

class Solution:
    def countBits(self, n: int) -> List[int]:
        # return a000120_list(n + 1)
        return a000120_list2(n + 1)

if __name__ == '__main__':
    print(Solution().countBits(5))
    print('ok')

