#include <iostream>

using namespace std;

/**
 * Note: By default, pointers are passed by value. This means that if we point
 * the pointer to a different memory location inside the function, it still
 * points to the previous memory location outside the function. Here, we are
 * passing pointers by reference, so if we change the address stored in a
 * pointer inside the function, it is also changed outside the function.
 */

void delete_element(int *&arr, int size, int index)
{
  int *newArr = new int[size - 1];

  for (int i = 0; i < size - 1; i++)
  {
    if (i < index)
    {
      newArr[i] = arr[i];
    }
    else
    {
      newArr[i] = arr[i + 1];
    }
  }

  delete[] arr;

  // if arr is not a reference then all I would be doing is updating the arr
  // parameter. Because rememeber that when we pass something to the argurments
  // of a function it is passed by value ie: copied. So if we did not have a
  // reference to the original arr then all we would be doing is updating the
  // values for arr in the argument of the function.
  arr = newArr;
}

int main()
{
  int size = 4;
  int *a = new int[size];

  for (int i = 0; i < size; i++)
  {
    a[i] = i + 1;
  }

  for (int i = 0; i < size; i++)
  {
    cout << a[i];
  }

  cout << "\n------------------\n";

  delete_element(a, size, 2);

  for (int i = 0; i < size - 1; i++)
  {
    cout << a[i];
  }
}