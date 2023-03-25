""" 482.py """

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
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        plain = "".join(s.upper().split("-"))[::-1]
        new_s = "-".join([plain[ind:ind+k] for ind in range(0, len(plain), k)])
        return new_s[::-1]

if __name__ == '__main__':
    # print(Solution().licenseKeyFormatting(s="5F3Z-2e-9-w", k=4))
    print(Solution().licenseKeyFormatting(s="2-4A0r7-4k", k=4))
    print('ok')

