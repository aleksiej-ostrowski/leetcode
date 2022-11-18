""" 69.py """

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
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        a = x >> 1
        if a == 0:
            return 1
        a_old = x
        an = 0
        while True:
            an = (a + x // a) >> 1
            # an = round((a + x / a) * .5)
            if ((a == an) or (a_old == an)) and (an * an <= x):
                break
            a_old = a
            a = an
        return an

if __name__ == '__main__':
    print(Solution().mySqrt(4))





