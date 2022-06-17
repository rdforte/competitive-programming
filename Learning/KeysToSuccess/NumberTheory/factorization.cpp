#include <iostream>
#include <fstream>
#define lli long long int
using namespace std;

// time complexity is √n as we iterate over n i doubles each time.
int print_factors_count(lli N)
{
  int cnt = 0;

  for (int i = 1; i * i <= N; i++)
    // N % i gives us a number which is divisible and therefore a foctor
    if (N % i == 0)
    {
      cnt++;
      // Now what times i gives us N. We can rearange this by doing N/i.
      if (i != N / i)
        cnt++;
    }

  return cnt;
}

int main()
{
  ifstream cin("input.txt");

  int N;
  cin >> N;
  cout << print_factors_count(N);
  return 0;
}