""" 500.py """

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

class Solution:
    def __init__(self):
        self.rows = ["qwertyuiop", "asdfghjkl", "zxcvbnm"]
        for idx, row in enumerate(self.rows):
            self.rows[idx] = frozenset(row)
        # print(self.rows)    

    def findWords(self, words: List[str]) -> List[str]:
        res = []
        for word in words:
            set_word = frozenset(word.lower())
            for row in self.rows:
                if set_word.issubset(row):
                    res.append(word)
        return res            


if __name__ == '__main__':
    print(Solution().findWords(["Hello","Alaska","Dad","Peace"]))
    print('ok')

