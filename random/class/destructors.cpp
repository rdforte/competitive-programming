#include <iostream>
#include <string>
using namespace std;

class Collector
{
  int *list;
  int size;
  int capacity;

public:
  // Default constructor
  Collector()
  {
    // We must define the default values for the data members
    list = nullptr;
    size = 0;
    capacity = 0;
  }

  // Parameterized constructor
  Collector(int cap)
  {
    // The arguments are used as values
    capacity = cap;           // the capacity of the array.
    size = 0;                 // the number of elements currently in the array.
    list = new int[capacity]; // the array with a fixed memory capacity.
  }

  bool append(int v)
  {
    if (size < capacity)
    {
      list[size++] = v;
      return true;
    }
    return false;
  }

  // A simple print function
  void dump()
  {
    for (int i = 0; i < size; i++)
    {
      cout << list[i] << " ";
    }
    cout << endl;
  }

  ~Collector()
  {
    cout << "Deleting the object " << endl;
    if (capacity > 0)
      delete[] list;
  }
};

int main()
{
  Collector c(10);
  for (int i = 0; i < 15; i++)
  {
    cout << c.append(i) << endl;
  }
}