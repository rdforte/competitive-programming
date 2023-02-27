#include <bits/stdc++.h>

using namespace std;

// Recursion

class Solution
{
public:
  string decodeString(string s)
  {
    return recursion(s, 0).first;
  }

private:
  pair<string, int> recursion(const string &s, int p)
  {

    string res = "";
    string repetitionsString = "";
    for (int i = p; i < s.size(); i++)
    {
      if (isdigit(s[i]))
      {
        repetitionsString += s[i];
        continue;
      }

      if (s[i] == '[')
      {
        int reps = stoi(repetitionsString);
        int next;
        for (int r = 1; r <= reps; r++)
        {
          auto si = recursion(s, i + 1);
          res += si.first;
          next = si.second;
        }
        i = next;
        repetitionsString = "";
        continue;
      }

      if (s[i] == ']')
      {
        return {res, i};
      }

      res += s[i];
    }

    return {res, p};
  }
};

int main()
{
  cout << boolalpha << (Solution().decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef");
}