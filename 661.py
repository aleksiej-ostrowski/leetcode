""" 661.py """

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
import math
import itertools
import copy

class Solution:
    def imageSmoother(self, img: List[List[int]]) -> List[List[int]]:
        iter3m = list(itertools.product([0,1,-1], repeat=2))
        def checkXY(x: int, len_vx: int, y: int, len_vy: int) -> bool:
            return x < 0 or x > len_vx - 1 or y < 0 or y > len_vy - 1 
        new_img = copy.deepcopy(img)
        for ind_x, row in enumerate(img):
            for ind_y, _ in enumerate(row):
                sm = .0
                cnt_sm = 0
                # ls = []
                for p1, p2 in iter3m:
                    if not checkXY(ind_x+p1, len(img), ind_y+p2, len(row)):
                        sm += img[ind_x+p1][ind_y+p2]
                        # ls.append(img[ind_x+p1][ind_y+p2])
                        cnt_sm += 1
                # print(ls)        
                if cnt_sm > 0:
                    new_img[ind_x][ind_y] = math.floor(sm / cnt_sm)

        return new_img

if __name__ == '__main__':
    print(Solution().imageSmoother([[100,200,100],[200,50,200],[100,200,100]]))
    print('ok')

