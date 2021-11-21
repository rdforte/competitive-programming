
#include <iostream>
#include <vector>

class RangeFreqQuery
{
private:
    std::vector<int> a;

public:
    RangeFreqQuery(std::vector<int> &arr)
    {
        a = arr;
    }

    int query(int left, int right, int value)
    {
        int count = 0;
        for (int i = left; i <= right; i++)
        {
            if (a[i] == value)
            {
                count++;
            }
        }
        return count;
    }
};

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * RangeFreqQuery* obj = new RangeFreqQuery(arr);
 * int param_1 = obj->query(left,right,value);
 */