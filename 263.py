
""" 263.py """

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
    def isUgly(self, n: int) -> bool:    
        if n == 0:    
            return False
        if n == 1:    
            return True        
        if n % 2 == 0:    
            return self.isUgly(n>>1)        
        if n % 3 == 0:    
            return self.isUgly(n//3)        
        if n % 5 == 0:    
            return self.isUgly(n//5)        
        return False  

if __name__ == '__main__':
    # print(Solution().isUgly(6))
    print(Solution().isUgly(14))
    print('ok')

