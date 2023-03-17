""" 1114.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2023      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from collections.abc import Callable

def printFirst():
    print("first")

def printSecond(): 
    print("second")

def printThird():
    print("third")

from threading import Thread, Lock

class Foo:
    def __init__(self):

        self.lock = Lock()
        self.num = 0

        t1 = Thread(target=self.first(printFirst))
        t2 = Thread(target=self.second(printSecond))
        t3 = Thread(target=self.third(printThird))

        t1.start()
        t2.start()
        t3.start()

        t1.join()
        t2.join()
        t3.join()
        
    def first(self, printFirst: 'Callable[[], None]') -> None:

        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()

        with self.lock:
            self.num = 1

    def second(self, printSecond: 'Callable[[], None]') -> None:
        
        while True:
            with self.lock:
                if self.num == 1:
                    break

        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()

        with self.lock:
            self.num = 2


    def third(self, printThird: 'Callable[[], None]') -> None:
        
        while True:
            with self.lock:
                if self.num == 2:
                    break

        # printThird() outputs "third". Do not change or remove this line.
        printThird()


if __name__ == '__main__':
    f = Foo()
    print("ok")
