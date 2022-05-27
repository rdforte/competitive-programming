#include <iostream>

using namespace std;

int main()
{
  int total_seconds = 3870;

  int hours = total_seconds / (60 * 60);
  total_seconds -= (hours * 60 * 60);

  int minutes = total_seconds / 60;
  total_seconds -= (minutes * 60);

  int seconds = total_seconds;

  cout << hours << " : " << minutes << " : " << seconds << "\n";
}