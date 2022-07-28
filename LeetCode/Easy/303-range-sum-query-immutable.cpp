// https://leetcode.com/problems/range-sum-query-immutable/
#include <iostream>
#include <vector>

// BruteForce
class NumArrayV1
{
    std::vector<int> n;

public:
    NumArrayV1(std::vector<int> &nums)
    {
        n = nums;
    }

    int sumRange(int left, int right)
    {
        int sum = 0;
        for (int i = left; i <= right; i++)
        {
            sum += n[i];
        }
        return sum;
    }
};

// Preffered approach using Dynamic Programming
// Most expensive call is the initialisation of the NumArray which is O(n)
// then each call to sumRange is O(1)
// runtime = O(1)
// spacetime = O(n) keeping track of the vector n
class NumArray
{
    std::vector<int> n;

public:
    NumArray(std::vector<int> &nums)
    {
        n = nums;
        for (int i = 1; i < n.size(); i++)
        {
            n[i] += n[i - 1];
        }
    }

    int sumRange(int left, int right)
    {
        return left == 0 ? n[right] : n[right] - n[left - 1];
    }
};

int main()
{
    std::vector<int> v{-2, 0, 3, -5, 2, -1};
    NumArray *numArray = new NumArray(v);
    std::cout << numArray->sumRange(0, 5);
}