#include <string>
#include <set>
#include <iostream>

class Solution
{
public:
    int minimizedStringLength(std::string s)
    {
        std::set<char> st;
        for (auto c : s)
        {
            st.insert(c);
        }
        return st.size();
    }
};

int main()
{
    std::cout << Solution().minimizedStringLength("aaabc");
}