
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

class Solution:
    def reverseBits(self, n: int) -> int:
        res = 0
        # print(f"{n:032b}")
        for ind, x in enumerate([(n >> bit) & 1 for bit in range(0, 32)][::-1]):
            if x == 1:
                res |= 1 << ind
        return res            

if __name__ == '__main__':
    print(Solution().reverseBits(0b00000010100101000001111010011100))
    print('ok')

