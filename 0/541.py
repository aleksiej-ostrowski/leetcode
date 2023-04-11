""" 541.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        # print(s)
        res = []
        for idx in range(0, len(s), k):
            ad = s[idx:idx+k]
            if idx % (k * 2) == 0:
                ad = ad[::-1]
            res += ad
        # print(res)
        return ''.join(res)        
     
if __name__ == '__main__':
    # print(Solution().reverseStr(s="abcdefg", k=2))
    # print(Solution().reverseStr(s="abcd", k=2))
    print(Solution().reverseStr(s="abcdefg", k=4))
    print('ok')

