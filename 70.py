""" 70.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

class Solution:
    '''

    def __init__(self):
        self.sum_comb = 0

    # https://www.cyberforum.ru/python-beginners/thread2748774.html
    def count_comb(self, n, k=[]):
        if n != 0:
            for i in range(n, 0, -1):
                if i <= 2: 
                    self.count_comb(n-i, k+[i])
        else:
            print(*k, sep=' + ')
            self.sum_comb += 1

    def climbStairs(self, n: int) -> int:
        self.count_comb(n)
        return self.sum_comb

    '''

    # https://leetcode.com/problems/climbing-stairs/solutions/2765140/python-dp/
    def climbStairs(self, n: int) -> int:

        '''
        dp = [0, 1]
        
        for i in range(2, n+2):
            dp.append(dp[i-2] + dp[i-1])
            
        print(dp)

        return dp[n+1]    

        '''

        d1 = 0
        d2 = 1
        d3 = 1
       
        for i in range(2, n+2):
            d3 = d1 + d2
            # print(d3)
            d1 = d2
            d2 = d3

        return d3

        
if __name__ == '__main__':
    print(Solution().climbStairs(35)) # 14930352

