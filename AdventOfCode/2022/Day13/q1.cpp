#include <bits/stdc++.h>

using namespace std;

pair<int, pair<int, int>> recursiveComparisonIsInOrder(string left, string right, int l, int r)
{
  // cout << "left in: => " << left << "\n";
  // cout << "right in: => " << right << "\n";
  for (int li = l, ri = r; li < left.size() and ri < right.size();)
  {
    // cout << "left: " << left[li] << "\n";
    // cout << "right: " << right[ri] << "\n";
    if (left[li] == ']' and right[ri] == ']')
      return {0, {li + 1, ri + 1}};
    if (left[li] == ']')
      return {1, {li + 1, ri + 1}};
    if (right[ri] == ']')
      return {-1, {li + 1, ri + 1}};

    if (isdigit(left[li]) and isdigit(right[ri]))
    {
      int leftDigit = stoi(to_string(left[li])), rightDigit = stoi(to_string(right[ri]));
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
        return {1, {0, 0}};
      }

      return {-1, {0, 0}};
    }

    if (left[li] == '[' and right[ri] == '[')
    {
      // cout << "STEP IN ---------"
      //      << "\n";
      auto val = recursiveComparisonIsInOrder(left.substr(li + 1, left.size() - li - 1), right.substr(ri + 1, right.size() - ri - 1), li, ri);
      // cout << "STEP OUT --------"
      //      << "\n";
      if (val.first != 0)
        return val;
      li += val.second.first;
      ri += val.second.second;
    }

    if (left[li] == '[' and isdigit(right[ri]))
    {
      li++;
      if (stoi(to_string(left[li])) == stoi(to_string(right[ri])))
      {
        while (left[li] != ']')
        {
          li++;
        }
      }
      else if (stoi(to_string(left[li])) <= stoi(to_string(right[ri])))
      {
        return {1, {0, 0}};
      }
      else
      {
        return {-1, {0, 0}};
      }
    }

    if (isdigit(left[li]) and right[ri] == '[')
    {
      ri++;
      if (stoi(to_string(left[li])) == stoi(to_string(right[ri])))
      {
        while (right[ri] != ']')
        {
          ri++;
        }
      }
      else if (stoi(to_string(left[li])) <= stoi(to_string(right[ri])))
      {
        return {1, {0, 0}};
      }
      else
      {
        return {-1, {0, 0}};
      }
    }
    li++;
    ri++;
  }

  return {1, {0, 0}};
}

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  int count = 0;
  int index = 1;

  string left;
  string right;
  while (getline(cin, left) and getline(cin, right))
  {
    int result = recursiveComparisonIsInOrder(left, right, 0, 0).first;
    cout << "--------------> " << result << "\n";
    if (result == 1)
    {
      count += index;
    }
    // cout << recursiveComparisonIsInOrder(left, right, 0, 0).first;
    index++;
    cin.ignore();
  }

  cout << "\n\n\n"
       << count;
}