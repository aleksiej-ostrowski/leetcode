""" 28.py """

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
    def strStr(self, haystack: str, needle: str) -> int:
        idx1 = 0
        idx2 = 0
        cnt = 0
        start = 0
        len_needle = len(needle)
        len_haystack = len(haystack)
        while True:
            if idx2 >= len_haystack:
                break
            if cnt == 0:
                start = idx2
            flag = haystack[idx2] == needle[idx1]
            if flag:
                cnt += 1
                idx1 += 1
                idx2 += 1
            if cnt == len_needle:
                return start
            if not flag or idx1 >= len_needle:
                cnt = 0
                idx2 = start + 1
                idx1 = 0
        return -1  


if __name__ == '__main__':
    print(Solution().strStr("leetcode", "leeto"))
    print('ok')

