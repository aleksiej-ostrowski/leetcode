""" 405.py """

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
    def toHex(self, num: int) -> str:    
        allcodes = '0123456789abcdef'    
        res = ""    
        if num < 0:    
            num += 2**32    
        while True:    
            if num < 16:    
                res += allcodes[num]    
                return res[::-1]    
            else:        
                res += allcodes[num % 16]    
                num >>= 4  

if __name__ == '__main__':
    print(Solution().toHex(-2))
    print('ok')

