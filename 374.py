""" 374.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:

def guess(num: int) -> int:
    NUM = 1702766719 #50
    if num == NUM:
        return 0
    elif num < NUM:
        return 1
    else:
        return -1

class Solution:
    def guessNumber(self, n: int) -> int:
        if n == 1:
            return 1
        border_start = 1
        border_end = n
        n0 = (1 + n) >> 1
        memory = {}
        cnt = 0
        
        while True:
            cnt += 1
            match guess(n0):
                case 0:
                    print(f"cnt={cnt}\n")    
                    return n0
                case 1:
                    border_start = n0 
                    if n0 in memory:
                        n0 += 1
                    else:
                        n0 = (border_end + n0) >> 1
                        memory[n0] = True
                case -1:
                    border_end = n0 
                    if n0 in memory:
                        n0 -= 1
                    else:
                        n0 = (border_start + n0) >> 1
                        memory[n0] = True 
            print(n0, "\n")       
 
if __name__ == '__main__':
    # print(Solution().guessNumber(1000))
    print(Solution().guessNumber(2126753390))
