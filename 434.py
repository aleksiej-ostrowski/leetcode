""" 434.py """

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
    def countSegments(self, s: str) -> int: 
        return len(list(filter(lambda el: el != "", s.split(" "))))

if __name__ == '__main__':
    # print(Solution().countSegments(""))
    print(Solution().countSegments(", , , ,        a, eaefa"))
    print('ok')



