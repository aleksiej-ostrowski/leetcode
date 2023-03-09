""" 409.py """

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
    def longestPalindrome(self, s: str) -> int:
        memory = defaultdict(lambda:0)
        for el in s:
            memory[el] += 1
        sum_even = 0   
        sum_1 = 0
        for el in memory:
            sum_even += memory[el] >> 1
            if sum_1 == 0 and memory[el] % 2 != 0:
                sum_1 += 1
        sum_even <<= 1
        if sum_1 > 0:
            sum_even += 1
        return sum_even 


if __name__ == '__main__':
    print(Solution().longestPalindrome("abccccdd"))
    print('ok')

