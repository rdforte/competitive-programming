// https://leetcode.com/problems/longest-repeating-character-replacement/?envType=study-plan&id=level-1
#include <bits/stdc++.h>

using namespace std;

// Sliding Window
// time complexity = the window always moves forward with each iteration. O(n) where n is the length of the string.
// space complexity = we have an array which keeps track of all uppercase characters. We can call this m. the space
// is then O(m).

// Things to note here. If possible use an array over a map because it is ~100x faster. A map can have some large constant
// values under the hood which can increase the overall time of our algorithm. All letters are uppercase so we have 26
// letters therefore a vector with a constant length of 26 is a good approach here.

class Solution
{
public:
  int characterReplacement(string s, int k)
  {
    vector<int> characterCount(26, 0);
    int highestFrequency = 1;
    int maxLength = 0;

    for (int l = 0, h = 0; h < s.size(); h++)
    {
      characterCount[s[h] - 'A']++;
      highestFrequency = max(highestFrequency, characterCount[s[h] - 'A']);

      int curLen = (h + 1) - l;
      int changes = curLen - highestFrequency;
      if (changes <= k)
      {
        maxLength++;
      }
      else
      {
        characterCount[s[l] - 'A']--;
        l++;
      }
    }

    return maxLength;
  }
};

int main()
{
  cout << Solution().characterReplacement("BAAAB", 2);
}