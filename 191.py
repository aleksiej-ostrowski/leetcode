
""" 191.py """

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
    def hammingWeight(self, n: int) -> int:
        # print(f"{n:032b}")
        return sum([(n >> bit) & 1 for bit in range(0, 32)])

if __name__ == '__main__':
    print(Solution().hammingWeight(0b010100101))
    print('ok')

