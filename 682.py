""" 682.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from typing import List

def tryConvToInt(st: str):
    res = st
    try:
        res = int(st)
    except:
        pass
    return res

class Solution:
    def calPoints(self, operations: List[str]) -> int:
        res = []
        for el in map(tryConvToInt, operations):
            # print("res = ", res, "el = ", el)
            match el:
                case int():
                    res.append(el)
                case "C":
                    res.pop()
                case "D":
                    res.append(res[-1]<<1)
                case "+":
                    res.append(res[-1]+res[-2])
        return sum(res)


if __name__ == '__main__':
    print(Solution().calPoints(operations=["5","-2","4","C","D","9","+","+"]))
