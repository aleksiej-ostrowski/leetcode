""" _551.py """

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
    def checkRecord(self, s: str) -> bool:
        pred_L = -1
        A = 0
        L = 0
        all_L = []
        for idx, el in enumerate(s):
            match el:
                case "A":
                    A += 1
                case "L":
                    if (pred_L + 1 == idx) or (pred_L == -1):
                        L += 1
                    else:
                        all_L.append(L)
                        L = 1
                    pred_L = idx
        if L != 0:
            all_L.append(L)
        return A < 2 and all(el < 3 for el in all_L)


if __name__ == "__main__":
    # print(Solution().checkRecord(s="PPALLL"))
    print(Solution().checkRecord(s="LLPPPLL"))
    print("ok")
