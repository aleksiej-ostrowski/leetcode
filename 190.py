
""" 190.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

# https://stackoverflow.com/questions/63759207/circular-shift-of-a-bit-in-python-equivalent-of-fortrans-ishftc
def ISHFTC(n, d, N):  
    return ((n << d) % (1 << N)) | (n >> (N - d))

class Solution:
    def reverseBits(self, n: int) -> int:
        return ISHFTC(n, 1, 32)

if __name__ == '__main__':
    print(Solution().reverseBits(34))
    print('ok')

