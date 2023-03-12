""" 415.py """

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
    def addStrings(self, num1: str, num2: str) -> str:
        len_num1 = len(num1)
        len_num2 = len(num2)
        idx1 = len_num1 - 1
        idx2 = len_num2 - 1
        memory = 0
        res = ""
        while True:
            d1 = 0
            d2 = 0
            if idx1 >= 0:
                d1 = int(num1[idx1])
            if idx2 >= 0:            
                d2 = int(num2[idx2])
            if idx1 < 0 and idx2 < 0:
                break
            sum_d1_d2 = d1 + d2 + memory
            if sum_d1_d2 <= 9:
                res += str(sum_d1_d2)
                memory = 0
            else:
                s_d = sum_d1_d2 % 10
                memory = sum_d1_d2 // 10
                res += str(s_d)
            idx1 -= 1
            idx2 -= 1
        if memory > 0: 
           res += str(memory)      
        return res[::-1]     


if __name__ == '__main__':
    print(Solution().addStrings("456", "77"))
    print('ok')

