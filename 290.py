""" 290.py """

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
    def __init__(self):
        self.new_codes_s = {}
        self.new_codes_p = {}

    def new_code(self, new_el, new_codes):
        # print(new_el)
        if new_el in new_codes:
            return new_codes[new_el]
        else:
            code = len(new_codes) + 1
            new_codes[new_el] = code
            return code

    def wordPattern(self, pattern: str, s: str) -> bool:
        return [self.new_code(el, self.new_codes_p) for el in list(pattern)] == \
               [self.new_code(el, self.new_codes_s) for el in list(s.split())]

if __name__ == "__main__":
    # print(Solution().wordPattern(pattern="abba", s="dog cat cat dog"))
    print(Solution().wordPattern(pattern="a", s="a"))
    print("ok")
