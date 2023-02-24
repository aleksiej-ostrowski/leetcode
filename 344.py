""" 344.py """

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
    def reverseString(self, s: List[str]) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        print(s)

        # s = s[::-1]

        len_s_d2 = (len(s) >> 1) - 1

        if len(s) % 2 == 0:
            len_s_d2 -= 1
            
        for idx, x in enumerate(s):
            s[idx], s[-(idx+1)] = s[-(idx+1)], s[idx]
            if idx > len_s_d2:
                break
        print(s)
        print("="*30)


if __name__ == '__main__':
    Solution().reverseString(s=["h","e","3","4","o"])
    Solution().reverseString(s=["h","e","3","4","o", "2"])
    print('ok')

