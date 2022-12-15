#include <bits/stdc++.h>

using namespace std;

typedef long long ll;

struct Monkey
{
private:
  vector<ll> items;
  string oper;
  string additionMultiplier;
  ll devisor;
  int monkeyTrue;
  int monkeyFalse;
  int operationMod;

public:
  Monkey()
  {
  }

  void addItem(ll item)
  {
    items.push_back(item);
  }

  bool hasNextItem()
  {
    return items.size() > 0;
  }

  void setOperator(string op)
  {
    oper = op;
  }
  void setAdditionMultiplier(string am)
  {
    additionMultiplier = am;
  }

  ll applyOperationAndRetrieveItem()
  {
    ll item = items[items.size() - 1];
    items.pop_back();

    int multiplyAddBy = additionMultiplier == "old" ? item : stoi(additionMultiplier);

    ll operatedItem;
    if (oper == "+")
    {
      operatedItem = (item + multiplyAddBy);
    }
    else if (oper == "*")
    {
      operatedItem = (item * multiplyAddBy);
    }
    return operatedItem % operationMod;
  }

  void setDevisor(ll d)
  {
    devisor = d;
  }

  void setThrowMonkeyTrue(int mt)
  {
    monkeyTrue = mt;
  }

  void setThrowMonkeyFalse(int mf)
  {
    monkeyFalse = mf;
  }

  int passToMonkey(ll itemWorry)
  {
    return itemWorry % devisor == 0 ? monkeyTrue : monkeyFalse;
  }

  void setOperationMod(ll mod)
  {
    operationMod = mod;
  }
};

int main()
{
  freopen("input.txt", "r", stdin);
  freopen("q2output.txt", "w", stdout);

  vector<Monkey> monkeys;
  ll operationMod = 1;

  string monkeyIndex;
  while (getline(cin, monkeyIndex))
  {
    int worryLevel = 0;
    Monkey monkey = Monkey();

    // Add items to monkey
    string itemsInput;
    getline(cin, itemsInput);
    smatch itemsMatch;
    regex reg("[0-9]+");
    auto iteratorStart = itemsInput.cbegin();
    while (regex_search(itemsInput, itemsMatch, reg))
    {
      monkey.addItem(stoi(itemsMatch.str(0)));
      itemsInput = itemsMatch.suffix().str();
    }

    // Add operation to monkey
    string operationInput;
    getline(cin, operationInput);
    smatch operationMatch;
    int operationIndex = 0;
    while (regex_search(operationInput, operationMatch, regex("old|\\+|\\*|[0-9]+")))
    {
      if (operationIndex == 1)
      {
        monkey.setOperator(operationMatch.str(0));
      }
      if (operationIndex == 2)
      {
        monkey.setAdditionMultiplier(operationMatch.str(0));
      }
      operationInput = operationMatch.suffix().str();
      operationIndex++;
    }

    // Add the devisor
    string devisorInput;
    smatch devisorMatch;
    getline(cin, devisorInput);
    regex_search(devisorInput, devisorMatch, regex("[0-9]+"));
    monkey.setDevisor((ll)stoi(devisorMatch.str()));
    operationMod *= (ll)stoi(devisorMatch.str());

    // Add the pass to monkey true
    string monkeyTrueInput;
    smatch monkeyTrueMatch;
    getline(cin, monkeyTrueInput);
    regex_search(monkeyTrueInput, monkeyTrueMatch, regex("[0-9]+"));
    monkey.setThrowMonkeyTrue(stoi(monkeyTrueMatch.str()));

    // Add the pass to monkey true
    string monkeyFalseInput;
    smatch monkeyFalseMatch;
    getline(cin, monkeyFalseInput);
    regex_search(monkeyFalseInput, monkeyFalseMatch, regex("[0-9]+"));
    monkey.setThrowMonkeyFalse(stoi(monkeyFalseMatch.str()));

    cin.ignore();

    monkeys.push_back(monkey);
  }

  for (int i = 0; i < monkeys.size(); i++)
  {
    monkeys[i].setOperationMod(operationMod);
  }

  vector<ll> monkeyInspect(monkeys.size(), 0);
  for (int round = 0; round < 10000; round++)
  {
    for (int i = 0; i < monkeys.size(); i++)
    {
      while (monkeys[i].hasNextItem())
      {
        ll itemWithWorry = monkeys[i].applyOperationAndRetrieveItem();
        int monkeyIndex = monkeys[i].passToMonkey(itemWithWorry);
        monkeys[monkeyIndex].addItem(itemWithWorry);
        monkeyInspect[i]++;
      }
    }
  }

  sort(monkeyInspect.begin(), monkeyInspect.end(), greater<int>());

  cout << (long long)(monkeyInspect[0] * monkeyInspect[1]);
}