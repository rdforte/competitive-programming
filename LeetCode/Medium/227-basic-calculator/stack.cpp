#include "../../../stdc++.h"

using namespace std;

int calculator(string str)
{
  stack<int> stk;
  int currentNum = 0;
  char currentOperator = '+';

  for (int i = 0; i < str.size(); i++)
  {
    if (isdigit(str[i]))
    {
      if (currentOperator == '+')
      {
        currentNum *= 10;
        currentNum += (str[i] - '0');
      }
    }
    if (i == str.size() - 1 || !isdigit(str[i]))
    {
      stk.push(currentNum);
      currentNum = 0;
      if (str[i] == ' ')
        continue;
      currentOperator = str[i];
    }
  }

  int res = 0;
  while (!stk.empty())
  {
    int num = stk.top();
    cout << "num: " << num << "\n";
    stk.pop();

    res += num;
  }

  return res;
}

int main()
{
  cout << calculator("3+4+5");
}