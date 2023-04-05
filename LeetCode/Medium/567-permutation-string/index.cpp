#include "../../../stdc++.h"

using namespace std;

class Solution
{
public:
    bool checkInclusion(string s1, string s2)
    {
        int counter = s1.size();
        vector<int> letters(26, 0);
        vector<int> l(26, 0);
        for (auto s : s1)
            letters[s - 'a']++;

        l = letters;

        for (int i = 0, j = 0; j < s2.size(); j++)
        {
            if (l[s2[j] - 'a'] == 0)
            {
                l = letters;
                counter = s1.size();
                j = i;
                i++;
                continue;
            }

            l[s2[j] - 'a']--;
            counter--;

            if (counter == 0)
                return true;
        }

        return false;
    }
};

int main()
{
    cout << boolalpha << Solution().checkInclusion("adc", "dcda");
}