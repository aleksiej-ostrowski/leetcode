""" 459.py """

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
    def repeatedSubstringPattern(self, s: str) -> bool:
        len_s = len(s)
        if len_s == 1:
            return False
        idx = 0
        new_s = ""
        len_new_s = 0
        cnt = 0
        while True:
            if len_new_s > 0:
                # if s[idx] == new_s[0] 
                if s.find(new_s, idx, idx + len_new_s) == idx:
                    idx += len_new_s
                    cnt += len_new_s
                else:
                    idx = len_new_s
                    new_s += s[idx]
                    len_new_s = len(new_s)
                    idx += 1
                    cnt = 0
            else:        
                new_s += s[idx]
                len_new_s = len(new_s)
                idx += 1
            if idx >= len_s:
                break
            if len_new_s > len_s >> 1:
                break
        return (cnt > 0) and (cnt + len_new_s == len_s)

if __name__ == '__main__':
    # print(Solution().repeatedSubstringPattern("abcapcabcabc"))
    print(Solution().repeatedSubstringPattern("ab"))
    print('ok')

