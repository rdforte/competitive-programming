#include <bits/stdc++.h>

using namespace std;

pair<bool, int> recursiveComparisonIsInOrder(string left, string right)
{
  cout << "left in: => " << left << "\n";
  cout << "right in: => " << right << "\n";
  for (int i = 0; i < left.size(); i++)
  {
    cout << "left: " << left[i] << "\n";
    cout << "right: " << right[i] << "\n";
    if (left[i] == ']')
      return {true, i + 1};
    if (right[i] == ']')
      return {false, i + 1};

    if (isdigit(left[i]) and isdigit(right[i]))
    {
      int leftDigit = stoi(to_string(left[i])), rightDigit = stoi(to_string(right[i]));
      if (leftDigit == rightDigit)
        continue;

      return {leftDigit < rightDigit, i + 1};
    }

    if (left[i] == '[' or right[i] == '[')
    {
      cout << "STEP IN ---------"
           << "\n";
      auto val = recursiveComparisonIsInOrder(left.substr(i + 1, left.size() - i - 1), right.substr(i + 1, right.size() - i - 1));
      cout << "STEP OUT --------"
           << "\n";
      if (!val.first)
        return val;
      i += val.second;
    }
  }

  return {true, 0};
}

int main()
{
  freopen("test.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  string left;
  string right;
  while (getline(cin, left))
  {
    getline(cin, right);
    cout << boolalpha;
    cout << recursiveComparisonIsInOrder(left, right).first << "\n";
    cin.ignore();
  }
}