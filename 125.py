
""" 125.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

import string 

class Solution:
    def __init__(self):
        self.ascii = set(string.ascii_lowercase + string.digits)
    def isPalindrome(self, s: str) -> bool:
        return (lambda x: x == x[::-1])(''.join(filter(lambda x: x in self.ascii, s.lower())))

if __name__ == '__main__':
    print(Solution().isPalindrome('A man, a plan, a canal: Panama'))





