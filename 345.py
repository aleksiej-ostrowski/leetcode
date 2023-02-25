""" 345.py """

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

class Solution:
    def reverseVowels(self, s: str) -> str:

        s = list(s)

        print(s)

        vowels = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']    

        border_begin = 0 
        border_end = len(s) - 1

        while True:

            if border_begin >= border_end:
                break

            while True:
                if s[border_begin] not in vowels:
                    border_begin += 1
                    if border_begin >= border_end:
                        break
                else:
                    break

            while True:
                if s[border_end] not in vowels:
                    border_end -= 1
                    if border_begin >= border_end:
                        break
                else:
                    break

            if border_begin < border_end:
                s[border_begin], s[border_end] = s[border_end], s[border_begin]
            else:
                break

            border_begin += 1
            border_end -= 1
        print(s)
        print("="*30)
        return "".join(s)


if __name__ == '__main__':
    Solution().reverseVowels(s=["h","e","3","4","o"])
    Solution().reverseVowels(s=["h","e","3","4","o", "2"])
    print('ok')

