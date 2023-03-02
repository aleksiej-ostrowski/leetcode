""" 383.py """

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
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        memory = defaultdict(lambda:0)
        for val in magazine:
            memory[val] += 1
        for val in ransomNote:
            if val in memory:
                memory[val] -= 1
                if memory[val] == 0:
                    del memory[val]
            else:
                return False
        return True    
     
if __name__ == '__main__':
    print(Solution().canConstruct("aa", "aab"))
    print(Solution().canConstruct("aa", "ab"))
    print('ok')

