""" 455.py """

from typing import List

# https://leetcode.com/problems/assign-cookies/solutions/3277329/455-solution-with-step-by-step-explanation/
class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        len_g = len(g)
        len_s = len(s)
        if len_g > 1:
            g.sort()
        if len_s > 1:     
            s.sort()
        i, j = 0, 0
        while i < len_g and j < len_s:
            if s[j] >= g[i]:
                i += 1
            j += 1
        return i
if __name__ == '__main__':
    print(Solution().findContentChildren(g = [1,2,3], s = [1,1]))
    print('ok')

