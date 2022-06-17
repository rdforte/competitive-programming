#include <iostream>
#include <vector>
#include <algorithm>
#include <fstream>
using namespace std;

// A prime number is a number greater than one whose only factors are 1 and itself.

// time complexity is √n as we iterate over n i doubles each time.
string is_prime(int N)
{
  for (int i = 2; i * i <= N; i++)
    if (N % i == 0)
    {
      return "no";
    }
  return "yes";
}

// the overal time complexity is t*√n as we calculate is_prime for each t.
int main()
{
  ifstream cin("input.txt");
  int T, N;

  cin >> T;
  while (T--)
  {
    cin >> N;
    cout << is_prime(N) << "\n";
  }
  return 0;
}