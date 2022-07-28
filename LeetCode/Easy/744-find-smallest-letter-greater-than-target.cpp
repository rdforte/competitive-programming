// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
#include <iostream>
#include <vector>
#include <algorithm>

class Solution
{
public:
  // You can compare characters using comparison operators.
  // Spacetime = O(1) no extra space needed.
  // Runtime = O(n) have to loop through all items if target is last item in array.
  char nextGreatestLetter(std::vector<char> &letters, char target)
  {
    for (int i = 0; i < letters.size(); i++)
    {
      if (letters[i] > target)
        return letters[i];
    }
    return letters[0];
  };
};

int main()
{
  std::vector<char> v{'a', 'b', 'c', 'd'};

  char sol = Solution().nextGreatestLetter(v, 'd');
  std::cout << sol;
}