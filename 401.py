""" 401.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from itertools import combinations, product

DEBUG = False

def wrap(lst):
    if lst == []:
        return [0]
    else:
        return list(lst)

def wrap0(val):
    if type(val) == int and val == 0:
        return [val]
    else:
        return val

class Solution:

    def __init__(self):
        H = [8,4,2,1]
        M = [32,16,8,4,2,1]

        # print(H)
        # print(M)

        partH = []
        partM = []

        for el_len in range(1, len(H)+1): 
            partH += filter(lambda x: sum(x) <= 11, combinations(H, el_len))
        
        for el_len in range(1, len(M)+1): 
            partM += filter(lambda x: sum(x) <= 59, combinations(M, el_len))

        if DEBUG: print("H", [list(x) for x in partH])
        if DEBUG: print("M", [list(x) for x in partM])

        if DEBUG: print("="*30)

        self.partH = partH
        self.partM = partM


    def readBinaryWatch(self, turnedOn: int): # -> List[str]:
        res = []
        for num in range(0, turnedOn+1):
            a = num
            b = turnedOn - num
            if b < a:
                break
            if (a + b == turnedOn):
                if (0 <= a <= 3) and (0 <= b <= 6):   
                    if DEBUG: print(f"{turnedOn} = H {a} + M {b}")
                    H_ = wrap(list(filter(lambda x: len(x) == a, self.partH)))
                    M_ = wrap(list(filter(lambda x: len(x) == b, self.partM)))

                    if DEBUG: print("H_", H_)
                    if DEBUG: print("M_", M_)

                    if DEBUG:
                        print(f"a = {a}")
                        print(f"b = {b}")
                        print(f"H_ == [0] {H_== [0]}")
                        print(f"M_ == [0] {M_== [0]}")

                    if  (\
                        a == 0 and H_ == [0] \
                        ) or \
                        (\
                        b == 0 and M_ == [0] \
                        ) or \
                        (\
                        a != 0 and H_ != [0] and \
                        b != 0 and M_ != [0] \
                        ):

                        product_ = product(H_, M_)
                        product_ = [list(x) for x in product_]

                        for h_, m_ in product_:
                            h_ = wrap0(h_)
                            m_ = wrap0(m_)
                            if h_ == [0] and m_ == [0] and turnedOn > 0:
                                continue
                            if DEBUG: print(h_, m_)
                            res.append(f"{sum(h_)}:{sum(m_):02}")

                        if DEBUG: print("P ", list(product_))

                if (0 <= b <= 3) and (0 <= a <= 6) and (a != b):    
                    if DEBUG: print(f"{turnedOn} = H {b} + M {a}")
                    H_ = wrap(list(filter(lambda x: len(x) == b, self.partH)))
                    M_ = wrap(list(filter(lambda x: len(x) == a, self.partM)))

                    if DEBUG: print("H_", H_)
                    if DEBUG: print("M_", M_)
                     
                    if DEBUG:
                        print(f"a = {a}")
                        print(f"b = {b}")
                        print(f"H_ == [0] {H_== [0]}")
                        print(f"M_ == [0] {M_== [0]}")

                    if  (\
                        b == 0 and H_ == [0] \
                        ) or \
                        (\
                        a == 0 and M_ == [0] \
                        ) or \
                        (\
                        b != 0 and H_ != [0] and \
                        a != 0 and M_ != [0] \
                        ):

                        product_ = product(H_, M_)
                        product_ = [list(x) for x in product_]
                        
                        for h_, m_ in product_:
                            h_ = wrap0(h_)
                            m_ = wrap0(m_)
                            if h_ == [0] and m_ == [0] and turnedOn > 0:
                                continue
                            if DEBUG: print(h_, m_)
                            res.append(f"{sum(h_)}:{sum(m_):02}")
                        

                        if DEBUG: print("P ", list(product_))

        return res

if __name__ == '__main__':
    print(Solution().readBinaryWatch(6))
    # print(Solution().readBinaryWatch(1))
    # print(Solution().readBinaryWatch(9))
    print('ok')



