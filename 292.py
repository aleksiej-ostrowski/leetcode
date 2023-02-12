""" 292.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

# https://leetcode.com/problems/nim-game/solutions/2773531/easy-java-solution-beats-100-0ms/
class Solution:
    def canWinNim(self, n: int) -> bool:
        return n % 4 != 0

if __name__ == '__main__':
    print(Solution().canWinNim(4))
    print('ok')
