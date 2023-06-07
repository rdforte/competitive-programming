#include <vector>
#include <unordered_map>

typedef std::vector<std::vector<int>> vv;

class Solution
{
public:
  long long matrixSumQueries(int n, vv &queries)
  {
    long long count = 0;
    int rows = n, cols = n;
    std::vector<bool> r(n, false), c(n, false);

    for (int i = queries.size() - 1; i >= 0; i--)
    {
      long long type = queries[i][0], index = queries[i][1], value = queries[i][2];
      if (type == 0 && !r[index])
      {
        r[index] = true;
        cols--;
        count += (rows * value);
      }

      if (type == 1 && !c[index])
      {
        c[index] = true;
        rows--;
        count += (cols * value);
      }
    }

    return count;
  }
};