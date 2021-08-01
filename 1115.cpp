/*

#----------------------------------------------------------------------#
#                                                                      #
#  version 0.0.1                                                       #
#  https://leetcode.com/problems/print-foobar-alternately/submissions/ #
#                                                                      #
#  Aleksiej Ostrowski, 2020                                            #
#                                                                      #
#  https://aleksiej.com                                                #
#                                                                      #
#----------------------------------------------------------------------#

*/


#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

std::mutex m;
std::condition_variable cv;

int flag = 1;

void task1(const int n)
{
    int i = 1;

    for(;;) {

        std::unique_lock lk(m);

        if (flag == 1) {
           std::cout << "foo";
           i++;
           flag = 2;
           lk.unlock();
           cv.notify_all();
        } else {
           cv.wait(lk, []{return flag == 1;});
        }

        if (i > n) return;
   }
}

void task2(const int n)
{
    int i = 1;

    for(;;) {

        std::unique_lock lk(m);

        if (flag == 2) {
           std::cout << "bar";
           i++;
           flag = 1;
           lk.unlock();
           cv.notify_all();
        } else {
           cv.wait(lk, []{return flag == 2;});
        }

        if (i > n) return;
   }
}

int main()
{
    flag = 1;

    std::thread t1(task1, 3);
    std::thread t2(task2, 3);

    t1.join();
    t2.join();
}
