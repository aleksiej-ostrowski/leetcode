
""" 278.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

def isBadVersion(version: int) -> bool:
    return version >= 4

class Solution:
    def firstBadVersion(self, n: int) -> int:

        if n == 1:
            return 1

        low = 1
        high = n
     
        while True:
            mid = (low + high) >> 1
            if isBadVersion(mid):
                high = mid - 1
                if not isBadVersion(high):
                    return mid
            else:
                low = mid + 1
                if isBadVersion(low):
                    return low
            

if __name__ == '__main__':
    print(Solution().firstBadVersion(20))
    print('ok')

