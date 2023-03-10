""" 412.py """

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
    def fizzBuzz(self, n: int) -> List[str]:
        return [rsl if (rsl := (\
                                 ("Fizz" if el % 3 == 0 else "") + \
                                 ("Buzz" if el % 5 == 0 else "")\
                               ) ) != "" \
                    else str(el) for el in range(1, n+1)]

if __name__ == '__main__':
    print(Solution().fizzBuzz(5))
    print('ok')

