#include <iostream>
#include <string>
using namespace std;

class Account
{
  string name;

public:
  // Default constructor
  Account()
  {
    // We must define the default values for day, month, and year
    name = "no name";
  }

  // Parameterized constructor
  Account(string name)
  {
    // The arguments are used as values
    this->name = name;
  }

  // Copy constructor, copies object into new account.
  Account(const Account &other)
  {
    name = other.name;
  }

  // A simple print function
  void printAccount()
  {
    cout << name << endl;
  }
};

int main()
{
  // Call the Account constructor to create its object;

  Account a("ryan"); // Object created with specified values!
  a.printAccount();

  Account b(a);
  b.printAccount();
}