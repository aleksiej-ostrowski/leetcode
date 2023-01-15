
""" 205.py """

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
    def isIsomorphic(self, s: str, t: str) -> bool:
        return self.isEqMap(s, t) and self.isEqMap(t, s)
    def isEqMap(self, s: str, t: str) -> bool:	
        research = {}
        for s_ind, s_elm in enumerate(s):
            if s_elm in research:
                if research[s_elm] != t[s_ind]:
                    return False
            else:    
                research[s_elm] = t[s_ind]
        return True        

if __name__ == '__main__':
    # print(Solution().isIsomorphic(s="egg", t="add"))
    print(Solution().isIsomorphic(s="foo", t="bar"))
    print('ok')


