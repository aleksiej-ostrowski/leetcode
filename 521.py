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

class Solution:
    def findLUSlength(self, a: str, b: str) -> int:
        dict_a = defaultdict(lambda: '')
        dict_b = defaultdict(lambda: '')
        for idx, el_a in enumerate(a):
            dict_a[el_a] = idx
        for idx, el_b in enumerate(b):
            dict_a[el_b] = idx





if __name__ == '__main__':
    print(Solution().findLUSlength("aba", "cdc"))
    print('ok')

