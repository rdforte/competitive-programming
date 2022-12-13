#include <bits/stdc++.h>

using namespace std;

struct Monkey
{
  // private:
  vector<int> items;
  string oper;
  string additionMultiplier;
  int devisor;
  int monkeyTrue;
  int monkeyFalse;

public:
  Monkey() {}

  void addItem(int item)
  {
    items.push_back(item);
  }

  vector<int> getItems()
  {
    return items;
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

  int applyOperationAndRetrieveItem()
  {
    if (items.size() > 0)
    {
      int item = items.back();
      items.pop_back();

      int multiplyAddBy = additionMultiplier == "old" ? item : stoi(additionMultiplier);

      if (oper == "+")
      {
        return (item + multiplyAddBy) / 3;
      }
      else if (oper == "*")
      {
        return (item * multiplyAddBy) / 3;
      }
    }
    return 0;
  }

  void setDevisor(int d)
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

  int passToMonkey(int itemWorry)
  {
    return itemWorry % devisor == 0 ? monkeyTrue : monkeyFalse;
  }
};

int main()
{
  freopen("test.txt", "r", stdin);
  freopen("q1output.txt", "w", stdout);

  vector<Monkey> monkeys;

  string monkeyIndex;
  while (getline(cin, monkeyIndex))
  {
    cout << "------------\n";
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
    monkey.setDevisor(stoi(devisorMatch.str()));
    // cout << monkey.devisor << "\n";

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

  for (auto monkey : monkeys)
  {
  }
}