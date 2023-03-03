""" 387.py """

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
    def firstUniqChar(self, s: str) -> int:
        memory1 = defaultdict(lambda:0)
        memory2 = defaultdict(lambda:0)
        for idx, el in enumerate(s):
            memory1[el] += 1
            if el not in memory2:
                memory2[el] = idx 
        for key, value in memory1.items():   
            if value == 1:
                return memory2[key]
        return -1    

if __name__ == '__main__':
    print(Solution().firstUniqChar("loveleetcode"))
    print('ok')



