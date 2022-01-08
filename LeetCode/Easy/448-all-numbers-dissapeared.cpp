#include <iostream>
#include <vector>
#include <set>

class Solution
{
public:
    // runtime is 197ms
    // time complexity is O(n) the Set is O(log.n) se we use O(n) as that is the slower of the two
    // space is the same O(n)
    /** NOTE:
     * The reason why the set is slower than the use of 2 vectors is because a set has each of its values
        in different locations in memory so the find is a lot slower compared to a vector which is just in one single location
        in memory
     */
    std::vector<int> findDisappearedNumbersV1(std::vector<int> &nums)
    {
        std::set<int> s;
        for (int n = 0; n < nums.size(); n++)
            s.insert(nums[n]);

        std::vector<int> n;
        for (int i = 1; i <= nums.size(); i++)
        {
            if (s.find(i) == s.end())
                n.push_back(i);
        }

        return n;
    }
    // runtime is 48ms
    // even though we use 2 vectors here to store the data this is still a lot more performant than a set
    // time complexity is O(n) we have at most two loops n.n but this can be reduced to n
    // space complexity is the same O(n) we have two vectors n.n but can be reduced to n
    std::vector<int> findDisappearedNumbers(std::vector<int> &nums)
    {
        std::vector<int> m(nums.size() + 1, 0); // we want the index to align with the corresponding number
        std::vector<int> e;

        for (int i = 0; i < nums.size(); i++)
            m[nums[i]]++;

        for (int i = 1; i < m.size(); i++)
            if (m[i] == 0)
                e.push_back(i);

        return e;
    }
};

int main()
{
    std::vector<int> v;

    int n;
    while (std::cin >> n)
    {
        v.push_back(n);
    }

    std::vector<int> sol = Solution().findDisappearedNumbers(v);
    for (int i = 0; i < sol.size(); i++)
    {
        std::cout << sol[i] << " ";
    }
}