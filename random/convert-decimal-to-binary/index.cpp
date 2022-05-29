#include <iostream>
#include <string>

using namespace std;

int main()
{
  int number;
  cout << "what is the number you want converted to binary?\n";
  cin >> number;

  string bin = "";

  do
  {
    int remainder = number % 2;
    bin = to_string(remainder) + bin;
    number /= 2;
  } while (number > 0);

  cout << stoi(bin);
}