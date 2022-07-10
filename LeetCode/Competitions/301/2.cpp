#include <set>

using namespace std;

class SmallestInfiniteSet
{
private:
  set<int> s;

public:
  SmallestInfiniteSet()
  {
    for (int i = 1; i <= 1000; i++)
    {
      s.insert(i);
    }
  }

  int popSmallest()
  {
    int num = *s.begin();
    s.erase(s.begin());
    return num;
  }

  void addBack(int num)
  {
    s.insert(num);
  }
};