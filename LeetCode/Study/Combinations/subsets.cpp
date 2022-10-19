#include "../../../stdc++.h"

using namespace std;

/**
 * @brief
 *
 * @param set this is the set of numbers to process
 * @param i this is the index for the number in the set we are processing
 * @param subsetCur this is the current subset which we want to add to the set
 * @param subsetRes
 */
void generateSubsets(vector<int> &set, int i, vector<int> &subsetCur, vector<vector<int>> &subsetRes)
{
  // base case
  if (i >= set.size())
  {
    subsetRes.push_back(subsetCur);
    return;
  }

  // include set[i] (left decision)
  subsetCur.push_back(set[i]);
  generateSubsets(set, i + 1, subsetCur, subsetRes);

  // don't include set[i] (right decision)
  subsetCur.pop_back();
  generateSubsets(set, i + 1, subsetCur, subsetRes);
}

int main()
{
  vector<vector<int>> res;

  vector<int> set{1, 2, 3};

  vector<int> subsetCur;

  generateSubsets(set, 0, subsetCur, res);

  for (auto r : res)
  {
    cout << "-----------\n";
    for (auto i : r)
    {
      cout << i << ",";
    }
  }
}