# python3.10 42.py

# ------------------------------------------------------#
#                                                       #
#  version 0.0.1                                        #
#  https://leetcode.com/problems/trapping-rain-water/   #
#                                                       #
#  Aleksiej Ostrowski, 2022                             #
#                                                       #
#  https://aleksiej.com                                 #
#                                                       #
# ------------------------------------------------------#


class Solution:
    def __init__(self, height: list[int]):
        self.height = height

    def trap(self) -> int:

        h = self.height
        w = [1] * len(h)
        sm = 0

        while True:

            h_new = []
            w_new = []

            ind_start = 0
            ind = ind_start + 1

            while True:

                flag = 1
                w_d = w[ind_start]

                while ind <= len(h) - 1:
                    if h[ind_start] == h[ind]:
                        w_d += w[ind]
                        ind += 1
                    else:
                        flag = 2
                        break

                w_new.append(w_d)
                h_new.append(h[ind_start])

                if flag == 2:
                    ind_start = ind
                elif flag == 1:
                    break

                ind += 1

            cases = 0
            for i, e in enumerate(h_new):
                if (i > 0) and (i < len(h_new) - 1):
                    if h_new[i - 1] > h_new[i] < h_new[i + 1]:
                        d = min(h_new[i - 1], h_new[i + 1]) - h_new[i]
                        h_new[i] += d
                        sm += d * w_new[i]
                        cases += 1

            if cases == 0:
                break

            h = h_new
            w = w_new

        return sm


# Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
# Output: 6

# Input: height = [4,2,0,3,2,5]
# Output: 9

if __name__ == "__main__":
    height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    print(height)
    x = Solution(height)
    print(x.trap())
