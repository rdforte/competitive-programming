#include <bits/stdc++.h>

using namespace std;

pair<bool, pair<int, int>> recursiveComparisonIsInOrder(string left, string right, int l, int r)
{
  cout << "left in: => " << left << "\n";
  cout << "right in: => " << right << "\n";
  for (int li = l, ri = r; li < left.size() and ri < right.size();)
  {
    cout << "left: " << left[li] << "\n";
    cout << "right: " << right[ri] << "\n";
    if (left[li] == ']')
      return {true, {li + 1, ri + 1}};
    if (right[ri] == ']')
      return {false, {li + 1, ri + 1}};

    if (isdigit(left[li]) and isdigit(right[ri]))
    {
      int leftDigit = stoi(to_string(left[ri])), rightDigit = stoi(to_string(right[ri]));
      if (leftDigit == rightDigit)
      {
        li++;
        ri++;
        continue;
      }

      if (leftDigit < rightDigit)
      {
        while (left[li] != ']')
        {
          li++;
        }
      }

      return {leftDigit < rightDigit, {li + 1, ri + 1}};
    }

    if (left[li] == '[' or right[ri] == '[')
    {
      cout << "STEP IN ---------"
           << "\n";
      auto val = recursiveComparisonIsInOrder(left.substr(li + 1, left.size() - li - 1), right.substr(ri + 1, right.size() - ri - 1), li, ri);
      cout << "STEP OUT --------"
           << "\n";
      if (!val.first)
        return val;
      li += val.second.first;
      ri += val.second.second;
    }
    li++;
    ri++;
  }

  return {true, {0, 0}};
}

string addBracketsToString(string ogString, int i)
{
  string newString = ogString.substr(0, i) + "[";
  int iterator = i;
  while (iterator < ogString.size() and ogString[iterator] != ',' and ogString[iterator] != ']')
  {
    newString += ogString[iterator];
    iterator++;
  }
  newString += "]" + ogString.substr(iterator, ogString.size() - 1);

  return newString;
}

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  int count = 0;
  int index = 1;

  string left;
  string right;
  while (getline(cin, left))
  {
    getline(cin, right);
    cout << boolalpha;
    for (int l = 0, r = 0; l < left.size() or r < left.size();)
    {
      if (left[l] == '[' and isdigit(right[r]))
      {
        right = addBracketsToString(right, r);
      }
      if (right[r] == '[' and isdigit(left[l]))
      {
        left = addBracketsToString(left, l);
      }
      l++;
      r++;
    }
    cout << left << "\n";
    cout << right << "\n";
    if (recursiveComparisonIsInOrder(left, right, 0, 0).first)
    {
      count += index;
    }
    index++;
    cin.ignore();
  }

  cout << "\n\n\n"
       << count;
}