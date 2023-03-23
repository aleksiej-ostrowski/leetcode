""" 463.py """

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
# from itertools import product

class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        perimeter = 0
        size_x = len(grid)
        # comb = set(product([-1,0,1], repeat=2)) - {(0,0)}
        comb = [(0,1),(1,0),(-1,0),(0,-1)]
        # print(comb)
        for x, row in enumerate(grid):
            size_y = len(row)
            for y, col in enumerate(row):
                el = grid[x][y]
                if el == 0:
                    continue
                for (ad_x, ad_y) in comb:
                    new_x = x + ad_x
                    new_y = y + ad_y
                    if (new_x >= 0) and (new_x < size_x) and \
                       (new_y >= 0) and (new_y < size_y):
                        # print(new_x, new_y)
                        if grid[new_x][new_y] == 0:
                            perimeter += 1
                    else:
                        perimeter += 1

        return perimeter

if __name__ == "__main__":
    # print(Solution().islandPerimeter([[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]))
    print(Solution().islandPerimeter([[1,0]]))
    print("ok")
