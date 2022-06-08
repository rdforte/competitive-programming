#include <iostream>
using namespace std;

class A
{
protected:
  int ID;

public:
  A() : ID(0) {}
};

class B : virtual public A
{
public:
  int length;
};

class C : virtual public A
{
public:
  int radius;
};

class D : public B, public C
{
public:
  int getID() { return ID; }
};

int main(void)
{
  D d;
  cout << d.getID();
  return 0;
}