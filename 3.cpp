/*
#----------------------------------------------------------------------------------------------#
#                                                                                              #
#  version 0.0.1                                                                               #
*  https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/   #
#                                                                                              #
#  Aleksiej Ostrowski, 2020                                                                    #
#                                                                                              #
#  https://aleksiej.com                                                                        #
#                                                                                              #
#----------------------------------------------------------------------------------------------#
*/

#include <string>
#include <iostream>


int max(const int a, const int b) {
    if (a > b) return a; else return b;
}

class Solution {
public:
    int lengthOfLongestSubstring(std::string s) {

        auto len_s = s.length();

        if (len_s <= 1) return len_s;

        std::string t;

        int max_len = -1;

        int i = 0;

        for (;;) {

           auto found = t.find(s[i]);

           if (found != -1) t = t.substr(found + 1);

           t += s[i];

           max_len = max(max_len, t.length());

           std::cout << "i = " << i <<  " s[" << i << "] = " << s[i] << " max_len = " << max_len << " string = " << t << std::endl;

           if (++i >= len_s) break;
        }

        return max_len;

    }
};

int main() {

    // std::string s = "aabac"; // 3
    // std::string s = "pwwkew"; // 3
    // std::string s = "bbbbb"; // 1
    // std::string s = " "; // 1
    std::string s = "aab"; // 2

    std::cout << "input " << s << std::endl;
    auto S = Solution();
    std::cout << S.lengthOfLongestSubstring(s) << std::endl;

    return 0;
}
